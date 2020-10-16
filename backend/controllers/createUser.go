package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "../models"
	utils "./utils"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var db = utils.ConnectDB()

func CreateUser(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption failed.",
		}
		json.NewEncoder(res).Encode(err)
	}

	user.Password = string(pass)
	createdUser := db.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}

	json.NewEncoder(res).Encode(createdUser)
}
