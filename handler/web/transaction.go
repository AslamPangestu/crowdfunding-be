package handler

import (
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
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	models, err := h.service.GetTransactions(page, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": models})
}
