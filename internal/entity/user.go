package entity

import "time"

type User struct {
	ID           string
	FirstName    string
	LastName     string
	MidName      string
	Username     string
	Email        string
	PasswordHash string
	Provider     string
	ProviderID   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
