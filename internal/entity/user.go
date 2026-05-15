package entity

import "time"

type User struct {
	ID string
	Name string
	Email string
	PasswordHash string
	Provider string
	ProviderID string
	CreatedAt time.Time
	UpdatedAt time.Time
}