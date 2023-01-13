package controllers

import (
	"net/http"
	"os"
)

var QTRADE_MANUAL_TOKEN = os.Getenv("QTRADE_MANUAL_TOKEN")
var AUTH_URL = "https://login.questrade.com/oauth2/token?grant_type=refresh_token&refresh_token=" + QTRADE_MANUAL_TOKEN

func CheckAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("access-token")
		if err != nil {

		}

		valid := CheckCookie(cookie)
		if !valid {

		}

	})
}

func GenerateAccessToken()

func CheckCookie(cookie *http.Cookie) (valid bool) {
	if cookie.MaxAge < 0 {
		return false
	}

	return true
}
