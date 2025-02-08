package models

type Transaction struct {
	TransactionId string
	AccountId     string
	Type          string
	Amount        string
	Fee           string
	Timestamp     string
	Reference     string
}
