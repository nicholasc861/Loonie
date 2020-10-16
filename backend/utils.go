package utils

import (
	"auth/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(
		&models.User{}
	)

	fmt.Println("Successfully connected!", db)
	return db


}