package models

import (
	"nexacore/banking-system/pkg/enums"
	"time"
)

type Transaction struct {
	TransactionId string
	AccountId     string
	Type          enums.TransactionType
	Amount        float64
	Fee           float64
	Timestamp     time.Time
	Reference     string
}
