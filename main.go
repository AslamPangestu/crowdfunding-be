package main

import (
	"crowdfunding/config"
	"crowdfunding/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	cookiesStore := cookie.NewStore([]byte(config.SECRET_KEY))
	router.Use(sessions.Sessions("crowdfunding", cookiesStore))

	//Testing Route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong")
	})

	//Static Routing
	router.Static("/css", "./assets/css")
	router.Static("/js", "./assets/js")
	router.Static("/webfonts", "./assets/webfonts")
	router.Static("/statics/avatars", "./storage/avatars")
	router.Static("/statics/campaigns", "./storage/campaigns")

	//CMS Routing
	router.HTMLRender = config.LoadTemplates("./views")
	routes.WebRoute(router, db)

	//APIV1 Routing
	apiV1 := router.Group("/api/v1")
	routes.APIRoute(apiV1, db)

	router.Run()
}
