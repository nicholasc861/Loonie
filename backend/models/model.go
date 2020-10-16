package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"varchar(100)"`
	LastName  string `gorm:"varchar(100)"`
	Email     string `gorm:"varchar(100);unique_index"`
	Password  string `json:"Password"`
}

type LoginToken struct {
	UserID         uint
	Name           string
	Email          string
	StandardClaims jwt.StandardClaims
}
