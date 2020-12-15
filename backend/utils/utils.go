package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
	return db

}
