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
	"github.com/zahidmahfudz/collabforge-platform/internal/service/token"
	"github.com/zahidmahfudz/collabforge-platform/utils"

	"golang.org/x/crypto/bcrypt"
)

var Logger = config.Logger

type AuthUseCase struct {
	userRepo *repository.UserRepository
	refreshTokenRepo *repository.RefreshTokenRepository
	pasetoService *token.PasetoService
}

func NewAuthUseCase(userRepo *repository.UserRepository, refreshTokenRepo *repository.RefreshTokenRepository, pasetoService *token.PasetoService) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo, refreshTokenRepo: refreshTokenRepo, pasetoService: pasetoService}
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
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MidName:      req.MidName,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Provider:     "local",
		ProviderID:   "",
		Bio:          "",
		AvatarURL:    "",
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
		Name:  user.FirstName + " " + user.MidName + " " + user.LastName,
		Email: user.Email,
	}, nil
}

func (u *AuthUseCase) Login(ctx context.Context, req request.LoginRequest) (*dtoresponse.LoginResponse, string, error) {
	Logger.Debug("Memasuki Login UseCase")

	user, err := u.userRepo.FindByEmail(ctx, req.Email)

	if err != nil {
		Logger.Errorf("Error saat mencari user: %v", err)
		return nil, "", errors.New("INVALID_CREDENTIALS")
	}
	Logger.Debugf("User ditemukan: %+v", user)

	// cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		Logger.Errorf("Error saat membandingkan password: %v", err)
		return nil, "", errors.New("INVALID_CREDENTIALS")
	}
	Logger.Debug("Password valid")

	// generate access token
	accessToken, err := u.pasetoService.GenerateAccessToken(user.ID, user.Email, 5*time.Minute)
	if err != nil {
		Logger.Errorf("Error saat menghasilkan access token: %v", err)
		return nil, "", errors.New("FAILED_TO_GENERATE_TOKEN")
	}
	Logger.Debug("Access token berhasil dibuat")

	// generate refresh token
	refreshToken, err := u.pasetoService.GenerateRefreshToken(user.ID, 7*24*time.Hour)
	if err != nil {
		Logger.Errorf("Error saat menghasilkan refresh token: %v", err)
		return nil, "", errors.New("FAILED_TO_GENERATE_TOKEN")
	}
	Logger.Debug("Refresh token berhasil dibuat")

	// hash refresh token sebelum disimpan ke database
	refreshTokenHash, err := utils.HashToken(refreshToken)
	if err != nil {
		Logger.Errorf("Error saat menghash refresh token: %v", err)
		return nil, "", errors.New("FAILED_TO_GENERATE_TOKEN")
	}

	// generate id untuk refresh token
	refreshTokenID, err := utils.GenerateID("rft")
	if err != nil {
		Logger.Errorf("Error saat menghasilkan ID refresh token: %v", err)
		return nil, "", errors.New("FAILED_TO_GENERATE_TOKEN")
	}

	// mapping entity refresh token
	refreshTokenEntity := entity.RefreshToken{
		ID: refreshTokenID,
		UserID: user.ID,
		TokenHash: refreshTokenHash,
		ExpiresAt: time.Now().Add(7*24*time.Hour),
		CreatedAt: time.Now(),
	}

	// simpan refresh token ke database
	err = u.refreshTokenRepo.CreateRefreshToken(ctx, &refreshTokenEntity)
	if err != nil {
		Logger.Errorf("Error saat menyimpan refresh token: %v", err)
		return nil, "", errors.New("FAILED_TO_GENERATE_TOKEN")
	}

	// mapping response
	return &dtoresponse.LoginResponse{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		MidName: user.MidName,
		Username: user.Username,
		Email: user.Email,
		AccessToken: accessToken,
	}, refreshToken, nil
}