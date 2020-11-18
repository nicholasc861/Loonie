package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	models "../models"
)

func AddRefreshToken(res http.ResponseWriter, req *http.Request) {
	var test *map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &test)
	if err != nil {
		panic(err)
	}
	new := *test
	fmt.Println(new["AccountID"])

}

func AddQuestradeAccount(res http.ResponseWriter, req *http.Request) {
	account := &models.QuestradeAccount{}

	err := json.NewDecoder(req.Body).Decode(account)

	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Invalid Request",
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
