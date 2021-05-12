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
	UserRepository := repository.NewUserRepository(db)

	//SERVICE
	UserService := services.NewUserService(UserRepository)

	// HANDLER
	UserHandler := handler.UserHandlerInit(UserService)

	//ROUTING
	//User
	route.GET("/users", UserHandler.Index)
	route.GET("/users/create", UserHandler.Create)
	route.POST("/users", UserHandler.PostCreate)
	route.GET("/users/edit/:id", UserHandler.Edit)
	route.POST("/users/:id/update", UserHandler.PostEdit)
	route.GET("/users/avatar/:id", UserHandler.UploadAvatar)
	route.POST("/users/:id/avatar", UserHandler.PostUploadAvatar)
}
