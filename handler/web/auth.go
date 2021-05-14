package handler

import (
	"crowdfunding/entity"
	"crowdfunding/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service services.UserInteractor
}

func AuthHandlerInit(service services.UserInteractor) *authHandler {
	return &authHandler{service}
}

func (h *authHandler) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *authHandler) PostLogin(c *gin.Context) {
	var form entity.LoginRequest
	err := c.ShouldBind(&form)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	model, err := h.service.Login(form)
	if err != nil || model.RoleID != 1 {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	stringify, err := json.Marshal(model)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	session := sessions.Default(c)
	session.Set("userID", strconv.Itoa(model.ID))
	session.Set("user", stringify)
	session.Save()

	c.Redirect(http.StatusFound, "/users")
}

func (h *authHandler) PostLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
