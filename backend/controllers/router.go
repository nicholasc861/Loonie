package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicholasc861/questrack-backend/utils"
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
	Route{
		"Get All Questrade Positions for All Accounts",
		"GET",
		"/positions",
		true,
		GetAllQuestradePositions,
	},
	Route{
		"Get All Questrade Balances for All Accounts",
		"GET",
		"/balances",
		true,
		GetAllQuestradeBalances,
	},
	Route{
		"Get Quote for a Position ID",
		"GET",
		"/hisquote/{positionID}",
		true,
		GetHistoricalQuote,
	},
	Route{
		"Get Current Quote for a Position ID",
		"GET",
		"/quote/{positionID}",
		true,
		GetCurrentQuote,
	},
	Route{
		"Get Live P&L for a Position ID",
		"POST",
		"/livepl",
		true,
		GetLivePL,
	},
	Route{
		"Get Options Chain For Position ID",
		"GET",
		"/option/{questradeID}",
		true,
		GetOptionsChain,
	},
	Route{
		"Get Query Results",
		"GET",
		"/query/{searchterm}",
		true,
		GetSearchResults,
	},
}
