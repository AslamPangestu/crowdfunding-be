package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CampaignRoute : Campaign Routing
func CampaignRoute(api *gin.RouterGroup, db *gorm.DB) {
	// repository := repository.RoleRepositoryInit(db)
	// service := services.RoleServiceInit(repository)
	// handler := handler.RoleHandlerInit(service)
	// api.POST("/campaigns", handler.Create)
}
