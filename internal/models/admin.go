package models

type Admin struct {
	Id          string
	Role        string
	Permissions []Permission
}

type Permission struct {
	ID          string
	Name        string // e.g., "view_reports"
	Description string // e.g., "Allows viewing system-wide reports"
}
