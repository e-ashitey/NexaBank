package transaction

import (
	"math/rand"
	"nexacore/banking-system/db"
	"nexacore/banking-system/internal/models"
	"nexacore/banking-system/pkg/enums"
	"strconv"
	"time"
)

func RecordTransaction(transactionType enums.TransactionType, amount float64) {
	var transaction = models.Transaction{
		TransactionId: strconv.Itoa(rand.Intn(1000000)),
		AccountId:     db.LoggedInAccount.CustomerID,
		Type:          transactionType,
		Amount:        amount,
		Fee:           0,
		Timestamp:     time.Now(),
		Reference:     "",
	}

	db.TransactionsStore = append(db.TransactionsStore, transaction)
}

func TransactionsByAccount(accountId string) []models.Transaction {
	var listTransactions []models.Transaction
	for _, transaction := range db.TransactionsStore {
		if transaction.AccountId == accountId {
			listTransactions = append(listTransactions, transaction)
		}
	}
	return listTransactions
}
