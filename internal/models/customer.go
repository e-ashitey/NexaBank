package models

type Customer struct {
	ID       string
	Name     string
	Email    string
	Password string // hashed
	Phone    string
	Address  string
	Status   string
}
