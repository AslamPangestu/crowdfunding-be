package routes

import (
	"crowdfunding/config"
	"crowdfunding/handler"
	"crowdfunding/middleware"
	"crowdfunding/repository"
	"crowdfunding/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRoute : User Routing
func UserRoute(api *gin.RouterGroup, db *gorm.DB) {
	repository := repository.UserRepositoryInit(db)
	userService := services.UserServiceInit(repository)
	authService := config.AuthServiceInit()
	handler := handler.UserHandlerInit(userService, authService)
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)
	api.POST("/email-avaiable", handler.IsEmailAvaiable)
	api.POST("/upload-avatar", middleware.AuthMiddleware(authService, userService), handler.UploadAvatar)
}
