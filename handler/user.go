package handler

import (
	"crowdfunding/adapter"
	"crowdfunding/config"
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     services.UserInteractor
	authService config.AuthService
}

// UserHandlerInit Initiation
func UserHandlerInit(service services.UserInteractor, authService config.AuthService) *userHandler {
	return &userHandler{service, authService}
}

/**
ROUTE: api/v1/register
METHOD: POST
*/
func (h *userHandler) Register(c *gin.Context) {
	//GET REQUEST REGISTER
	var request entity.RegisterRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Register Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//SAVE USER DB
	user, err := h.service.Register(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("Register Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GENERATE TOKEN
	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GenerateToken Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := adapter.RegsiterAdapter(user, token)
	res := helper.ResponseHandler("Register Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/login
METHOD: POST
*/
func (h *userHandler) Login(c *gin.Context) {
	//GET REQUEST LOGIN
	var request entity.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Login Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//GET USER LOGGED
	userLogged, err := h.service.Login(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("Login Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GENERATE TOKEN
	token, err := h.authService.GenerateToken(userLogged.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GenerateToken Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := adapter.LoginAdapter(userLogged, token)
	res := helper.ResponseHandler("Login Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/email-validate
METHOD: POST
*/
func (h *userHandler) IsEmailAvaiable(c *gin.Context) {
	//GET REQUEST EMAIL VALIDATE
	var request entity.EmailValidationRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("EmailValidate Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}

	//GET EMAIL STATUS VALIDATE
	isEmailAvaiable, err := h.service.IsEmailAvaiable(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("IsEmailAvaiable Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := gin.H{"is_avaiable": isEmailAvaiable}
	responseMessage := "Email already register"
	if isEmailAvaiable {
		responseMessage = "Email is avaiable"
	}
	res := helper.ResponseHandler(responseMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/upload-avatar
METHOD: POST
*/
func (h *userHandler) UploadAvatar(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	//Get File from Storage
	file, err := c.FormFile("file")
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
		errResponse := helper.ResponseHandler("Get File Avatar Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//Store File to Storage
	filename := fmt.Sprintf("%d-%s.jpg", currentUser.ID, currentUser.Username)
	path := "storage/avatars/" + filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
		errResponse := helper.ResponseHandler("Store File Avatar Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//Save Filename to DB
	_, err = h.service.UploadAvatar(currentUser.ID, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
		errResponse := helper.ResponseHandler("UploadAvatar Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := gin.H{"is_uploaded": true}
	res := helper.ResponseHandler("UploadAvatar Success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/login
METHOD: POST
*/
func (h *userHandler) FetchUser(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	//RESPONSE
	data := adapter.LoginAdapter(currentUser, "")
	res := helper.ResponseHandler("FetchUser Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
