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

// AddRefreshToken Update refresh token for user
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

// GetQuestradeAccounts Retrieves all active accounts on Questrade
func GetQuestradeAccounts(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id").(uint)
	newAccounts := []models.QuestradeAccount{}
	accounts := []models.QuestradeAccount{}

	user, err := GetUserInfo(contextUserID)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem finding user",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := GetQuestradeClient(user)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem retrieving Questrade Client. Invalid Refresh Token or Problem Updating Refresh Token.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	_, qAccounts, err := client.GetAccounts()
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error getting new accounts, please contact support.",
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
			UserID:      contextUserID,
			AccountType: account.Type,
			Status:      accountStatus,
		})
	}

	if err := db.Table("user_accounts").Clauses(clause.OnConflict{DoNothing: true}).Create(&newAccounts).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error creating record in database.",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("user_accounts").Where("user_id = ?", contextUserID).Find(&accounts).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding records in the database.",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	res.WriteHeader(200)
	var resp = map[string]interface{}{
		"status":     200,
		"statusText": "ok",
		"message":    "Successfully retrieved accounts.",
		"data":       accounts,
	}
	json.NewEncoder(res).Encode(resp)
}

// GetAllQuestradePositions Retrieves all positions in all Questrade accounts
func GetAllQuestradePositions(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id").(uint)
	accounts := []models.QuestradeAccount{}
	newPositions := []models.QuestradePosition{}
	closedPositions := []models.QuestradePosition{}
	positions := []models.QuestradePosition{}

	user, err := GetUserInfo(contextUserID)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem finding user",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := GetQuestradeClient(user)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem retrieving Questrade Client. Invalid Refresh Token or Problem Updating Refresh Token.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("user_accounts").Where("user_id = ?", contextUserID).Find(&accounts).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding specified user in database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	// Only Options have Numbers within
	// Ex. AAPL19Feb21C125.00 (Option) vs AAPL (Stock)
	checkIsOption := regexp.MustCompile(`[0-9]`).MatchString

	for _, account := range accounts {
		if account.Status == true {
			qPositions, err := client.GetPositions(fmt.Sprint(account.AccountID))
			if err != nil {
				res.WriteHeader(500)
				var resp = map[string]interface{}{
					"status":     500,
					"statusText": "INTERNAL_SERVER_ERROR",
					"message":    "Problem retrieving positions from Questrade",
				}
				fmt.Println(err)
				json.NewEncoder(res).Encode(resp)
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

				if position.AverageEntryPrice != 0 {
					newPositions = append(newPositions, models.QuestradePosition{
						AccountID:         account.AccountID,
						QuestradeID:       uint(position.SymbolID),
						Symbol:            position.Symbol,
						OpenQuantity:      position.OpenQuantity,
						ClosedQuantity:    position.ClosedQuantity,
						ClosedPNL:         position.ClosedPnL,
						AverageEntryPrice: position.AverageEntryPrice,
						TotalEntry:        position.TotalCost,
						IsOption:          positionIsOption,
						Status:            positionStatus,
					})
				} else {
					closedPositions = append(closedPositions, models.QuestradePosition{
						AccountID:         account.AccountID,
						QuestradeID:       uint(position.SymbolID),
						Symbol:            position.Symbol,
						OpenQuantity:      position.OpenQuantity,
						ClosedQuantity:    position.ClosedQuantity,
						ClosedPNL:         position.ClosedPnL,
						AverageEntryPrice: position.AverageEntryPrice,
						TotalEntry:        position.TotalCost,
						IsOption:          positionIsOption,
						Status:            positionStatus,
					})
				}
			}
		}
	}

	// Only add Position if it does not exist, else update the position
	if err := db.Table("user_positions").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "account_id"}, {Name: "questrade_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"open_quantity", "closed_quantity", "average_entry_price", "total_entry", "status"}),
	}).Create(&newPositions).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding record in the Database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	// Closed positions change the average-entry-price, separate SQL query
	if err := db.Table("user_positions").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "account_id"}, {Name: "questrade_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"open_quantity", "closed_quantity", "total_entry", "status"}),
	}).Create(&closedPositions).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding record in the Database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	// Find all positions associated with the accounts passed in
	if err := db.Table("user_positions").Where("account_id in ?", accounts).Find(&positions).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding record in the Database",
		}
		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}
	res.WriteHeader(200)
	var resp = map[string]interface{}{
		"status":     200,
		"statusText": "ok",
		"message":    "Successfully retrieved accounts.",
		"data":       positions,
	}
	json.NewEncoder(res).Encode(resp)
}

