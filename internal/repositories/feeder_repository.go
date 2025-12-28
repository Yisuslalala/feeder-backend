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

type FeederRepository struct {
	db * sql.DB
}

func NewFeederRepository(db *sql.DB) *FeederRepository {
	return &FeederRepository {
		db: db,
	}
}

func (r *FeederRepository) Create(ctx context.Context, feeder *models.Feeder) error {
	result, err := r.db.ExecContext(ctx, `
			INSERT INTO feeders (house_id, mac_address, name, pet_type)
			VALUES(?, ?, ?, ?)
		`,
		feeder.HouseID,
		feeder.MacAddress,
		feeder.Name,
		feeder.PetType,
	)

	if err != nil {
		return fmt.Errorf("Create feeder: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("Get last insert id: %w", err)
	}

	feeder.ID = id
	return nil
}
