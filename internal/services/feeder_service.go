package services

import (
	"context"
	"feeder-backend/internal/models"
 	"feeder-backend/internal/repositories/feeder_repository"
)

// Defines what the service does
type FeederService interface {
	RegisterFeeder(ctx context.Context, feeder *models.Feeder) error
}

type feederService struct {
	repo feeder_repository.FeederRepository
}

type NewFeederService(repo FeederRepository) FeederService {
	return &FeederRepository{
		repo: repo
	}
}




