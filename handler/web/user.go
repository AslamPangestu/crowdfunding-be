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
)

type userHandler struct {
	service services.UserInteractor
}

func UserHandlerInit(service services.UserInteractor) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) Index(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	models, err := h.service.GetAllUsers(page, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	pagination := helper.PaginationAdapterHandler(models.Pagination)
	c.HTML(http.StatusOK, "user_index.html", gin.H{"User": user, "users": models.Data, "pagination": pagination})
}

func (h *userHandler) Create(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "user_create.html", gin.H{"User": user})
}

func (h *userHandler) PostCreate(c *gin.Context) {
	var form entity.CreateUserForm

	err := c.ShouldBind(&form)
	if err != nil {
		form.Error = err
		c.HTML(http.StatusOK, "user_create.html", form)
		return
	}

	registerInput := entity.RegisterRequest{
		Name:       form.Name,
		Username:   form.Username,
		Occupation: form.Occupation,
		Email:      form.Email,
		Password:   form.Password,
	}

	_, err = h.service.Register(registerInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) Edit(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	id := c.Param("id")
	model, err := h.service.GetUserByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	form := entity.EditUserForm{
		ID:         id,
		Name:       model.Name,
		Username:   model.Username,
		Email:      model.Email,
		Occupation: model.Occupation,
		User:       user,
	}
	c.HTML(http.StatusOK, "user_edit.html", form)
}

func (h *userHandler) PostEdit(c *gin.Context) {
	id := c.Param("id")
	var form entity.EditUserForm
	err := c.ShouldBind(&form)
	if err != nil {
		form.Error = err
		c.HTML(http.StatusOK, "user_edit.html", form)
		return
	}
	form.ID = id
	_, err = h.service.UpdateUser(form)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// _, err := h.service.GetUserByID(id)
	// if err != nil {
	// 	c.HTML(http.StatusInternalServerError, "error.html", nil)
	// 	return
	// }
	c.HTML(http.StatusOK, "user_avatar.html", gin.H{"ID": id})
}

func (h *userHandler) PostUploadAvatar(c *gin.Context) {
	var path string
	id := c.Param("id")
	file, err := c.FormFile("avatar")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "local" {
		//Store File to Storage
		filename := fmt.Sprintf("%d-%s.jpg", id, user.Username)
		path = helper.GeneratePath("avatars", filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.HTML(http.StatusOK, "user_avatar.html", gin.H{"Error": err})
			return
		}
		//Generate URL
		path = helper.GenerateURL("avatars", filename)
	}
	if storageType == "cloud" {
		openedFile, err := file.Open()
		if err != nil {
			c.HTML(http.StatusOK, "user_avatar.html", gin.H{"Error": err})
			return
		}
		var ctx = context.Background()
		cloudinary := config.NewCloudStorage()
		uploadResponse, err := cloudinary.Upload.Upload(ctx, openedFile, config.ConfigCloudStorage("avatars"))
		if err != nil {
			c.HTML(http.StatusOK, "user_avatar.html", gin.H{"Error": err})
			return
		}
		//Generate URL
		path = uploadResponse.SecureURL
	}
	//Save URL To DB
	_, err = h.service.UploadAvatar(user.ID, path)
	if err != nil {
		c.HTML(http.StatusOK, "user_avatar.html", gin.H{"Error": err})
		return
	}
	c.Redirect(http.StatusFound, "/users")
}
