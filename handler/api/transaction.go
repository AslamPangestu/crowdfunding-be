package handler

import (
	"crowdfunding/adapter"
	"crowdfunding/entity"
	"crowdfunding/helper"
	"crowdfunding/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type transactionHandler struct {
	service        services.TransactionInteractor
	paymentService services.PaymentInteractor
}

// TransactionHandlerInit Initiation
func TransactionHandlerInit(service services.TransactionInteractor, paymentService services.PaymentInteractor) *transactionHandler {
	return &transactionHandler{service, paymentService}
}

/*
*
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

/*
*
ROUTE: api/v1/campaigns/:id/transactions
METHOD: GET
*/
func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	//GET ID CAMPAIGN
	var request entity.CampaignTransactionsRequest
	err := c.ShouldBindUri(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCampaignTransactions Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//SET OWNER CAMPAIGN
	request.CampaignerID = currentUser.ID
	//GET TRANSACTIONS CAMPAIGN
	//Get Query
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	transactions, err := h.service.GetTransactionsByCampaignID(request, page, pageSize)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetCampaignTransactions Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	models := []entity.Transaction{}
	mapstructure.Decode(transactions.Data, &models)
	data := adapter.CampaignTransactionsAdapter(models)
	res := helper.ResponseHandler("GetCampaignTransactions Successful", http.StatusOK, "success", helper.ResponsePagination{
		Data:       data,
		Pagination: transactions.Pagination,
	})
	c.JSON(http.StatusOK, res)
}

/*
*
ROUTE: api/v1/users/transactions
METHOD: GET
*/
func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	//Get User Logged
	currentUser := c.MustGet("currentUser").(entity.User)
	//Get Query
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	//GET TRANSACTIONS
	transactions, err := h.service.GetTransactionsByUserID(currentUser.ID, page, pageSize)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errResponse := helper.ResponseHandler("GetUserTransactions Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	models := []entity.Transaction{}
	mapstructure.Decode(transactions.Data, &models)
	data := adapter.UserTransactionsAdapter(models)
	res := helper.ResponseHandler("GetUserTransactions Successful", http.StatusOK, "success", helper.ResponsePagination{
		Data:       data,
		Pagination: transactions.Pagination,
	})
	c.JSON(http.StatusOK, res)
}

/*
*
ROUTE: api/v1/users/transactions
METHOD: GET
*/
func (h *transactionHandler) GetNotification(c *gin.Context) {
	//GET REQUEST
	var request entity.TransactionNotificationRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("GetNotification Validation Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, errResponse)
		return
	}
	//PROCESS
	err = h.paymentService.ProcessPayment(request)
	if err != nil {
		errorMessage := gin.H{"errors": helper.ErrResponseValidationHandler(err)}
		errResponse := helper.ResponseHandler("GetNotification Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	//RESPONSE
	c.JSON(http.StatusOK, request)
}
