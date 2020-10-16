package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	models "../models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Invalid Request",
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(res).Encode(resp)
}

func FindOne(email, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Email or password not found",
		}
		return resp
	}

	expiresAt := time.Now().add(time.Minute * 10000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Invalid login credentials",
		}
		return resp
	}

	tk := &models.LoginToken{
		user.ID,
		user.FirstName,
		user.Email,
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}

	var resp = map[string]interface{}{
		"status":  false,
		"message": "Logged in successfully",
	}

	resp["token"] = tokenString
	resp["user"] = user
	return resp
}
