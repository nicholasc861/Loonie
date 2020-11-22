package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)

}

func ConnectDB() *gorm.DB {
	username := GetEnv("databaseUser")
	password := GetEnv("databasePassword")
	databaseName := GetEnv("databaseName")
	databaseHost := GetEnv("databaseHost")
	databasePort := GetEnv("databasePort")
	fmt.Println(username)

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, databasePort, username, databaseName, password)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
	return db

}
