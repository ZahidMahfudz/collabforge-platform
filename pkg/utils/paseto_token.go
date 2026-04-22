package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto/v2"
)

type TokenPayload struct {
	UserID string    `json:"user_id"`
	Email  string    `json:"email"`
	Exp    time.Time `json:"exp"`
}

func GeneratePasetoToken(payload TokenPayload, secretKey string) (string, error) {
	if len(secretKey) < 32 {
		return "", errors.New("secret key must be at least 32 characters")
	}

	token := paseto.NewV2()

	// Create map for claims
	claims := make(map[string]interface{})
	claims["user_id"] = payload.UserID
	claims["email"] = payload.Email
	claims["exp"] = payload.Exp.Format(time.RFC3339)

	signed, err := token.Encrypt([]byte(secretKey), claims, nil)
	if err != nil {
		fmt.Printf("Error encrypting token: %v\n", err)
		return "", err
	}

	return signed, nil
}
