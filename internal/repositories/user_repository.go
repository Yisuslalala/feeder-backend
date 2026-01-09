package repositories

import (	
	"fmt"
	"context"
	"database/sql"
	"errors"

	"feeder-backend/internal/models"
)
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id int64) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	result, err := r.db.ExecContext(ctx, `
		INSERT INTO users (email, password, role)
		VALUES(?, ?, ?)
	`,
		user.Email, user.Password, user.Role,
	)
	
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("get last insert id: %w", err)
	}
	
	user.ID = id
	return nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, email, password, role
		FROM users
		WHERE email = ?
	`, email)
	
	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found by email")
	}

	if err != nil {
		return nil, fmt.Errorf("find user by email: %w", err)
	}

	return &user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*models.User, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, email, password, role
		FROM users
		WHERE id = ?
	`, id)
	
	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found by id")
	}

	if err != nil {
		return nil, fmt.Errorf("find user by id: %w", err)
	}

	return &user, nil
}
