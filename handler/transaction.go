package handler

import (
	"crowdfunding/adapter"
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service services.TransactionInteractor
}

// TransactionHandlerInit Initiation
func TransactionHandlerInit(service services.TransactionInteractor) *transactionHandler {
	return &transactionHandler{service}
}

/**
ROUTE: api/v1/campaigns/:id/transactions
METHOD: POST
*/
func (h *transactionHandler) GetCamapaignTransactions(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	var request entity.GetCampaignTransactionsRequest
	//GET ID CAMPAIGN
	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCamapaignTransactions Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//GET CAMPAIGN
	request.User = currentUser
	transactions, err := h.service.GetTransactionsByCampaignID(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCamapaignTransactions Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	data := adapter.TransactionsAdapter(transactions)
	res := helper.ResponseHandler("GetCamapaignTransactions Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
