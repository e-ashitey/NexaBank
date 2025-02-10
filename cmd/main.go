package main

import (
	"fmt"
	"nexacore/banking-system/cmd/cli"
	"nexacore/banking-system/db"
)

// Volatile data
// This data will be lost when the program is restarted
// This is where we will store the data for the banking system

// var transactions = make([]models.Transaction, 0)
func main() {

	// var accountsStore = make([]models.Account, 0)
	// Welcome message
	fmt.Println("         Welcome To NexaCore         ")
	fmt.Println("\"Your Financial Future, Connected.\"")
	fmt.Print("_______________________________________\n\n")

	// Main loop
	for {
		// List out the available commands
		fmt.Println("Available Commands:")
		fmt.Println("1. Sign In")
		fmt.Println("2. Create Account")
		fmt.Println("99. Exit")

		// Get the command from the user
		var command int
		fmt.Scanln(&command)

		// Handle the command
		switch command {
		case 1:
			var isAuthenticated = cli.Auth()
			if isAuthenticated {
				cli.MainSystem()
			}
		case 2:
			cli.CustomerUi()
		case 99:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Printf("Invalid Command: %v", db.CustomersStore)
			// fmt.Printf("Invalid Command: %v", db.LoggedInCustomer)
		}
	}
}
