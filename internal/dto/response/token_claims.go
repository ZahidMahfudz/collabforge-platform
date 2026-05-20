package response

import "time"

type TokenClaims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
}