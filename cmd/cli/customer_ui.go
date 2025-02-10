package cli

import (
	"fmt"
	"nexacore/banking-system/internal/customer"
)

func CustomerUi() {
	// Ask user for customer details
	fmt.Println("Enter Customer Name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter Customer Email:")
	var email string
	fmt.Scanln(&email)

	fmt.Println("Enter Customer Password:")
	var password string
	fmt.Scanln(&password)

	fmt.Println("Enter Customer Phone:")
	var phone string
	fmt.Scanln(&phone)

	fmt.Println("Enter Customer Address:")
	var address string
	fmt.Scanln(&address)

	fmt.Println("Enter Initial Deposit:")
	var initialDeposit float64
	fmt.Scanln(&initialDeposit)

	// Call Customer CreateCustomer
	customer.CreateCustomer(name, email, password, phone, address, initialDeposit)
}
