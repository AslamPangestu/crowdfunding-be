package config

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCORS() gin.HandlerFunc {
	var FRONTEND_URL = os.Getenv("FRONTEND_URL")
	return cors.New(cors.Config{
		AllowOrigins: []string{FRONTEND_URL},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	})
}
