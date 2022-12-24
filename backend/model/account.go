package model

import "time"

type Account struct {
	ID          string `gorm:"primary_key" json:"id"`
	ProviderID  string `json:"provider_id"`
	Description string `json:"description"`
}

type Balance struct {
	ID        string    `gorm:"primary_key" json:"id"`
	AccountID string    `json:"account_id"`
	Date      time.Time `json:"date"`
	Amount    float64   `json:"amount"`
}
