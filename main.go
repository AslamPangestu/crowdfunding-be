package main

import (
	"crowdfunding/config"
	"crowdfunding/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load ENV
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Initialize DB
	db := config.InitDB()

	//Routing
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	routes.RoleRoute(apiV1, db)
	routes.UserRoute(apiV1, db)
	router.Run()
}