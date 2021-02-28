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
	var input entity.RegisterRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("User Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
	}
	user, err := h.service.Register(input)
	if err != nil {
		errResponse := helper.ResponseHandler("User Failed Register", http.StatusUnprocessableEntity, "failed", err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
	}
	token := "TEST"
	data := entity.RegsiterAdapter(user, token)
	res := helper.ResponseHandler("User Successful Register", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
