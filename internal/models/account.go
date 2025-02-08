package models

import (
	"nexacore/banking-system/pkg/enums"
	"time"
)

type Account struct {
	Number    string
	UserId    string
	Type      enums.AccountType
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
	// Currency  string
}
