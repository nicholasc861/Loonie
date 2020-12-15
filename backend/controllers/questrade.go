package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/nicholasc861/qapi"
	"github.com/nicholasc861/questrack-backend/models"
	"gorm.io/gorm/clause"
)

func AddRefreshToken(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	currentUser := &models.User{
		UserID: contextUserID.(uint),
	}

	err := json.NewDecoder(req.Body).Decode(currentUser)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Invalid Request",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Model(&currentUser).Table("users").Update("refresh_token", currentUser.RefreshToken).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error creating record in Database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	json.NewEncoder(res).Encode(currentUser)
}

func UpdateRefreshToken(c *qapi.Client, currentUser models.User) error {
	newRefreshToken := c.Credentials.RefreshToken

	if err := db.Model(&currentUser).Table("users").Update("refresh_token", newRefreshToken).Error; err != nil {

		return err
	}
	return nil
}

func GetQuestradeAccounts(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	user := models.User{}
	newAccounts := []models.QuestradeAccount{}
	accounts := []models.QuestradeAccount{}

	if err := db.Table("users").Find(&user).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error using refresh token, please contact support.",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	err = UpdateRefreshToken(client, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, qAccounts, err := client.GetAccounts()
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error getting new accounts, please contact support.",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	for _, account := range qAccounts {
		accountNumber, _ := strconv.ParseUint(account.Number, 10, 64)
		var accountStatus bool
		if account.Status == "Active" {
			accountStatus = true
		} else {
			accountStatus = false
		}

		newAccounts = append(newAccounts, models.QuestradeAccount{
			AccountID:   uint(accountNumber),
			UserID:      contextUserID.(uint),
			AccountType: account.Type,
			Status:      accountStatus,
		})
	}

	if err := db.Table("user_accounts").Clauses(clause.OnConflict{DoNothing: true}).Create(&newAccounts).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error creating record in Database",
		}
		fmt.Println(err)

		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("user_accounts").Find(&accounts).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error creating record in Database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	json.NewEncoder(res).Encode(accounts)
}

func GetAllQuestradePositions(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	user := models.User{}
	accounts := []models.QuestradeAccount{}
	newPositions := []models.QuestradePosition{}
	positions := []models.QuestradePosition{}

	if err := db.Table("users").Find(&user).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("user_accounts").Find(&accounts).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error using refresh token, please contact support.",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	err = UpdateRefreshToken(client, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Only Options have Numbers within
	// Ex. AAPL19Feb21C125.00 (Option) vs AAPL (Stock)
	checkIsOption := regexp.MustCompile(`[0-9]`).MatchString

	for _, account := range accounts {
		if account.Status == true {
			qPositions, err := client.GetPositions(fmt.Sprint(account.AccountID))
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, position := range qPositions.Positions {
				var positionStatus bool
				var positionIsOption bool
				if position.OpenQuantity == 0 {
					positionStatus = true
				} else {
					positionStatus = false
				}

				if checkIsOption(position.Symbol) {
					positionIsOption = true
				} else {
					positionIsOption = false
				}

				newPositions = append(newPositions, models.QuestradePosition{
					AccountID:         account.AccountID,
					QuestradeID:       uint(position.SymbolID),
					Symbol:            position.Symbol,
					OpenQuantity:      position.OpenQuantity,
					ClosedQuantity:    position.ClosedQuantity,
					AverageEntryPrice: position.AverageEntryPrice,
					TotalEntry:        position.TotalCost,
					IsOption:          positionIsOption,
					Status:            positionStatus,
				})
			}

		}
	}

	// Only add Position if it does not exist, else update the position
	if err := db.Table("user_positions").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "account_id"}, {Name: "questrade_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"open_quantity", "closed_quantity", "average_entry_price", "total_entry", "status"}),
	}).Create(&newPositions).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding record in the Database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	// Find all positions associated with the accounts passed in
	if err := db.Table("user_positions").Find(&positions).Where("account_id in ?", accounts).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding record in the Database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}
	json.NewEncoder(res).Encode(positions)
}

// func GetQuestradePositions(res http.ResponseWriter, req *http.Request) {

// }

