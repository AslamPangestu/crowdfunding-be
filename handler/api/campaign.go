package handler

import (
	"context"
	"crowdfunding/adapter"
	"crowdfunding/config"
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"fmt"
	"net/http"
	"os"
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
	//GET REQUEST CAMPAIGN
	var request entity.FormCampaignRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("CreateCampaign Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//SET CAMPAIGNER OWNER
	request.CampaignerID = currentUser.ID
	//SAVE CAMPAIGN DB
	newCampaign, err := h.service.CreateCampaign(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("CreateCampaign Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := adapter.CampaignAdapter(newCampaign)
	res := helper.ResponseHandler("CreateCampaign Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/campaigns
QUERY: user_id
METHOD: GET
*/
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	//GET CAMPAIGNS
	campaigns, err := h.service.GetCampaigns(userID, page, pageSize)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCampaigns Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := adapter.CampaignsAdapter(campaigns)
	res := helper.ResponseHandler("GetCampaigns Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/campaigns/:id
METHOD: GET
*/
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	//GET ID CAMPAIGN
	var request entity.CampaignIDRequest
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
	//RESPONSE
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
	var uri entity.CampaignIDRequest

	//GET ID CAMPAIGN
	err := c.ShouldBindUri(&uri)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("EditCampaign Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GET REQUEST CAMPAIGN
	var request entity.FormCampaignRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("EditCampaign Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//SET CAMPAIGNER OWNER
	request.CampaignerID = currentUser.ID
	//UPDATE CAMPAIGN DB
	updateCampaign, err := h.service.EditCampaign(uri, request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("EditCampaign Failed Created", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := adapter.CampaignAdapter(updateCampaign)
	res := helper.ResponseHandler("EditCampaign Successful Created", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/campaigns
METHOD: POST
*/
func (h *campaignHandler) UploadImage(c *gin.Context) {
	var path string
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)

	//GET REQUEST UPLOAD IMAGES
	var request entity.UploadCampaignImageRequest
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

	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "local" {
		//Store File to Storage
		cleanFilename := helper.RemoveFileExt(file.Filename)
		filename := fmt.Sprintf("%d-%d-%s.jpg", currentUser.ID, request.CampaignID, cleanFilename)
		path = helper.GeneratePath("campaigns", filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
			errResponse := helper.ResponseHandler("[Local]Store UploadImage Failed", http.StatusBadRequest, "failed", errorMessage)
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}
		path = helper.GenerateURL("campaigns", filename)
	}
	if storageType == "cloud" {
		var ctx = context.Background()
		cloudinary := config.NewCloudStorage()
		uploadResponse, err := cloudinary.Upload.Upload(ctx, file, config.ConfigCloudStorage("campaigns"))
		if err != nil {
			errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
			errResponse := helper.ResponseHandler("[Cloud]Store UploadImage Failed", http.StatusBadRequest, "failed", errorMessage)
			c.JSON(http.StatusBadRequest, errResponse)
			return
		}
		//Generate URL
		path = uploadResponse.SecureURL
	}
	//Save Filename to DB
	request.UserID = currentUser.ID
	_, err = h.service.UploadCampaignImages(request, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
		errResponse := helper.ResponseHandler("UploadImage Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := gin.H{"is_uploaded": true, "file": path}
	res := helper.ResponseHandler("UploadImage Success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
