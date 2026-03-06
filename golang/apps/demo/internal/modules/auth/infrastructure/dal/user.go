package dal

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/thumbrise/demo/golang-demo/internal/modules/shared/database"
)

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*User, error) {
	sql := "SELECT id, name, email, created_at FROM users WHERE id = $1 LIMIT 1"

	rows, err := r.db.Pool().Query(ctx, sql, id)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", database.ErrFailedQuery, err)
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, fmt.Errorf("%w: %w", database.ErrFailedCollectExactlyOneRow, err)
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	sql := "SELECT id, name, email, created_at FROM users WHERE email = $1 LIMIT 1"

	rows, err := r.db.Pool().Query(ctx, sql, email)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", database.ErrFailedQuery, err)
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, fmt.Errorf("%w: %w", database.ErrFailedCollectExactlyOneRow, err)
	}

	return &user, nil
}

func (r *UserRepository) Count(ctx context.Context) (int, error) {
	var count int

	err := r.db.Pool().QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", database.ErrFailedQuery, err)
	}

	return count, nil
}

func (r *UserRepository) Create(ctx context.Context, user *User) error {
	sql := "INSERT INTO users (name, email, created_at) VALUES ($1, $2, $3) RETURNING id"

	var id int
	if err := r.db.Pool().QueryRow(ctx, sql, user.Name, user.Email, user.CreatedAt).Scan(&id); err != nil {
		return fmt.Errorf("%w: %w", database.ErrFailedQuery, err)
	}

	user.ID = id

	return nil
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	sql := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)"

	var exists bool
	if err := r.db.Pool().QueryRow(ctx, sql, email).Scan(&exists); err != nil {
		return false, fmt.Errorf("%w: %w", database.ErrFailedQuery, err)
	}

	return exists, nil
}
