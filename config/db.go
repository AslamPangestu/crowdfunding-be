package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//NewDB : Initialize DB
func NewDB() *gorm.DB {
	//GET CONFIG
	var USERNAME = os.Getenv("DB_USERNAME")
	var PASSWORD = os.Getenv("DB_PASSWORD")
	var HOST = os.Getenv("DB_HOST")
	var PORT = os.Getenv("DB_PORT")
	var DBNAME = os.Getenv("DB_NAME")
	//Setup
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, DBNAME)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB Connection Failure, %v\n", err.Error())
	}
	fmt.Println("DB Connected")
	return db
}
