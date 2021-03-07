package adapter

import "crowdfunding/entity"

//TransactionAdapter : Adapter Transaction
func TransactionAdapter(transaction entity.Transaction) entity.GetCampaignTransactionsResponse {
	return entity.GetCampaignTransactionsResponse{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
}

//TransactionsAdapter : Adapter Transactions
func TransactionsAdapter(transactions []entity.Transaction) []entity.GetCampaignTransactionsResponse {
	if len(transactions) == 0 {
		return []entity.GetCampaignTransactionsResponse{}
	}
	var transactionsResponse []entity.GetCampaignTransactionsResponse

	for _, transaction := range transactions {
		transactionsResponse = append(transactionsResponse, TransactionAdapter(transaction))

	}
	return transactionsResponse
}
