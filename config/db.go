package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB : Initialize DB
func NewDB() *gorm.DB {
	//GET CONFIG
	var USERNAME = os.Getenv("DB_USERNAME")
	var PASSWORD = os.Getenv("DB_PASSWORD")
	var HOST = os.Getenv("DB_HOST")
	var PORT = os.Getenv("DB_PORT")
	var DBNAME = os.Getenv("DB_NAME")
	//Setup

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", HOST, USERNAME, PASSWORD, DBNAME, PORT)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB Connection Failure, %v\n", err.Error())
	}
	fmt.Println("DB Connected")
	return db
}
