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
	TransactionRepo := repository.NewTransactionRepository(db)
	UserRepo := repository.NewUserRepository(db)
	CampaignRepo := repository.NewCampaignRepository(db)

	TransactionService := services.NewTransactionService(TransactionRepo, CampaignRepo)
	UserService := services.NewUserService(UserRepo)
	AuthService := config.NewAuthService()

	TransactionHandler := handler.TransactionHandlerInit(TransactionService)

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(AuthService, UserService), TransactionHandler.GetCamapaignTransactions)
}
