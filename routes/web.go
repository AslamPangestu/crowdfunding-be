package routes

import (
	handler "crowdfunding/handler/web"
	"crowdfunding/middleware"
	"crowdfunding/repository"
	"crowdfunding/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// WebRoute : List Web Routing
func WebRoute(route *gin.Engine, db *gorm.DB) {
	//REPOSITORY
	userRepository := repository.NewUserRepository(db)
	campaignRepository := repository.NewCampaignRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	roleRepository := repository.NewRoleRepository(db)

	//SERVICE
	userService := services.NewUserService(userRepository, roleRepository)
	campaignService := services.NewCampaignService(campaignRepository)
	paymentService := services.NewPaymentService(transactionRepository, campaignRepository)
	transactionService := services.NewTransactionService(transactionRepository, campaignRepository, paymentService)
	roleService := services.NewRoleService(roleRepository)

	// HANDLER
	userHandler := handler.UserHandlerInit(userService)
	campaignHandler := handler.CampaignHandlerInit(campaignService, userService)
	transactionHandler := handler.TransactionHandlerInit(transactionService)
	roleHandler := handler.RoleHandlerInit(roleService)
	authHandler := handler.AuthHandlerInit(userService)

	//ROUTING
	//User
	route.GET("/users", middleware.WebAuthMiddleware(), userHandler.Index)
	route.GET("/users/create", middleware.WebAuthMiddleware(), userHandler.Create)
	route.POST("/users", middleware.WebAuthMiddleware(), userHandler.PostCreate)
	route.GET("/users/edit/:id", middleware.WebAuthMiddleware(), userHandler.Edit)
	route.POST("/users/:id/update", middleware.WebAuthMiddleware(), userHandler.PostEdit)
	route.GET("/users/avatar/:id", middleware.WebAuthMiddleware(), userHandler.UploadAvatar)
	route.POST("/users/:id/avatar", middleware.WebAuthMiddleware(), userHandler.PostUploadAvatar)

	//Campaign
	route.GET("/campaigns", middleware.WebAuthMiddleware(), campaignHandler.Index)
	route.GET("/campaigns/create", middleware.WebAuthMiddleware(), campaignHandler.Create)
	route.POST("/campaigns", middleware.WebAuthMiddleware(), campaignHandler.PostCreate)
	route.GET("/campaigns/image/:id", middleware.WebAuthMiddleware(), campaignHandler.UploadImages)
	route.POST("/campaigns/:id/image", middleware.WebAuthMiddleware(), campaignHandler.PostUploadImages)
	route.GET("/campaigns/edit/:id", middleware.WebAuthMiddleware(), campaignHandler.Edit)
	route.POST("/campaigns/:id/update", middleware.WebAuthMiddleware(), campaignHandler.PostEdit)
	route.GET("/campaigns/detail/:id", middleware.WebAuthMiddleware(), campaignHandler.Detail)

	//Transaction
	route.GET("/transactions", middleware.WebAuthMiddleware(), transactionHandler.Index)

	//Role
	route.GET("/roles", middleware.WebAuthMiddleware(), roleHandler.Index)
	route.GET("/roles/create", middleware.WebAuthMiddleware(), roleHandler.Create)
	route.POST("/roles", middleware.WebAuthMiddleware(), roleHandler.PostCreate)
	route.GET("/roles/edit/:id", middleware.WebAuthMiddleware(), roleHandler.Edit)
	route.POST("/roles/:id/update", middleware.WebAuthMiddleware(), roleHandler.PostEdit)
	route.GET("/roles/delete/:id", middleware.WebAuthMiddleware(), roleHandler.Remove)

	//Auth
	route.GET("/login", authHandler.Login)
	route.POST("/login", authHandler.PostLogin)
	route.GET("/logout", authHandler.PostLogout)
	route.GET("/profile", middleware.WebAuthMiddleware(), authHandler.Profile)
	route.POST("/profile/:id", middleware.WebAuthMiddleware(), authHandler.PostProfile)
}
