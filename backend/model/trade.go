package model

import "time"

type Side int
type EquityType int
type OptionType int

const (
	Buy Side = iota
	Sell
)

const (
	Stock EquityType = iota
	Option
	Crypto
)

const (
	Call OptionType = iota
	Put
)

// Trade
type TradeGroup struct {
	ID        string     `gorm:"primary_key" json:"id"`
	AccountID string     `json:"account_id"`
	Type      EquityType `json:"type"`
}

type StockInfo struct {
	ID          string `gorm:"primary_key" json:"id"`
	Ticker      string `json:"ticker"`
	Description string `json:"description"`
}

type OptionInfo struct {
	ID          string     `gorm:"primary_key" json:"id"`
	Ticker      string     `json:"ticker"`
	Description string     `json:"description"`
	Expiration  time.Time  `json:"expiration"`
	Strike      uint       `json:"strike"`
	Type        OptionType `json:"type"`
}

type CryptoInfo struct {
	ID          string `gorm:"primary_key" json:"id"`
	Ticker      string `json:"ticker"`
	Description string `json:"description"`
}

type Trade struct {
	ID           string    `gorm:"primary_key" json:"id"`
	TradeGroupID string    `json:"trade_id"`
	Datetime     time.Time `json:"date"`
	Price        uint      `json:"price"`
	Quantity     uint      `json:"quantity"`
	Commission   uint      `json:"commission"`
	GrossAmount  uint      `json:"gross_amount"`
	NetAmount    uint      `json:"net_amount"`
	Side         Side      `json:"side"`
}

type NewTrade struct {
	Trade     Trade      `json:"trade"`
	AccountID string     `json:"account_id"`
	Type      EquityType `json:"type"`
}

type TradeRequest struct {
	TradeGroup TradeGroup `json:"trade_group"`
	Trade      Trade      `json:"trade"`
}

func (TradeGroup) TableName() string {
	return "trade_group"
}

func (Trade) TableName() string {
	return "trade"
}
