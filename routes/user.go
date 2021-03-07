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
	UserRepo := repository.NewUserRepository(db)

	UserService := services.NewUserService(UserRepo)
	AuthService := config.NewAuthService()

	UserHandler := handler.UserHandlerInit(UserService, AuthService)

	api.POST("/register", UserHandler.Register)
	api.POST("/login", UserHandler.Login)
	api.POST("/email-avaiable", UserHandler.IsEmailAvaiable)
	api.POST("/upload-avatar", middleware.AuthMiddleware(AuthService, UserService), UserHandler.UploadAvatar)
}
