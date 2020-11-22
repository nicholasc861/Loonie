package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	models "../models"
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

func AddAccount(res http.ResponseWriter, req *http.Request) {
	account := &models.QuestradeAccount{}

	err := json.NewDecoder(req.Body).Decode(account)

	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Invalid Request Format",
		}
		log.Fatal(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("user_accounts").Create(account).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error creating record in Database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	json.NewEncoder(res).Encode(account)
}

func AddMultipleAccounts(res http.ResponseWriter, req *http.Request) {
	accounts := []models.QuestradeAccount{}

	err := json.NewDecoder(req.Body).Decode(&accounts)

	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Invalid Request Format",
		}
		log.Fatal(err)
		json.NewEncoder(res).Encode(resp)
		return
	}

	if err := db.Table("user_accounts").Create(accounts).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error creating record in Database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	json.NewEncoder(res).Encode(accounts)
}

func GetQuestradeAccounts(res http.ResponseWriter, req *http.Request) {
	contextUserID := req.Context().Value("user_id")
	accounts := []models.QuestradeAccount{}

	if err := db.Table("user_accounts").Find(&accounts).Where("user_id = ?", contextUserID.(uint)).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Error creating record in Database",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	json.NewEncoder(res).Encode(accounts)
}
