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
	UserRepository := repository.NewUserRepository(db)
	CampaignRepository := repository.NewCampaignRepository(db)
	TransactionRepo := repository.NewTransactionRepository(db)

	//SERVCE
	AuthService := config.NewAuthService()
	UserService := services.NewUserService(UserRepository)
	CampaignService := services.NewCampaignService(CampaignRepository)
	PaymentService := services.NewPaymentService(TransactionRepo, CampaignRepository)
	TransactionService := services.NewTransactionService(TransactionRepo, CampaignRepository, PaymentService)

	//HANDLER
	UserHandler := handler.UserHandlerInit(UserService, AuthService)
	CampaignHandler := handler.CampaignHandlerInit(CampaignService)
	TransactionHandler := handler.TransactionHandlerInit(TransactionService, PaymentService)

	//ROUTING
	//User
	api.POST("/register", UserHandler.Register)
	api.POST("/login", UserHandler.Login)
	api.GET("/profile", middleware.AuthMiddleware(AuthService, UserService), UserHandler.FetchUser)
	api.POST("/email-validate", UserHandler.IsEmailAvaiable)
	api.POST("/upload-avatar", middleware.AuthMiddleware(AuthService, UserService), UserHandler.UploadAvatar)

	//Campaign
	api.GET("/campaigns", CampaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", CampaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(AuthService, UserService), CampaignHandler.CreateCampaign)
	api.PATCH("/campaigns/:id", middleware.AuthMiddleware(AuthService, UserService), CampaignHandler.EditCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(AuthService, UserService), CampaignHandler.UploadImage)

	//Transaction
	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.GetCamapaignTransactions)
	api.GET("/users/transactions", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.GetUserTransactions)
	api.POST("/users/transactions", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.MakeTransaction)
	api.GET("/users/notification", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.MakeTransaction)
}
