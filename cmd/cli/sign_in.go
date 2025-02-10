package cli

import (
	"fmt"
	"nexacore/banking-system/db"
	"nexacore/banking-system/internal/customer"
)

func Auth() bool {
	// Ask user for login details
	fmt.Println("Enter Email:")
	var email string
	fmt.Scanln(&email)

	fmt.Println("Enter Password:")
	var password string
	fmt.Scanln(&password)

	// Call Customer Authenticate
	var cust = customer.Authenticate(email, password)
	if cust != nil {
		// Cache the customer
		db.LoggedInCustomer = *cust

		// Cache the customer's account
		db.LoggedInAccount = *customer.AuthenticatedAccount(cust.ID)

		fmt.Println("Login successful")
		return true
	}

	fmt.Print("No customer found with the provided credentials\n")
	return false
}
