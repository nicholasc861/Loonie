package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nicholasc861/Loonie/backend/model"
)

func CreateTradeGroup(res http.ResponseWriter, req *http.Request) {
	tradeInfo := model.NewTrade{}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&tradeInfo); err != nil {
		fmt.Errorf("", "")
		return
	}

	// Create trade group object containing information
	tradeGroup := model.TradeGroup{
		AccountID: tradeInfo.AccountID,
		Type:      tradeInfo.Type,
	}

	// Create new trade group in database
	db.Create(&tradeGroup)

	// Create new trade associated with the trade group
	tradeInfo.Trade.ID = tradeGroup.ID
	db.Create(&tradeInfo.Trade)

}

func CreateTrade(res http.ResponseWriter, req *http.Request) {
	trade := model.Trade{}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&trade); err != nil {
		fmt.Errorf("", "")
		return
	}

	db.Create(&trade)
}
