package handler

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"encoding/json"
	"net/http"

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
	if err != nil || model.RoleID != "" {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	stringify, err := json.Marshal(model)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	session := sessions.Default(c)
	session.Set("userID", model.ID)
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

func (h *authHandler) Profile(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	form := entity.EditUserForm{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		User:     user,
	}
	c.HTML(http.StatusOK, "profile.html", form)
}

func (h *authHandler) PostProfile(c *gin.Context) {
	id := c.Param("id")
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	var form entity.EditUserForm
	err = c.ShouldBind(&form)
	if err != nil {
		form.Error = err
		form.User = user
		c.HTML(http.StatusOK, "profile.html", form)
		return
	}
	//Update Profile
	form.ID = id
	model, err := h.service.UpdateUser(form)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	//Update Cache
	stringify, err := json.Marshal(model)
	if err != nil {
		form.Error = err
		form.User = user
		c.HTML(http.StatusOK, "profile.html", form)
		return
	}
	session := sessions.Default(c)
	session.Set("user", stringify)
	session.Save()
	c.Redirect(http.StatusFound, "/users")
}
