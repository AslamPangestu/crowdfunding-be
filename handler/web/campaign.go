package handler

import (
	"context"
	"crowdfunding/config"
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type campaignHandler struct {
	service     services.CampaignInteractor
	userService services.UserInteractor
}

func CampaignHandlerInit(service services.CampaignInteractor, userService services.UserInteractor) *campaignHandler {
	return &campaignHandler{service, userService}
}

func (h *campaignHandler) Index(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	models, err := h.service.GetCampaigns(0, page, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	pagination := helper.PaginationAdapterHandler(models.Pagination)
	c.HTML(http.StatusOK, "campaign_index.html", gin.H{"User": user, "campaigns": models.Data, "pagination": pagination})
}

func (h *campaignHandler) Create(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	users, err := h.userService.GetAllUsers(1, 0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	models := []entity.User{}
	mapstructure.Decode(users.Data, &models)

	form := entity.CreateCampaignForm{
		Users: models,
		User:  user,
	}
	c.HTML(http.StatusOK, "campaign_create.html", form)
}

func (h *campaignHandler) PostCreate(c *gin.Context) {
	var form entity.CreateCampaignForm

	err := c.ShouldBind(&form)
	if err != nil {
		users, errUsers := h.userService.GetAllUsers(1, 0)
		if errUsers != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		models := []entity.User{}
		mapstructure.Decode(users.Data, &models)
		form.Users = models
		form.Error = err
		c.HTML(http.StatusOK, "campaign_create.html", form)
		return
	}

	req := entity.FormCampaignRequest{
		Title:            form.Title,
		ShortDescription: form.ShortDescription,
		Description:      form.Description,
		TargetAmount:     form.TargetAmount,
		Perks:            form.Perks,
		CampaignerID:     form.UserID,
	}

	_, err = h.service.CreateCampaign(req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) UploadImages(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	c.HTML(http.StatusOK, "campaign_image.html", gin.H{"ID": id, "User": user})
}

func (h *campaignHandler) PostUploadImages(c *gin.Context) {
	var path string
	id, _ := strconv.Atoi(c.Param("id"))
	file, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	campaign, err := h.service.GetCampaignByID(entity.CampaignIDRequest{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "local" {
		//Store File to Storage
		cleanFilename := helper.RemoveFileExt(file.Filename)
		filename := fmt.Sprintf("%d-%d-%s.jpg", campaign.CampaignerID, campaign.ID, cleanFilename)
		path = helper.GeneratePath("campaigns", filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.HTML(http.StatusOK, "campaign_image.html", gin.H{"Error": err})
			return
		}
		//Generate URL
		path = helper.GenerateURL("campaigns", filename)
	}
	if storageType == "cloud" {
		var ctx = context.Background()
		cloudinary := config.NewCloudStorage()
		uploadResponse, err := cloudinary.Upload.Upload(ctx, file, config.ConfigCloudStorage("campaigns"))
		if err != nil {
			c.HTML(http.StatusOK, "campaign_image.html", gin.H{"Error": err})
			return
		}
		//Generate URL
		path = uploadResponse.SecureURL
	}
	//Save URL To DB
	req := entity.UploadCampaignImageRequest{
		CampaignID: campaign.ID,
		UserID:     campaign.CampaignerID,
		IsPrimary:  true,
	}
	_, err = h.service.UploadCampaignImages(req, path)
	if err != nil {
		c.HTML(http.StatusOK, "campaign_image.html", gin.H{"Error": err})
		return
	}
	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) Edit(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := h.service.GetCampaignByID(entity.CampaignIDRequest{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	form := entity.EditCampaignForm{
		ID:               id,
		Title:            campaign.Title,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		TargetAmount:     campaign.TargetAmount,
		Perks:            campaign.Perks,
		User:             user,
	}
	c.HTML(http.StatusOK, "campaign_edit.html", form)
}

func (h *campaignHandler) PostEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form entity.EditCampaignForm

	err := c.ShouldBind(&form)
	if err != nil {
		form.Error = err
		form.ID = id
		c.HTML(http.StatusOK, "campaign_create.html", form)
		return
	}

	campaign, err := h.service.GetCampaignByID(entity.CampaignIDRequest{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	updateForm := entity.FormCampaignRequest{
		Title:            form.Title,
		ShortDescription: form.ShortDescription,
		Description:      form.Description,
		TargetAmount:     form.TargetAmount,
		Perks:            form.Perks,
		CampaignerID:     campaign.CampaignerID,
	}

	_, err = h.service.EditCampaign(entity.CampaignIDRequest{ID: id}, updateForm)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/campaigns")
}

func (h *campaignHandler) Detail(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := h.service.GetCampaignByID(entity.CampaignIDRequest{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	form := entity.EditCampaignForm{
		ID:               id,
		Title:            campaign.Title,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		TargetAmount:     campaign.TargetAmount,
		Perks:            campaign.Perks,
		User:             user,
	}
	c.HTML(http.StatusOK, "campaign_detail.html", form)
}
