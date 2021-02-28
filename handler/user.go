package handler

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service services.UserService
}

// UserHandlerInit Initiation
func UserHandlerInit(service services.UserService) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) Register(c *gin.Context) {
	var request entity.RegisterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Register Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	user, err := h.service.Register(request)
	if err != nil {
		errResponse := helper.ResponseHandler("Register Failed", http.StatusUnprocessableEntity, "failed", err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	token := "TEST"
	data := entity.RegsiterAdapter(user, token)
	res := helper.ResponseHandler("User Successful Register", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) Login(c *gin.Context) {
	var request entity.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Login Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	logged, err := h.service.Login(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("Login Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	token := "TEST"
	data := entity.LoginAdapter(logged, token)
	res := helper.ResponseHandler("Login Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
