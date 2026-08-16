package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, displayName string) (User, error) {
	const query = `
		INSERT INTO users (display_name)
		VALUES ($1)
		RETURNING
			id::text,
			COALESCE(telegram_id, 0),
			COALESCE(telegram_username, ''),
			display_name,
			COALESCE(avatar_url, ''),
			created_at,
			updated_at
	`

	var user User

	err := r.db.QueryRow(
		ctx,
		query,
		displayName,
	).Scan(
		&user.ID,
		&user.TelegramID,
		&user.TelegramUsername,
		&user.DisplayName,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (User, error) {
	const query = `
		SELECT
			id::text,
			COALESCE(telegram_id, 0),
			COALESCE(telegram_username, ''),
			display_name,
			COALESCE(avatar_url, ''),
			created_at,
			updated_at
		FROM users
		WHERE id = $1::uuid
	`

	var user User

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.TelegramID,
		&user.TelegramUsername,
		&user.DisplayName,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
