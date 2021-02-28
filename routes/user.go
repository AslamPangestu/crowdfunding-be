package routes

import (
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
	handler := handler.UserHandlerInit(service)
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)
}
