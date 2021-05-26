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
	service services.RoleInteractor
}

func RoleHandlerInit(service services.RoleInteractor) *roleHandler {
	return &roleHandler{service}
}

func (h *roleHandler) Index(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	models, err := h.service.GetRoles(page, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	pagination := helper.PaginationAdapterHandler(models.Pagination)
	c.HTML(http.StatusOK, "role_index.html", gin.H{"User": user, "roles": models.Data, "pagination": pagination})
}

func (h *roleHandler) Create(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "role_create.html", gin.H{"User": user})
}

func (h *roleHandler) PostCreate(c *gin.Context) {
	var form entity.CreateRoleForm

	err := c.ShouldBind(&form)
	if err != nil {
		form.Error = err
		c.HTML(http.StatusOK, "role_create.html", form)
		return
	}

	createForm := entity.FormRoleRequest{
		Name: form.Name,
	}

	_, err = h.service.AddRole(createForm)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/roles")
}

func (h *roleHandler) Edit(c *gin.Context) {
	user, err := helper.GetUserLoggedIn(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	model, err := h.service.GetRoleByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	form := entity.EditRoleForm{
		ID:   id,
		Name: model.Name,
		User: user,
	}
	c.HTML(http.StatusOK, "role_edit.html", form)
}

func (h *roleHandler) PostEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form entity.EditRoleForm
	err := c.ShouldBind(&form)
	if err != nil {
		form.Error = err
		c.HTML(http.StatusOK, "user_edit.html", form)
		return
	}
	req := entity.FormRoleRequest{
		Name: form.Name,
	}
	_, err = h.service.EditRole(id, req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/roles")
}

func (h *roleHandler) Remove(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.RemoveRole(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/roles")
}
