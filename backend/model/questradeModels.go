package model

type QuestradeAccount struct {
	AccountID   uint `gorm:"primaryKey"`
	UserID      uint
	AccountType string `gorm:"varchar(50)"`
	Status      bool   `gorm:"bool"`
}

type QuestradePosition struct {
	PositionID        uint `gorm:"primaryKey;autoIncrement"`
	AccountID         uint
	QuestradeID       uint
	Symbol            string
	OpenQuantity      float32
	ClosedQuantity    float32
	AverageEntryPrice float32
	ClosedPNL         float32
	TotalEntry        float32
	IsOption          bool
	Status            bool
}

type QuestradeQuote struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	QuestradeID       uint
	Description       string
	Symbol            string
	BidPrice          float32
	LastTradePrice    float32
	TimeQuoted        int64
	OpenPrice         float32
	PrevDayClosePrice float32
}

type QuestradeLivePL struct {
	QuestradeID       uint    `json:"questrade_id"`
	AverageEntryPrice float32 `json:"average_entry_price"`
	OpenQuantity      float32 `json:"open_quantity"`
}
