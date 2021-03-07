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
	transactionRepository := repository.NewTransactionRepository(db)
	userRepository := repository.NewUserRepository(db)
	campaignRepository := repository.NewCampaignRepository(db)

	service := services.NewTransactionService(transactionRepository, campaignRepository)
	userService := services.NewUserService(userRepository)
	authService := config.AuthServiceInit()

	handler := handler.TransactionHandlerInit(service)

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(authService, userService), handler.GetCamapaignTransactions)
}
