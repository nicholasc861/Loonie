package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nicholasc861/Loonie/backend/model"
)

func AddAccount(res http.ResponseWriter, req *http.Request) {
	account := model.Account{}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&account); err != nil {
		return
	}

	// Create new account in database
	db.Create(&account)
}

func DeleteAccount(res http.ResponseWriter, req *http.Request) {

}

func GetAccounts(res http.ResponseWriter, req *http.Request) {
	accounts := []model.Account{}

	// Retrive all accounts
	db.Find(&accounts)

}
