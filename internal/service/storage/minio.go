package storage

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	Client *minio.Client
	Bucket string
}

func NewMinioStorage(endpoint, accessKeyID, secretAccessKey, bucketName string) (*MinioStorage, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return &MinioStorage{
		Client: client,
		Bucket: bucketName,
	}, nil
}

func (s *MinioStorage) GetPresignedURL(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	presignedURL, err := s.Client.PresignedGetObject(ctx, s.Bucket, objectName, expiry, nil)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}
