package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	models "../models"
	"github.com/dgrijalva/jwt-go"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookieToken, err := req.Cookie("access-token")
		if err != nil {
			json.NewEncoder(res).Encode(models.Exception{
				ErrorCode: 404,
				Message:   "Missing authentication token",
			})
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

		fmt.Printf("%v", claims.Name)

		if err != nil {
			res.WriteHeader(http.StatusForbidden)
			json.NewEncoder(res).Encode(models.Exception{
				Message: err.Error(),
			})
			return
		}

		ctx := context.WithValue(req.Context(), "user", claims)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
