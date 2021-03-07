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
	RoleRepo := repository.NewRoleRepository(db)
	RoleService := services.NewRoleService(RoleRepo)
	Handler := handler.NewRoleHandler(RoleService)

	api.GET("/roles", Handler.GetRoles)
	api.GET("/roles/:id", Handler.GetRoleByID)
	api.POST("/roles", Handler.AddRole)
	api.PATCH("/roles/:id", Handler.EditRole)
}
