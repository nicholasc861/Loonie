package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/nicholasc861/questrack-backend/models"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Cookies())
		fmt.Println(req)
		cookieToken, err := req.Cookie("access-token")
		if err != nil {
			res.WriteHeader(403)
			var resp = map[string]interface{}{
				"status":     403,
				"statusText": "INVALID_CREDENTIALS",
				"message":    "Please login before accessing this page",
			}

			fmt.Println(err)
			json.NewEncoder(res).Encode(resp)
			return
		}

		JWTToken := cookieToken.Value
		claims := &models.LoginToken{}

		_, err = jwt.ParseWithClaims(JWTToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if err != nil {
			res.WriteHeader(403)
			var resp = map[string]interface{}{
				"status":     403,
				"statusText": "INVALID_CREDENTIALS",
				"message":    "Please login before making requests",
			}

			fmt.Println(err)
			json.NewEncoder(res).Encode(resp)
			return
		}

		ctx := context.WithValue(req.Context(), "user_id", claims.UserID)

		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
