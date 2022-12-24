package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nicholasc861/Loonie/backend/controllers"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("GO_PORT")
	router := controllers.NewRouter()

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5000"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",
		},
		AllowCredentials: true,
	}).Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
