package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			res.WriteHeader(http.StatusForbidden)
			json.NewEncoder(res).Encode(models.Exception{
				Message: err.Error(),
			})
			return
		}

		ctx := context.WithValue(req.Context(), "user_id", claims.UserID)

		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
