package handler

import (
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service services.CampaignService
}

// CampaignHandlerInit Initiation
func CampaignHandlerInit(service services.CampaignService) *campaignHandler {
	return &campaignHandler{service}
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
	data := entity.CampaignsAdapter(campaigns)
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
	data := entity.CampaignDetailAdapter(campaign)
	res := helper.ResponseHandler("GetCampaign Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
