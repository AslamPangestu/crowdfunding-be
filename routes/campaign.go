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

// CampaignRoute : Campaign Routing
func CampaignRoute(api *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	repository := repository.NewCampaignRepository(db)

	service := services.CampaignServiceInit(repository)
	userService := services.UserServiceInit(userRepository)
	authService := config.AuthServiceInit()

	handler := handler.CampaignHandlerInit(service)

	api.GET("/campaigns", handler.GetCampaigns)
	api.GET("/campaigns/:id", handler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(authService, userService), handler.CreateCampaign)
}
