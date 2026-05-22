package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zahidmahfudz/collabforge-platform/internal/entity"
)

type RefreshTokenRepository struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepository(db *pgxpool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}


func (r *RefreshTokenRepository) CreateRefreshToken(ctx context.Context, token *entity.RefreshToken) error {
	query := `
		INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(ctx, query, token.ID, token.UserID, token.TokenHash, token.ExpiresAt, token.CreatedAt, token.UpdatedAt)
	return err
}

func (r *RefreshTokenRepository) FindByToken(ctx context.Context, tokenHash string) (*entity.RefreshToken, error) {
	query := `
		SELECT id, user_id, token_hash, expires_at, revoked, created_at, updated_at 
		FROM refresh_tokens 
		WHERE token_hash = $1
	`
	var token entity.RefreshToken
	err := r.db.QueryRow(ctx, query, tokenHash).Scan(
		&token.ID,
		&token.UserID,
		&token.TokenHash,
		&token.ExpiresAt,
		&token.Revoked,
		&token.CreatedAt,
		&token.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *RefreshTokenRepository) RevokeToken(ctx context.Context, tokenID string) error {
	query := `
		UPDATE refresh_tokens
		SET revoked = NOW(), updated_at = NOW()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, tokenID)
	return err
}