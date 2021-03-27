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
	UserRepository := repository.NewUserRepository(db)
	CampaignRepo := repository.NewCampaignRepository(db)

	CampaignService := services.NewCampaignService(CampaignRepo)
	UserService := services.NewUserService(UserRepository)
	AuthService := config.NewAuthService()

	CampaignHandler := handler.CampaignHandlerInit(CampaignService)

	api.GET("/campaigns", CampaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", CampaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(AuthService, UserService), CampaignHandler.CreateCampaign)
	api.PATCH("/campaigns/:id", middleware.AuthMiddleware(AuthService, UserService), CampaignHandler.EditCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(AuthService, UserService), CampaignHandler.UploadImage)
}
