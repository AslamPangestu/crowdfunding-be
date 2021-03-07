package adapter

import "crowdfunding/entity"

//CampaignTransactionAdapter : Adapter Campaign Transaction
func CampaignTransactionAdapter(transaction entity.Transaction) entity.CampaignTransactionsResponse {
	return entity.CampaignTransactionsResponse{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
}

//CampaignTransactionsAdapter : Adapter Campaign Transactions
func CampaignTransactionsAdapter(transactions []entity.Transaction) []entity.CampaignTransactionsResponse {
	if len(transactions) == 0 {
		return []entity.CampaignTransactionsResponse{}
	}
	var transactionsResponse []entity.CampaignTransactionsResponse

	for _, transaction := range transactions {
		transactionsResponse = append(transactionsResponse, CampaignTransactionAdapter(transaction))

	}
	return transactionsResponse
}

//UserTransactionAdapter : Adapter User Transaction
func UserTransactionAdapter(transaction entity.Transaction) entity.UserTransactionsResponse {
	imageURL := ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		imageURL = transaction.Campaign.CampaignImages[0].ImagePath
	}
	campaign := entity.CampaignTransaction{
		Title:    transaction.Campaign.Title,
		ImageURL: imageURL,
	}
	return entity.UserTransactionsResponse{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		TRXCode:   transaction.TRXCode,
		CreatedAt: transaction.CreatedAt,
		Campaign:  campaign,
	}
}

//UserTransactionsAdapter : Adapter User Transactions
func UserTransactionsAdapter(transactions []entity.Transaction) []entity.UserTransactionsResponse {
	if len(transactions) == 0 {
		return []entity.UserTransactionsResponse{}
	}
	var transactionsResponse []entity.UserTransactionsResponse

	for _, transaction := range transactions {
		transactionsResponse = append(transactionsResponse, UserTransactionAdapter(transaction))

	}
	return transactionsResponse
}
