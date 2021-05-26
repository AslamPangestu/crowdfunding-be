package handler

import (
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service services.TransactionInteractor
}

func TransactionHandlerInit(service services.TransactionInteractor) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) Index(c *gin.Context) {
	user := helper.GetUserLoggedIn(c)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	models, err := h.service.GetTransactions(page, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	pagination := helper.PaginationAdapterHandler(models.Pagination)
	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"user": user, "transactions": models.Data, "pagination": pagination})
}
