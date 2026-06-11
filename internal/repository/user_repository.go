package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zahidmahfudz/collabforge-platform/internal/entity"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// insert user ke database
func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (id, first_name, last_name, mid_name, username, email, password_hash, provider, provider_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.MidName,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.Provider,
		user.ProviderID,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// CEK EMAIL EXISTS
func (r *UserRepository) IsEmailExists(ctx context.Context, email string) (bool, error) {

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`

	var exists bool

	err := r.db.QueryRow(ctx, query, email).Scan(&exists)

	return exists, err
}

// TEMUKAN USER BERDASARKAN EMAIL
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, first_name, last_name, mid_name, username, email, password_hash, provider, provider_id, created_at, updated_at FROM users WHERE email=$1`

	var user entity.User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MidName,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Provider,
		&user.ProviderID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// TEMUKAN USER BERDASARKAN ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT id, first_name, last_name, mid_name, username, email, password_hash, provider, provider_id, created_at, updated_at FROM users WHERE id=$1`

	var user entity.User

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MidName,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Provider,
		&user.ProviderID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
