package helper

import (
	"crowdfunding/entity"
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserLoggedIn(c *gin.Context) entity.User {
	session := sessions.Default(c)
	var user entity.User
	userSession := session.Get("user")
	json.Unmarshal(userSession.([]byte), &user)
	return user
}