// GetAllQuestradeBalances Gets Account Balances from Questrade
func GetAllQuestradeBalances(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	user := models.User{}
	accounts := []models.QuestradeAccount{}
	balances := []qapi.AccountBalances{}

	if err := db.Table("users").Find(&user).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("user_accounts").Find(&accounts).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error using refresh token, please contact support.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	err = UpdateRefreshToken(client, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, account := range accounts {
		if account.Status == true {
			qBalance, err := client.GetBalances(fmt.Sprint(account.AccountID))
			if err != nil {
				return
			}

			balances = append(balances, qBalance)
		}

		json.NewEncoder(res).Encode(balances)
	}
}

// GetAllQuestradeQuote Gets the Quote for the specified symbol (by QuestradeID)
func GetHistoricalQuote(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	user := models.User{}
	historicalQuotes := []models.QuestradeQuote{}
	urlVars := mux.Vars(req)

	if err := db.Table("users").Find(&user).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error using refresh token, please contact support.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	positionID, _ := strconv.ParseInt(urlVars["positionID"], 10, 64)

	err = UpdateRefreshToken(client, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	qQuote, err := client.GetQuote(int(positionID))
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding quote for specified position.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("historical_quotes").Find(&historicalQuotes).Where("questrade_id = ?", positionID).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding quotes for specified position.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	quote := models.QuestradeQuote{
		QuestradeID:    uint(positionID),
		Symbol:         qQuote.Symbol,
		BidPrice:       qQuote.BidPrice,
		LastTradePrice: qQuote.LastTradePrice,
		TimeQuoted:     time.Now().Unix(),
		OpenPrice:      qQuote.OpenPrice,
	}

	if err := db.Table("historical_quotes").Create(&quote).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error creating row in database.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	historicalQuotes = append(historicalQuotes, quote)

	json.NewEncoder(res).Encode(quote)
}

func GetCurrentQuote(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	user := models.User{}
	urlVars := mux.Vars(req)

	if err := db.Table("users").Find(&user).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error using refresh token, please contact support.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	positionID, _ := strconv.ParseInt(urlVars["positionID"], 10, 64)

	err = UpdateRefreshToken(client, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	qQuote, err := client.GetQuote(int(positionID))
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding quote for specified position.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	quote := models.QuestradeQuote{
		QuestradeID:    uint(positionID),
		Symbol:         qQuote.Symbol,
		BidPrice:       qQuote.BidPrice,
		LastTradePrice: qQuote.LastTradePrice,
		TimeQuoted:     time.Now().Unix(),
		OpenPrice:      qQuote.OpenPrice,
	}

	json.NewEncoder(res).Encode(quote)
}

func GetLivePL(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	positions := []models.QuestradePosition{}
	user := models.User{}
	resp := make([]map[string]interface{}, 0)

	reqStruct := struct {
		Positions []models.QuestradePosition `json:"positions"`
	}{}

	json.NewDecoder(req.Body).Decode(&reqStruct)

	positions = reqStruct.Positions

	if err := db.Table("users").Find(&user).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error using refresh token, please contact support.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	err = UpdateRefreshToken(client, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, position := range positions {
		qQuote, err := client.GetQuote(int(position.QuestradeID))
		fmt.Println(qQuote)
		if err != nil {
			var resp = map[string]interface{}{
				"status":  false,
				"message": "Error finding quote for specified position.",
			}
			json.NewEncoder(res).Encode(resp)
			return
		}
		var multiplierOption float32 = 1
		if position.IsOption {
			multiplierOption = 100
		}
		PL := (qQuote.BidPrice * position.OpenQuantity * multiplierOption) - (position.AverageEntryPrice * position.OpenQuantity * multiplierOption)
		roundedPL := math.Round(float64(PL)*100) / 100
		isNegative := math.Signbit(roundedPL)

		var newResp = map[string]interface{}{
			"status":     true,
			"isNegative": isNegative,
			"P_L":        roundedPL,
		}

		resp = append(resp, newResp)

	}
	fmt.Println(resp)
	json.NewEncoder(res).Encode(resp)
}

func GetSearchResults(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	user := models.User{}
	urlVars := mux.Vars(req)

	if err := db.Table("users").Find(&user).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error using refresh token, please contact support.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	err = UpdateRefreshToken(client, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	results, err := client.SearchSymbols(urlVars["searchterm"], 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(res).Encode(results)
}

func GetOptionsChain(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	user := models.User{}
	urlVars := mux.Vars(req)

	if err := db.Table("users").Find(&user).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error finding specified user in database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error using refresh token, please contact support.",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	err = UpdateRefreshToken(client, user)
	if err != nil {
		fmt.Println(err)
		return
	}

	questradeID, err := strconv.ParseInt(urlVars["questradeID"], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	results, err := client.GetOptionChain(int(questradeID))
	fmt.Println(results)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(res).Encode(results)
}
