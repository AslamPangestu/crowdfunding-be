package main

import (
	"crowdfunding/config"
	"crowdfunding/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Failure, %v\n", err.Error())
	}

	//Initialize DB
	db := config.NewDB()

	//ROUTING
	router := gin.Default()
	router.Use(config.NewCORS())
	//Static Routing
	router.Static("/statics/avatars", "./storage/avatars")
	router.Static("/statics/campaigns", "./storage/campaigns")
	//APIV1 Routing
	apiV1 := router.Group("/api/v1")
	routes.RoleRoute(apiV1, db)
	routes.UserRoute(apiV1, db)
	routes.CampaignRoute(apiV1, db)
	routes.TransactionRoute(apiV1, db)
	router.Run()
}
