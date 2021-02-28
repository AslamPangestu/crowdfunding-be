package handler

import (
	"crowdfunding/config"
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     services.UserService
	authService config.JwtService
}

// UserHandlerInit Initiation
func UserHandlerInit(service services.UserService, authService config.JwtService) *userHandler {
	return &userHandler{service, authService}
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
		errResponse := helper.ResponseHandler("Register Failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		errResponse := helper.ResponseHandler("Register Failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
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

	userLogged, err := h.service.Login(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("Login Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	token, err := h.authService.GenerateToken(userLogged.ID)
	if err != nil {
		errResponse := helper.ResponseHandler("Login Failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	data := entity.LoginAdapter(userLogged, token)
	res := helper.ResponseHandler("Login Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) IsEmailAvaiable(c *gin.Context) {
	var request entity.EmailValidationRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("IsEmailAvaiable Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	isEmailAvaiable, err := h.service.IsEmailAvaiable(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("IsEmailAvaiable Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	data := gin.H{
		"is_avaiable": isEmailAvaiable,
	}
	responseMessage := "Email already register"
	if isEmailAvaiable {
		responseMessage = "Email is avaiable"
	}
	res := helper.ResponseHandler(responseMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	//Get File from Storage
	file, err := c.FormFile("avatar")
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		errResponse := helper.ResponseHandler("UploadAvatar Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//Store File to Storage
	path := "storage/avatars" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		errResponse := helper.ResponseHandler("UploadAvatar Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//Save Filename to DB
	_, err = h.service.UploadAvatar(1, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		errResponse := helper.ResponseHandler("UploadAvatar Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	data := gin.H{"is_uploaded": true}
	res := helper.ResponseHandler("UploadAvatar Success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
