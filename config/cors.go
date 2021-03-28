package config

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCORS() gin.HandlerFunc {
	var BASE_URL = os.Getenv("BASE_URL")
	return cors.New(cors.Config{
		AllowOrigins: []string{BASE_URL},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	})
}
