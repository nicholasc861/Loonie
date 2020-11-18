package controllers

import (
	"net/http"

	utils "../utils"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Protected   bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var db = utils.ConnectDB()

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	user := router.PathPrefix("/user").Subrouter()
	user.Use(JwtVerify)

	for _, route := range routes {
		if route.Protected {
			user.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		} else {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}

	return router
}

var routes = Routes{
	Route{
		"User Registration",
		"POST",
		"/register",
		false,
		CreateUser,
	},
	Route{
		"User Login",
		"POST",
		"/login",
		false,
		Login,
	},
	// User Routes /user/{}
	Route{
		"Add Questrade Account",
		"POST",
		"/account",
		true,
		AddQuestradeAccount,
	},
	Route{
		"Add Questrade API Refresh Token",
		"POST",
		"/addtoken",
		true,
		AddRefreshToken,
	},
}
