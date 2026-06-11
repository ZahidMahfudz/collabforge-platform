package storage

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
)

type MinioStorage struct {
	client *minio.Client
	bucket string
}

func NewMinioStorage(client *minio.Client, bucket string) *MinioStorage {
	return &MinioStorage{
		client: client,
		bucket: bucket,
	}
}

func (s *MinioStorage) GetPresignedURL(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	presignedURL, err := s.client.PresignedGetObject(ctx, s.bucket, objectName, expiry, nil)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}
