package handler

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	service services.RoleInteractor
}

// NewRoleHandler Initiation
func NewRoleHandler(service services.RoleInteractor) *roleHandler {
	return &roleHandler{service}
}

/**
ROUTE: api/v1/roles
METHOD: POST
*/
func (h *roleHandler) AddRole(c *gin.Context) {
	//GET REQUEST ROLE
	var request entity.RoleRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Role Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//SAVE ROLE DB
	role, err := h.service.AddRole(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("AddRole Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	res := helper.ResponseHandler("AddRole Successful", http.StatusOK, "success", role)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/roles
METHOD: GET
*/
func (h *roleHandler) GetRoles(c *gin.Context) {
	//GET ROLE DB
	roles, err := h.service.GetRoles()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetRoles Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	res := helper.ResponseHandler("GetRoles Successful", http.StatusOK, "success", roles)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/roles/:ID
METHOD: GET
*/
func (h *roleHandler) GetRoleByID(c *gin.Context) {
	var uri entity.RoleIdRequest
	//GET ID ROLE
	err := c.ShouldBindUri(&uri)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetRoleID Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GET ROLE DB
	roles, err := h.service.GetRoleByID(uri)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetRoleByID Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	res := helper.ResponseHandler("GetRoleByID Successful", http.StatusOK, "success", roles)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/roles?name=
METHOD: GET
*/
func (h *roleHandler) GetRolesByName(c *gin.Context) {
	var request entity.RoleNameRequest
	//GET ID ROLE
	err := c.ShouldBindQuery(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetRoleName Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GET ROLE DB
	roles, err := h.service.GetRolesByName(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetRoleByID Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	res := helper.ResponseHandler("GetRoleByID Successful", http.StatusOK, "success", roles)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/roles
METHOD: PATCH
*/
func (h *roleHandler) EditRole(c *gin.Context) {
	var uri entity.RoleIdRequest
	var request entity.RoleRequest
	//GET ID ROLE
	err := c.ShouldBindUri(&uri)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetRoleID Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GET REQUEST ROLE
	err = c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Role Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//UPDATE ROLE DB
	role, err := h.service.EditRole(uri, request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("EditRole Failed Edit", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	res := helper.ResponseHandler("EditRole Successful", http.StatusOK, "success", role)
	c.JSON(http.StatusOK, res)
}
