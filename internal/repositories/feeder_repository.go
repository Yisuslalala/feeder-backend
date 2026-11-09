package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"feeder-backend/internal/models"
)

type FeederRepository interface {
	Create(ctx context.Context, feeder *models.Feeder) error
}

type feederRepository struct {
	db *sql.DB
}

func NewFeederRepository(db *sql.DB) FeederRepository {
	return &feederRepository{
		db: db,
	}
}

func (r *feederRepository) Create(ctx context.Context, feeder *models.Feeder) error {
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO feeders (house_id, mac_address, name, pet_type)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		feeder.HouseID,
		feeder.MacAddress,
		feeder.Name,
		feeder.PetType,
	).Scan(&feeder.ID)

	if err != nil {
		return fmt.Errorf("Create feeder: %w", err)
	}

	return nil
}
