package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	UserID       uint    `gorm:"primaryKey;autoIncrement"`
	FirstName    string  `gorm:"varchar(100)" json:"firstname"`
	LastName     string  `gorm:"varchar(100)" json:"lastname"`
	Email        *string `gorm:"varchar(255);unique_index" json:"email"`
	Password     string  `json:"password"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	RefreshToken string         `gorm:"varchar(255)" json:"refresh_token"`
}

// Login Token Type
type LoginToken struct {
	UserID uint
	Name   string
	Email  string
	jwt.StandardClaims
}

type Accounts struct {
	AccountType   string
	AccountNumber string
	AccountStatus string
}

type ErrorResponse struct {
	Err string
}

type QuestradeAccount struct {
	AccountID   uint `gorm:"primaryKey"`
	UserID      uint
	AccountType string `gorm:"varchar(50)"`
	Status      bool   `gorm:"bool"`
}

type Exception struct {
	ErrorCode int
	Message   string
}
