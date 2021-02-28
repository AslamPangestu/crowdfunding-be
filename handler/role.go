package handler

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	service services.RoleService
}

// RoleHandlerInit Initiation
func RoleHandlerInit(service services.RoleService) *roleHandler {
	return &roleHandler{service}
}

/**
ROUTE: api/v1/roles
METHOD: POST
*/
func (h *roleHandler) AddRole(c *gin.Context) {
	var request entity.RoleRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Role Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	role, err := h.service.AddRole(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("Role Failed Created", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	res := helper.ResponseHandler("Role Successful Created", http.StatusOK, "success", role)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/roles
METHOD: GET
*/
func (h *roleHandler) GetRoles(c *gin.Context) {
	roles, err := h.service.GetRoles()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetRoles Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	res := helper.ResponseHandler("GetRoles Successful", http.StatusOK, "success", roles)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/roles
METHOD: PATCH
*/
func (h *roleHandler) EditRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var request entity.RoleRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Role Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	role, err := h.service.EditRole(id, request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("Role Failed Edit", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	res := helper.ResponseHandler("Role Successful Edit", http.StatusOK, "success", role)
	c.JSON(http.StatusOK, res)
}
