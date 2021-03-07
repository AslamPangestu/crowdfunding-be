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
ROUTE: api/users/transactions
METHOD: POST
*/
func (h *transactionHandler) MakeTransaction(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	//GET REQUEST TRANSACTION
	var request entity.TransactionRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("MakeTransaction Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//SET BACKER TRANSACTION
	request.Backer = currentUser
	//SAVE TRANSACTION DB
	newTransaction, err := h.service.MakeTransaction(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("MakeTransaction Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := adapter.TransactionAdapter(newTransaction)
	res := helper.ResponseHandler("MakeTransaction Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/campaigns/:id/transactions
METHOD: GET
*/
func (h *transactionHandler) GetCamapaignTransactions(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	//GET ID CAMPAIGN
	var request entity.CampaignTransactionsRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCamapaignTransactions Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//SET OWNER CAMPAIGN
	request.CampaignerID = currentUser.ID
	//GET TRANSACTIONS CAMPAIGN
	transactions, err := h.service.GetTransactionsByCampaignID(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCamapaignTransactions Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := adapter.CampaignTransactionsAdapter(transactions)
	res := helper.ResponseHandler("GetCamapaignTransactions Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

/**
ROUTE: api/v1/users/transactions
METHOD: GET
*/
func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	//GET TRANSACTIONS
	transactions, err := h.service.GetTransactionsByUserID(currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetUserTransactions Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	data := adapter.CampaignTransactionsAdapter(transactions)
	res := helper.ResponseHandler("GetUserTransactions Successful", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
