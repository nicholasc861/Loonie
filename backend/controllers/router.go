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

	// Middlewares to use
	router.Use(commonMiddleware)
	user.Use(JwtVerify)
	user.Use(commonMiddleware)

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
		AddAccount,
	},
	Route{
		"Get Questrade Accounts",
		"GET",
		"/accounts",
		true,
		GetQuestradeAccounts,
	},
	Route{
		"Add Questrade API Refresh Token",
		"POST",
		"/addtoken",
		true,
		AddRefreshToken,
	},
}
