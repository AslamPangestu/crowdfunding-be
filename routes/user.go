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

	Handler := handler.UserHandlerInit(UserService, AuthService)

	api.POST("/register", Handler.Register)
	api.POST("/login", Handler.Login)
	api.POST("/email-validate", Handler.IsEmailAvaiable)
	api.POST("/upload-avatar", middleware.AuthMiddleware(AuthService, UserService), Handler.UploadAvatar)
}
