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
ROUTE: api/v1/roles
METHOD: POST
*/
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	fmt.Println(currentUser)
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
