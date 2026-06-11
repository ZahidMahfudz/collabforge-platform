package storage

import (
	"context"
	"time"
)

type StorageService interface {
	GetPresignedURL(ctx context.Context, objectName string, expiry time.Duration) (string, error)
}
