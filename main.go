package main

import (
	"crowdfunding/config"
	"crowdfunding/routes"
	webRoutes "crowdfunding/web/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load ENV
	var APP_ENV = os.Getenv("APP_ENV")
	if APP_ENV != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("ENV Failure, %v\n", err.Error())
		}
	}

	//Initialize DB
	db := config.NewDB()

	//ROUTING
	router := gin.Default()
	router.Use(config.NewCORS())
	//Static Routing
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")
	router.Static("/statics/avatars", "./storage/avatars")
	router.Static("/statics/campaigns", "./storage/campaigns")
	//CMS Routing
	router.HTMLRender = config.LoadTemplates("./web/views")
	webRoutes.UserRoute(router, db)
	//APIV1 Routing
	apiV1 := router.Group("/api/v1")
	routes.RoleRoute(apiV1, db)
	routes.UserRoute(apiV1, db)
	routes.CampaignRoute(apiV1, db)
	routes.TransactionRoute(apiV1, db)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong")
	})

	router.Run()
}
