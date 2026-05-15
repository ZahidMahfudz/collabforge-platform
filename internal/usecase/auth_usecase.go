package usecase

import (
	"context"
	"time"

	"github.com/zahidmahfudz/collabforge-platform/config"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/response"
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
	return &AuthUseCase{userRepo: userRepo}
}

func (u *AuthUseCase) Register(ctx context.Context, req request.RegisterRequest) (*response.RegisterResponse, error) {
	Logger.Debug("Memasuki auth_usecase di fungsi Register")
	// CEK EMAIL EXISTS
	exists, err := u.userRepo.IsEmailExists(ctx, req.Email)
	if err != nil {
		Logger.Errorf("Error saat cek email exists: %v", err)
		return nil, err
	}

	if exists {
		Logger.Debug("Email sudah terdaftar")
		return nil, err
	}
	Logger.Debug("Email belum terdaftar")

	// HASH PASSWORD
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		Logger.Errorf("Error saat hash password: %v", err)
		return nil, err
	}
	Logger.Debug("Password berhasil di-hash")

	// GENERATE ID
	id, err := utils.GenerateID("usr")
	if err != nil {
		Logger.Errorf("Error saat generate ID: %v", err)
		return nil, err
	}
	Logger.Debugf("ID user yang dihasilkan: %s", id)

	// MAPPING KE ENTITY
	user := entity.User{
		ID: id,
		Name: req.Name,
		Email: req.Email,
		PasswordHash: string(hashedPassword),
		Provider: "local",
		ProviderID: "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// SIMPAN KE DATABASE
	err = u.userRepo.CreateUser(ctx, &user)
	if err != nil {
		Logger.Errorf("Error saat menyimpan user: %v", err)
		return nil, err
	}
	Logger.Debug("User berhasil dibuat")
	Logger.Debugf("User yang disimpan: %+v", user)

	// MAPPING KE RESPONSE
	Logger.Debug("Mapping user entity ke RegisterResponse")
	Logger.Debugf("Data user yang akan dikirim di response: %+v", user)
	return &response.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