// GetAllQuestradeBalances Gets Account Balances from Questrade
func GetAllQuestradeBalances(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id").(uint)
	accounts := []models.QuestradeAccount{}
	balances := []qapi.AccountBalances{}

	user, err := GetUserInfo(contextUserID)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem finding user",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := GetQuestradeClient(user)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem retrieving Questrade Client. Invalid Refresh Token or Problem Updating Refresh Token.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("users").Where("user_id = ?", contextUserID).Find(&user).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error creating entry in database",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("user_accounts").Where("user_id = ?", contextUserID).Find(&accounts).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error creating entry in database",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
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
	}

	res.WriteHeader(200)
	var resp = map[string]interface{}{
		"status":     200,
		"statusText": "ok",
		"message":    "Successfully retrieved historical quote for position",
		"data":       balances,
	}
	json.NewEncoder(res).Encode(resp)
}

// GetHistoricalQuote Gets the Quote for the specified symbol (by QuestradeID)
func GetHistoricalQuote(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id").(uint)
	historicalQuotes := []models.QuestradeQuote{}
	urlVars := mux.Vars(req)

	user, err := GetUserInfo(contextUserID)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem finding user",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := GetQuestradeClient(user)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem retrieving Questrade Client. Invalid Refresh Token or Problem Updating Refresh Token.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	positionID, _ := strconv.ParseInt(urlVars["positionID"], 10, 64)

	qQuote, err := client.GetQuote(int(positionID))
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding quote with specified id",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	quoteInfo, err := client.GetSymbols(int(positionID))
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding info for specified position.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("historical_quotes").Where("questrade_id = ?", positionID).Find(&historicalQuotes).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding quote with specified id",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	quote := models.QuestradeQuote{
		QuestradeID:       uint(positionID),
		Description:       quoteInfo[0].Description,
		Symbol:            qQuote.Symbol,
		BidPrice:          qQuote.BidPrice,
		LastTradePrice:    qQuote.LastTradePrice,
		TimeQuoted:        time.Now().Unix(),
		OpenPrice:         qQuote.OpenPrice,
		PrevDayClosePrice: quoteInfo[0].PrevDayClosePrice,
	}

	if err := db.Table("historical_quotes").Create(&quote).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error creating entry in database",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	historicalQuotes = append(historicalQuotes, quote)
	res.WriteHeader(200)
	var resp = map[string]interface{}{
		"status":     200,
		"statusText": "ok",
		"message":    "Successfully retrieved historical quote for position",
		"data":       historicalQuotes,
	}
	json.NewEncoder(res).Encode(resp)
}

// GetCurrentQuote Retrieves live quote for a specified position. Does not retrieve historical.
func GetCurrentQuote(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id").(uint)
	urlVars := mux.Vars(req)

	user, err := GetUserInfo(contextUserID)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem finding user",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := GetQuestradeClient(user)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem retrieving Questrade Client. Invalid Refresh Token or Problem Updating Refresh Token.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	positionID, _ := strconv.ParseInt(urlVars["positionID"], 10, 64)

	qQuote, err := client.GetQuote(int(positionID))
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding quote for specified position.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	quoteInfo, err := client.GetSymbols(int(positionID))
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding info for specified position.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	quote := models.QuestradeQuote{
		QuestradeID:       uint(positionID),
		Description:       quoteInfo[0].Description,
		Symbol:            qQuote.Symbol,
		BidPrice:          qQuote.BidPrice,
		LastTradePrice:    qQuote.LastTradePrice,
		TimeQuoted:        time.Now().Unix(),
		OpenPrice:         qQuote.OpenPrice,
		PrevDayClosePrice: quoteInfo[0].PrevDayClosePrice,
	}

	if err := db.Table("historical_quotes").Create(&quote).Error; err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error creating entry in database",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	res.WriteHeader(200)
	var resp = map[string]interface{}{
		"status":     200,
		"statusText": "ok",
		"message":    "Successfully retrieved quote for position",
		"data":       quote,
	}
	json.NewEncoder(res).Encode(resp)
}

