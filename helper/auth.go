package helper

import (
	"crowdfunding/entity"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserLoggedIn(c *gin.Context) (entity.User, error) {
	session := sessions.Default(c)
	var user entity.User
	userSession := session.Get("user")
	if userSession == nil {
		c.Redirect(http.StatusFound, "/login")
		return entity.User{}, errors.New("User Not Logged In")
	}
	err := json.Unmarshal(userSession.([]byte), &user)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return entity.User{}, err
	}
	return user, nil
}
