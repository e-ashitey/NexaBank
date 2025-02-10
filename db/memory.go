package db

import "nexacore/banking-system/internal/models"

// In-memory storage for Customers, Accounts, and Transactions
var CustomersStore = make([]models.Customer, 0)
var AccountsStore = make([]models.Account, 0)
var TransactionsStore = make([]models.Transaction, 0)

// Logged in customer
var LoggedInCustomer models.Customer
var LoggedInAccount models.Account