// GetLivePL Retrieves Live Profit/Loss for all positions
func GetLivePL(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id").(uint)
	positions := []models.QuestradePosition{}
	allLivePL := make([]map[string]interface{}, 0)

	reqStruct := struct {
		Positions []models.QuestradePosition `json:"positions"`
	}{}

	json.NewDecoder(req.Body).Decode(&reqStruct)
	positions = reqStruct.Positions

	user, err := GetUserInfo(contextUserID)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem finding user",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := GetQuestradeClient(user)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem retrieving Questrade Client. Invalid Refresh Token or Problem Updating Refresh Token.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	for _, position := range positions {
		qQuote, err := client.GetQuote(int(position.QuestradeID))
		fmt.Println(qQuote)
		if err != nil {
			res.WriteHeader(500)
			var resp = map[string]interface{}{
				"status":     500,
				"statusText": "INTERNAL_SERVER_ERROR",
				"message":    "Error retrieving quote for specified position.",
			}

			fmt.Println(err)
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

		allLivePL = append(allLivePL, newResp)

	}

	res.WriteHeader(200)
	var resp = map[string]interface{}{
		"status":     200,
		"statusText": "ok",
		"message":    "Successfully retrieved Live PL for positions",
		"data":       allLivePL,
	}
	json.NewEncoder(res).Encode(resp)
}

// GetSearchResults Retrieves matching symbols for query
func GetSearchResults(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id").(uint)
	urlVars := mux.Vars(req)

	user, err := GetUserInfo(contextUserID)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem finding user",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := GetQuestradeClient(user)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem retrieving Questrade Client. Invalid Refresh Token or Problem Updating Refresh Token.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	results, err := client.SearchSymbols(urlVars["searchterm"], 0)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error finding matches for query.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	res.WriteHeader(200)
	var resp = map[string]interface{}{
		"status":     200,
		"statusText": "ok",
		"message":    "Successfully retrieved Live PL for positions",
		"data":       results,
	}
	json.NewEncoder(res).Encode(resp)
}

// GetOptionsChain Retrieves Option Chain if it exists for a position
func GetOptionsChain(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id").(uint)
	urlVars := mux.Vars(req)

	user, err := GetUserInfo(contextUserID)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem finding user",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	client, err := GetQuestradeClient(user)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Problem retrieving Questrade Client. Invalid Refresh Token or Problem Updating Refresh Token.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	questradeID, err := strconv.ParseInt(urlVars["questradeID"], 10, 64)
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     400,
			"statusText": "INVALID_REQUEST_ERROR",
			"message":    "Error converting questrade ID.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	results, err := client.GetOptionChain(int(questradeID))
	if err != nil {
		res.WriteHeader(500)
		var resp = map[string]interface{}{
			"status":     500,
			"statusText": "INTERNAL_SERVER_ERROR",
			"message":    "Error retrieving options chain from Questrade.",
		}

		fmt.Println(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	res.WriteHeader(200)
	var resp = map[string]interface{}{
		"status":     200,
		"statusText": "ok",
		"message":    "Successfully retrieved Live PL for positions",
		"data":       results,
	}
	json.NewEncoder(res).Encode(resp)
}

func PositionInformation(res http.ResponseWriter, req *http.Request) {

}

// GetQuestradeClient Client for API Requests and Updating Refresh Token
// Retrieves client to make API request to Questrade and Updates Refresh Token in Database
func GetQuestradeClient(user *models.User) (*qapi.Client, error) {
	client, err := qapi.NewClient(user.RefreshToken, false, false)
	if err != nil {
		return nil, err
	}

	newRefreshToken := client.Credentials.RefreshToken
	if err := db.Model(&user).Table("users").Update("refresh_token", newRefreshToken).Error; err != nil {
		return nil, err
	}

	return client, nil
}
