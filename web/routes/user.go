package routes

import (
	"crowdfunding/repository"
	"crowdfunding/services"
	"crowdfunding/web/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRoute : User Routing
func UserRoute(route *gin.Engine, db *gorm.DB) {
	UserRepo := repository.NewUserRepository(db)

	UserService := services.NewUserService(UserRepo)

	Handler := handler.UserHandlerInit(UserService)

	route.GET("/users", Handler.Index)
	route.GET("/users/create", Handler.Create)
	route.POST("/users", Handler.PostCreate)
	route.GET("/users/edit/:id", Handler.Edit)
	route.POST("/users/:id/update", Handler.PostEdit)
	route.GET("/users/avatar/:id", Handler.UploadAvatar)
	route.POST("/users/:id/avatar", Handler.PostUploadAvatar)
}
