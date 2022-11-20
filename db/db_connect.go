package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func DbConnect() (*gorm.DB, error) {
	db_url := os.Getenv("POSTGRES_ADDRESS")
	db_name := os.Getenv("POSTGRES_DB")

	db, err := gorm.Open("postgres", fmt.Sprintf("%v/%v", db_url, db_name))

	if err != nil {
		log.Printf("Error connecting to database. \n%v", err)
	}

	return db, nil
}
