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

func CreateUser(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	json.NewDecoder(req.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := models.ErrorResponse{
			Err: "Password Encryption failed.",
		}
		json.NewEncoder(res).Encode(err)
	}

	user.UserID = 1
	user.Password = string(pass)
	createdUser := db.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}

	json.NewEncoder(res).Encode(createdUser)
}

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

	resp, JWTCookie := FindOne(user.Email, user.Password)

	http.SetCookie(res, JWTCookie)
	json.NewEncoder(res).Encode(resp)
}

func FindOne(email, password string) (map[string]interface{}, *http.Cookie) {
	user := &models.User{}

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Email or password not found",
		}
		return resp, nil
	}

	expiresAt := time.Now().Add(time.Minute * 720) // JWT is valid for 12 hours

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{
			"status":  false,
			"message": "Invalid login credentials",
		}
		return resp, nil
	}

	tk := &models.LoginToken{
		UserID: user.UserID,
		Name:   user.FirstName,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    "Questrack",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}

	var resp = map[string]interface{}{
		"status":  true,
		"message": "Logged in successfully",
	}

	JWTCookie := &http.Cookie{
		Name:     "access-token",
		Value:    tokenString,
		Expires:  expiresAt,
		HttpOnly: true,
	}

	return resp, JWTCookie
}