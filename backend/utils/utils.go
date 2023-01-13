package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	if user == "" || pass == "" {
		log.Fatalln("You must include POSTGRES_USER and POSTGRES_PASSWORD environment variables")
	}
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")
	if host == "" || port == "" || dbname == "" {
		log.Fatalln("You must include POSTGRES_HOST, POSTGRES_PORT, and POSTGRES_DB environment variables")
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)

	fmt.Println(databaseURL)
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
	return db

}
