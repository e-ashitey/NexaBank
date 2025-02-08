package enums

type AccountType string

const (
	Savings      AccountType = "savings"
	Checking     AccountType = "checking"
	FixedDeposit AccountType = "fixed_deposit"
)

type TransactionType string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
	Transfer   TransactionType = "transfer"
)
