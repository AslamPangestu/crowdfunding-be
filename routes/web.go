package routes

import (
	handler "crowdfunding/handler/web"
	"crowdfunding/repository"
	"crowdfunding/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// WebRoute : List Web Routing
func WebRoute(route *gin.Engine, db *gorm.DB) {
	//REPOSITORY
	userRepository := repository.NewUserRepository(db)
	dampaignRepository := repository.NewCampaignRepository(db)

	//SERVICE
	userService := services.NewUserService(userRepository)
	campaignService := services.NewCampaignService(dampaignRepository)

	// HANDLER
	userHandler := handler.UserHandlerInit(userService)
	campaignHandler := handler.CampaignHandlerInit(campaignService, userService)

	//ROUTING
	//User
	route.GET("/users", userHandler.Index)
	route.GET("/users/create", userHandler.Create)
	route.POST("/users", userHandler.PostCreate)
	route.GET("/users/edit/:id", userHandler.Edit)
	route.POST("/users/:id/update", userHandler.PostEdit)
	route.GET("/users/avatar/:id", userHandler.UploadAvatar)
	route.POST("/users/:id/avatar", userHandler.PostUploadAvatar)

	//Campaign
	route.GET("/campaigns", campaignHandler.Index)
	route.GET("/campaigns/create", campaignHandler.Create)
	route.POST("/campaigns", campaignHandler.PostCreate)
	route.GET("/campaigns/image/:id", campaignHandler.UploadImages)
	route.POST("/campaigns/:id/image", campaignHandler.PostUploadImages)
	route.GET("/campaigns/edit/:id", campaignHandler.Edit)
	route.POST("/campaigns/:id/update", campaignHandler.PostEdit)
	route.GET("/campaigns/detail/:id", campaignHandler.Detail)
}
