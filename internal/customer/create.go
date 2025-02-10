package customer

import (
	"fmt"
	"math/rand"
	"nexacore/banking-system/db"
	"nexacore/banking-system/internal/account"
	"nexacore/banking-system/internal/models"
	"nexacore/banking-system/pkg/utils"

	"strconv"
	"time"
)

func Authenticate(email string, password string) *models.Customer {
	utils.DelayedLoader()
	for _, customer := range db.CustomersStore {
		if customer.Email == email && customer.Password == password {
			return &customer
		}
	}
	return nil
}

func AuthenticatedAccount(customerID string) *models.Account {
	for _, account := range db.AccountsStore {
		if account.CustomerID == customerID {
			return &account
		}
	}
	return nil
}

func CreateCustomer(name string, email string, password string, phone string, address string, initialDeposit float64) {
	// Create Customer model
	newCustomer := models.Customer{
		ID:       strconv.Itoa(rand.Intn(1000000)),
		Name:     name,
		Email:    email,
		Password: password,
		Phone:    phone,
		Address:  address,
	}
	// Validate
	valid := models.Validate(&newCustomer)
	if valid != nil {
		fmt.Printf("%v\n\n", valid)
		return
	}

	db.CustomersStore = append(db.CustomersStore, newCustomer)

	// Create Account
	account.CreateAccount(newCustomer.ID, initialDeposit)

	fmt.Println("\nProcessing...")
	time.Sleep(2 * time.Second) // Simulate a delay
	fmt.Printf("\nAccount Created Successfully: %v\n\n", newCustomer.Email)
}
