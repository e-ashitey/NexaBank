package account

import (
	"math/rand"
	"nexacore/banking-system/db"
	"nexacore/banking-system/internal/models"
	"nexacore/banking-system/pkg/enums"
	"strconv"
	"time"
)

func CreateAccount(customerID string, initialDeposit float64) {
	newAccount := models.Account{
		CustomerID: customerID,
		Number:     strconv.Itoa(rand.Intn(1000000)),
		Type:       enums.Savings,
		Balance:    initialDeposit,
		Status:     "active",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	db.AccountsStore = append(db.AccountsStore, newAccount)
}
