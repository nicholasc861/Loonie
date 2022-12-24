package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicholasc861/Loonie/backend/utils"
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

	return router
}

var routes = Routes{
	Route{
		"Add Questrade API Refresh Token",
		"POST",
		"/addtoken",
		true,
		AddRefreshToken,
	},
	Route{
		"Get Query Results",
		"GET",
		"/query/{searchterm}",
		true,
		GetSearchResults,
	},
}
