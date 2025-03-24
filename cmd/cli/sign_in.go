package cli

import (
	"bufio"
	"fmt"
	"nexacore/banking-system/db"
	"nexacore/banking-system/internal/customer"
	"os"
	"strings"
)

func Auth() bool {
	reader := bufio.NewReader(os.Stdin)

	// Ask user for login details
	fmt.Println("Enter Email:")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Println("Enter Password:")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

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
