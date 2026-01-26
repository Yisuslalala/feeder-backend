package services

import (
	"context"
	"errors"
	"fmt"

	"feeder-backend/internal/models"
	"feeder-backend/internal/repositories"
	"feeder-backend/internal/utils"
)

type UserService interface {
	RegisterMember(ctx context.Context, email, password string) (*models.User, error)
	RegisterAdmin(ctx context.Context, email, password string) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService {
		userRepo: userRepo,
	}
}

func (s *userService) RegisterMember(ctx context.Context, email, password string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	if password == "" {
		return nil, errors.New("password is required")
	}

	existing, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("existing user: %w", err)
	}

	if existing != nil {
		return nil, errors.New("email already in use")
	}
	
	// hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := &models.User {
		Email: email,
		Password: hashedPassword,
		Role: "member",
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("create member: %w", err)
	}

	return user, nil
}

func (s *userService) RegisterAdmin(ctx context.Context, email, password string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}

	if password == "" {
		return nil, errors.New("password is required")
	}

	existing, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("existing user: %w", err)
	}

	if existing != nil {
		return nil, errors.New("email already in use")
	}
	
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := &models.User {
		Email: email,
		Password: hashedPassword,
		Role: "admin",
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("create admin: %w", err)
	}

	return user, nil
}
