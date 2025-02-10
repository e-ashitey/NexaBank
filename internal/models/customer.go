package models

import "errors"

type Customer struct {
	ID       string
	Name     string
	Email    string
	Password string // hashed
	Phone    string
	Address  string
	Status   string
}

func Validate(c *Customer) error {
	if c.Name == "" || len(c.Name) < 5 {
		return errors.New("name can't be empty and must be at least 5 characters")
	}
	if c.Email == "" || c.Password == "" || c.Phone == "" || c.Address == "" {
		return errors.New("all fields are required")
	}
	return nil
}
