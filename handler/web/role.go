package handler

import (
	"crowdfunding/entity"
	"crowdfunding/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	service services.RoleInteractor
}

func RoleHandlerInit(service services.RoleInteractor) *roleHandler {
	return &roleHandler{service}
}

func (h *roleHandler) Index(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	fmt.Println("USER", user)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	models, err := h.service.GetRoles(page, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "role_index.html", gin.H{"roles": models.Data, "pagination": models.Pagination})
}

func (h *roleHandler) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "role_create.html", nil)
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
	id, _ := strconv.Atoi(c.Param("id"))
	model, err := h.service.GetRoleByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	user := entity.EditRoleForm{
		ID:   id,
		Name: model.Name,
	}
	c.HTML(http.StatusOK, "role_edit.html", user)
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
