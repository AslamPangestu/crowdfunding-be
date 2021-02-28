package routes

import (
	"crowdfunding/handler"
	"crowdfunding/repository"
	"crowdfunding/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RoleRoute : Role Routing
func RoleRoute(api *gin.RouterGroup, db *gorm.DB) {
	repository := repository.RoleRepositoryInit(db)
	service := services.RoleServiceInit(repository)
	handler := handler.RoleHandlerInit(service)
	api.POST("/roles", handler.Create)
}
