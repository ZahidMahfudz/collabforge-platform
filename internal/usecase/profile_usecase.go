package usecase

import (
	"context"
	"time"

	"github.com/zahidmahfudz/collabforge-platform/internal/service/storage"
)

type ProfileUseCase struct {
	StorageService *storage.MinioStorage
}

func NewProfileUseCase(storageService *storage.MinioStorage) *ProfileUseCase {
	return &ProfileUseCase{StorageService: storageService}
}

func (u *ProfileUseCase) GetAvatarURL(ctx context.Context) (string, error) {
	return u.StorageService.GetPresignedURL(ctx, "avatar/foto_zahid (1) (1).jpg", 15*time.Minute)
}
