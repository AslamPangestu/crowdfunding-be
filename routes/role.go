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
	repository := repository.NewRoleRepository(db)
	service := services.RoleServiceInit(repository)
	handler := handler.RoleHandlerInit(service)

	api.GET("/roles", handler.GetRoles)
	api.POST("/roles", handler.AddRole)
	api.PATCH("/roles/:id", handler.EditRole)
}
