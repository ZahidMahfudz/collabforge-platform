package entity

import "time"

type RefreshToken struct {
	ID string
	UserID string
	TokenHash string
	ExpiresAt time.Time
	Revoked time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
