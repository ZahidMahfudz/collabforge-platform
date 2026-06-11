package usecase

import (
	"context"
	"time"

	"github.com/zahidmahfudz/collabforge-platform/internal/service/storage"
)

type ProfileUseCase struct {
	storage storage.StorageService
}

func NewProfileUseCase(storage storage.StorageService) *ProfileUseCase {
	return &ProfileUseCase{storage: storage}
}

func (u *ProfileUseCase) GetAvatarURL(ctx context.Context) (string, error) {
	return u.storage.GetPresignedURL(ctx, "avatar/foto_zahid (1) (1).jpg", 15*time.Minute)
}
