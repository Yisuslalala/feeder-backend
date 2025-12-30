package services

import (
	"context"
	"errors"
	"fmt"

	"feeder-backend/internal/models"
 	"feeder-backend/internal/repositories"
)

// Defines what the service does
type FeederService interface {
	RegisterFeeder(ctx context.Context, feeder *models.Feeder) error
}

type feederService struct {
	repo repositories.FeederRepository
}

func NewFeederService(repo repositories.FeederRepository) FeederService {
	return &feederService{
		repo: repo,
	}
}

func (s *feederService) RegisterFeeder(ctx context.Context, feeder *models.Feeder) error  {
	if feeder.HouseID == 0 {
		return errors.New("house_id is required")
	}

	if feeder.MacAddress == "" {
		return errors.New("mac_address is required")
	}
		
	if err := s.repo.Create(ctx, feeder); err != nil {
		return fmt.Errorf("register feeder: %w", err)
	}

	return nil
}

