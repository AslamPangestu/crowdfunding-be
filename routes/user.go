package routes

import (
	"crowdfunding/config"
	"crowdfunding/handler"
	"crowdfunding/repository"
	"crowdfunding/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRoute : User Routing
func UserRoute(api *gin.RouterGroup, db *gorm.DB) {
	repository := repository.UserRepositoryInit(db)
	service := services.UserServiceInit(repository)
	authService := config.JwtServiceInit()
	handler := handler.UserHandlerInit(service, authService)
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)
	api.POST("/email-avaiable", handler.IsEmailAvaiable)
	api.POST("/upload-avatar", handler.UploadAvatar)
}
