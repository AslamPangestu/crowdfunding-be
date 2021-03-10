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

// TransactionRoute : Transaction Routing
func TransactionRoute(api *gin.RouterGroup, db *gorm.DB) {
	UserRepo := repository.NewUserRepository(db)
	CampaignRepo := repository.NewCampaignRepository(db)
	TransactionRepo := repository.NewTransactionRepository(db)

	PaymentService := services.NewPaymentService(TransactionRepo, CampaignRepo)
	UserService := services.NewUserService(UserRepo)
	AuthService := config.NewAuthService()
	TransactionService := services.NewTransactionService(TransactionRepo, CampaignRepo, PaymentService)

	TransactionHandler := handler.TransactionHandlerInit(TransactionService, PaymentService)

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.GetCamapaignTransactions)
	api.GET("/users/transactions", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.GetUserTransactions)
	api.POST("/users/transactions", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.MakeTransaction)
	api.GET("/users/notification", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.MakeTransaction)
}
