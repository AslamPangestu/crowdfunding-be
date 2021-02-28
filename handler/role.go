package handler

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	service services.RoleService
}

// RoleHandlerInit Initiation
func RoleHandlerInit(service services.RoleService) *roleHandler {
	return &roleHandler{service}
}

func (h *roleHandler) Create(c *gin.Context) {
	var input entity.RoleRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Role Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	role, err := h.service.Create(input)
	if err != nil {
		errResponse := helper.ResponseHandler("Role Failed Created", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	res := helper.ResponseHandler("Role Successful Created", http.StatusOK, "success", role)
	c.JSON(http.StatusOK, res)
}
