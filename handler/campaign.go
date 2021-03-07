package handler

import (
	"crowdfunding/adapter"
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service services.CampaignInteractor
}

// CampaignHandlerInit Initiation
func CampaignHandlerInit(service services.CampaignInteractor) *campaignHandler {
	return &campaignHandler{service}
}

/**
ROUTE: api/v1/campaigns
METHOD: POST
*/
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	var request entity.CreateCampaignRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("CreateCampaign Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	request.CampaignerID = currentUser.ID
	newCampaign, err := h.service.CreateCampaign(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("CreateCampaign Failed Created", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	data := adapter.CampaignAdapter(newCampaign)
	res := helper.ResponseHandler("CreateCampaign Successful Created", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/campaigns
METHOD: GET
*/
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCampaigns Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	data := adapter.CampaignsAdapter(campaigns)
	res := helper.ResponseHandler("GetCampaigns Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/campaigns/:id
METHOD: GET
*/
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var request entity.CampaignDetailRequest

	//GET ID CAMPAIGN
	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCampaign Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GET CAMPAIGN
	campaign, err := h.service.GetCampaignByID(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCampaign Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	data := adapter.CampaignDetailAdapter(campaign)
	res := helper.ResponseHandler("GetCampaign Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/campaigns
METHOD: PATCH
*/
func (h *campaignHandler) EditCampaign(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	var requestID entity.CampaignDetailRequest

	//GET ID CAMPAIGN
	err := c.ShouldBindUri(&requestID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("EditCampaign Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	var request entity.CreateCampaignRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("EditCampaign Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	request.CampaignerID = currentUser.ID
	updateCampaign, err := h.service.EditCampaign(requestID, request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("EditCampaign Failed Created", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	data := adapter.CampaignAdapter(updateCampaign)
	res := helper.ResponseHandler("EditCampaign Successful Created", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/campaigns
METHOD: POST
*/
func (h *campaignHandler) UploadImage(c *gin.Context) {
	var request entity.UploadCampaignImageRequest
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)

	err := c.ShouldBind(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("Get UploadImage Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//Get File from Storage
	file, err := c.FormFile("file")
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
		errResponse := helper.ResponseHandler("UploadImage Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	//Store File to Storage
	filename := fmt.Sprintf("%d-%s.jpg", currentUser.ID, file.Filename)
	path := "storage/campaigns/" + filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
		errResponse := helper.ResponseHandler("Store UploadImage Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	request.UserID = currentUser.ID
	//Save Filename to DB
	_, err = h.service.UploadCampaignImages(request, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
		errResponse := helper.ResponseHandler("UploadImage Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	data := gin.H{"is_uploaded": true}
	res := helper.ResponseHandler("UploadImage Success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
