package main

import (
	"log"
	"net/http"

	controllers "./controllers"
	"github.com/rs/cors"
)

func main() {
	router := controllers.NewRouter()

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
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

	log.Fatal(http.ListenAndServe(":8080", handler))

}
