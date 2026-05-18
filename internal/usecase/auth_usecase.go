package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	dtoresponse "github.com/zahidmahfudz/collabforge-platform/internal/dto/response"
	"github.com/zahidmahfudz/collabforge-platform/internal/entity"
	"github.com/zahidmahfudz/collabforge-platform/internal/repository"
	"github.com/zahidmahfudz/collabforge-platform/utils"

	"golang.org/x/crypto/bcrypt"
)

var Logger = config.Logger

type AuthUseCase struct {
	userRepo *repository.UserRepository
}

func NewAuthUseCase(userRepo *repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo,}
}

func (u *AuthUseCase) Register(ctx context.Context,req request.RegisterRequest,) (*dtoresponse.RegisterResponse, error) {
	Logger.Debug("Memasuki Register UseCase")

	// cek email exists
	exists, err := u.userRepo.IsEmailExists(ctx,req.Email)
	if err != nil {
		Logger.Errorf("Error cek email: %v",err)
		return nil, err
	}

	if exists {
		Logger.Debug("Email sudah terdaftar",)
		return nil, errors.New(
			"EMAIL_ALREADY_EXISTS",
		)
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	// generate id
	id, err := utils.GenerateID("usr")

	if err != nil {
		return nil, err
	}

	// mapping entity
	user := entity.User{
		ID:           id,
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Provider:     "local",
		ProviderID:   "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// save database
	err = u.userRepo.CreateUser(ctx,&user,)

	if err != nil {
		return nil, err
	}

	Logger.Debug("User berhasil dibuat")

	// mapping response
	return &dtoresponse.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}