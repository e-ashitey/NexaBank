package models

import "nexacore/banking-system/pkg/enums"

type Transaction struct {
	TransactionId string
	AccountId     string
	Type          enums.TransactionType
	Amount        string
	Fee           string
	Timestamp     string
	Reference     string
}
