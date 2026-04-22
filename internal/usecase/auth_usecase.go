package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/zahidmahfudz/collabforge-platform/internal/domain/entity"
	"github.com/zahidmahfudz/collabforge-platform/internal/domain/repository"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/request"
	"github.com/zahidmahfudz/collabforge-platform/internal/dto/response"
	AppError "github.com/zahidmahfudz/collabforge-platform/pkg/errors"
	"github.com/zahidmahfudz/collabforge-platform/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

type AuthUsecase interface {
	Register(ctx context.Context, req request.RegisterRequest) (*response.AuthResponse, error)
	LoginGoogle(ctx context.Context, req request.GoogleLoginRequest) (*response.AuthResponse, error)
}

type authUsecaseImpl struct {
	userRepo       repository.UserRepository
	secretKey      string
	googleClientID string
}

func NewAuthUsecase(userRepo repository.UserRepository, secretKey string, googleClientID string) AuthUsecase {
	return &authUsecaseImpl{
		userRepo:       userRepo,
		secretKey:      secretKey,
		googleClientID: googleClientID,
	}
}

func (u *authUsecaseImpl) Register(ctx context.Context, req request.RegisterRequest) (*response.AuthResponse, error) {
	// 1. Cek email sudah ada atau belum
	existingUser, _ := u.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, AppError.ErrEmailAlreadyUsed
	}

	// 2. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashedPasswordStr := string(hashedPassword)

	// 3. Generate ID
	id, err := utils.GenerateID("usr", 16)
	if err != nil {
		return nil, err
	}

	// 4. Create entity
	now := time.Now()
	user := &entity.User{
		ID:        id,
		Name:      req.Name,
		Email:     req.Email,
		Password:  &hashedPasswordStr,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	// 5. Save user
	if err := u.userRepo.Create(ctx, user); err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return nil, err
	}

	// 6. Generate Token
	payload := utils.TokenPayload{
		UserID: user.ID,
		Email:  user.Email,
		Exp:    now.Add(24 * time.Hour),
	}
	token, err := utils.GeneratePasetoToken(payload, u.secretKey)
	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{
		Token: token,
		User: response.UserDataResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (u *authUsecaseImpl) LoginGoogle(ctx context.Context, req request.GoogleLoginRequest) (*response.AuthResponse, error) {
	// 1. Verifikasi token google
	payload, err := idtoken.Validate(ctx, req.IDToken, u.googleClientID)
	if err != nil {
		return nil, AppError.New("Invalid Google Token", "UNAUTHORIZED", 401)
	}

	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)
	providerID := payload.Subject

	// 2. Cek user
	user, _ := u.userRepo.FindByEmail(ctx, email)
	if user == nil {
		// Create new user
		id, err := utils.GenerateID("usr", 16)
		if err != nil {
			return nil, err
		}
		now := time.Now()
		user = &entity.User{
			ID:         id,
			Name:       name,
			Email:      email,
			Provider:   "google",
			ProviderID: &providerID,
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		if err := u.userRepo.Create(ctx, user); err != nil {
			return nil, err
		}
	}

	// 3. Generate Token
	now := time.Now()
	tokenPayload := utils.TokenPayload{
		UserID: user.ID,
		Email:  user.Email,
		Exp:    now.Add(24 * time.Hour),
	}
	token, err := utils.GeneratePasetoToken(tokenPayload, u.secretKey)
	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{
		Token: token,
		User: response.UserDataResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
