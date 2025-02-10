package cli

import (
	"fmt"
	"nexacore/banking-system/db"
	"nexacore/banking-system/internal/models"
	"nexacore/banking-system/internal/transaction"
	"nexacore/banking-system/pkg/enums"
	"nexacore/banking-system/pkg/utils"
)

func MainSystem() {
	fmt.Printf("\n%30s\n", fmt.Sprintf("Welcome %v to NexaCore Banking System", db.LoggedInCustomer.Name))
	for {
		// List out the available commands
		fmt.Println("Available Commands:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Transfer")
		fmt.Println("4. View Balance")
		fmt.Println("5. View Transactions")
		fmt.Println("99. Exit")

		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "1":
			var depositAmount float64
			fmt.Println("How much would you like to deposit?")
			fmt.Scanln(&depositAmount)
			if depositAmount <= 0 {
				fmt.Print("Invalid Amount\n\n")
				break
			}
			utils.DelayedLoader()
			db.LoggedInAccount.Balance += depositAmount
			transaction.RecordTransaction(enums.Deposit, depositAmount)
			fmt.Print("_____________________________________\n\n")
			fmt.Println("Deposit Success")
			fmt.Println("_____________________________________")
		case "2":
			var withdrawAmount float64
			fmt.Println("How much would you like to withdraw?")
			fmt.Scanln(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Print("Invalid Amount\n\n")
				break
			}
			if withdrawAmount > db.LoggedInAccount.Balance {
				fmt.Print("Insufficient funds\n\n")
				break
			}
			utils.DelayedLoader()
			db.LoggedInAccount.Balance -= withdrawAmount
			transaction.RecordTransaction(enums.Withdrawal, withdrawAmount)
			fmt.Print("_____________________________________\n\n")
			fmt.Println("Withdrawal Success")
			fmt.Println("_____________________________________")
		case "3":
			fmt.Println("Transfer")
		case "4":
			utils.DelayedLoader()
			fmt.Print("_____________________________________")
			fmt.Printf("Current Balance: %v\n\n", db.LoggedInAccount.Balance)
			fmt.Print("_____________________________________\n\n")
		case "5":
			fmt.Println("Transaction History")
			utils.DelayedLoader()
			var transactions = transaction.TransactionsByAccount(db.LoggedInAccount.CustomerID)
			if len(transactions) == 0 {
				fmt.Println("_____________________________________")
				fmt.Println("No Transactions Yet...")
				fmt.Print("_____________________________________\n\n")

			}
			for _, transaction := range transactions {
				fmt.Println("_____________________________________")
				fmt.Printf("Transaction Type: %v", transaction.Type)
				fmt.Printf("Amount: %v", transaction.Amount)
				fmt.Printf("Time: %v", transaction.Timestamp)
				fmt.Printf("Reference: %v\n", transaction.Reference)
				fmt.Printf("_____________________________________\n\n")
			}
		case "99":
			fmt.Println("Sign out...")
			db.LoggedInCustomer = models.Customer{}
			return
		}
	}
}
