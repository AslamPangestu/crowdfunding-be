package routes

import (
	"crowdfunding/config"
	handler "crowdfunding/handler/api"
	"crowdfunding/middleware"
	"crowdfunding/repository"
	"crowdfunding/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// APIRoute : List API Routing
func APIRoute(api *gin.RouterGroup, db *gorm.DB) {
	//REPOSITORY
	userRepository := repository.NewUserRepository(db)
	campaignRepository := repository.NewCampaignRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	roleRepository := repository.NewRoleRepository(db)

	//SERVICE
	authService := config.NewAuthService()
	userService := services.NewUserService(userRepository, roleRepository)
	campaignService := services.NewCampaignService(campaignRepository)
	paymentService := services.NewPaymentService(transactionRepository, campaignRepository)
	transactionService := services.NewTransactionService(transactionRepository, campaignRepository, paymentService)

	//HANDLER
	userHandler := handler.UserHandlerInit(userService, authService)
	campaignHandler := handler.CampaignHandlerInit(campaignService)
	transactionHandler := handler.TransactionHandlerInit(transactionService, paymentService)

	//ROUTING
	//User
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/profile", middleware.APIAuthMiddleware(authService, userService), userHandler.FetchUser)
	api.PATCH("/profile", middleware.APIAuthMiddleware(authService, userService), userHandler.UpdateUser)
	api.POST("/email-validate", userHandler.IsEmailAvailable)
	api.POST("/upload-avatar", middleware.APIAuthMiddleware(authService, userService), userHandler.UploadAvatar)

	//Campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.APIAuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PATCH("/campaigns/:id", middleware.APIAuthMiddleware(authService, userService), campaignHandler.EditCampaign)
	api.POST("/campaign-images", middleware.APIAuthMiddleware(authService, userService), campaignHandler.UploadImage)

	//Transaction
	api.GET("/campaigns/:id/transactions", middleware.APIAuthMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/users/transactions", middleware.APIAuthMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/users/transactions", middleware.APIAuthMiddleware(authService, userService), transactionHandler.MakeTransaction)
	api.POST("/users/notification", transactionHandler.GetNotification)
}
