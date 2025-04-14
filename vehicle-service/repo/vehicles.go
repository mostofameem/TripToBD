package repo

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"
	"vehicles/logger"
	"vehicles/types"

	sq "github.com/Masterminds/squirrel"
)

func (repo *vehiclesRepo) AddVehicle(ctx context.Context, params *types.Vehicles) error {
	columns := map[string]interface{}{
		"name":        params.Name,
		"catagory":    params.Catagory,
		"description": params.Descriptions,
		"picture_url": params.ProfilePictureUrl,
		"rating":      0.0,
		"vote_count":  0,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"is_active":   false,
	}
	var colNames []string
	var colValues []any
	for colName, colValue := range columns {
		colNames = append(colNames, colName)
		colValues = append(colValues, colValue)
	}

	query, args, err := NewQueryBuilder().
		Insert(repo.table).
		Columns(colNames...).
		Values(colValues...).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		slog.Error("Failed to create insert query")
		return err
	}

	_, err = repo.db.ExecContext(ctx, query, args...)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (repo *vehiclesRepo) GetVehicleById(ctx context.Context, id int) (*types.Vehicles, error) {
	query, args, err := NewQueryBuilder().
		Select("*").
		From(repo.table).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		slog.Error("Failed to create select query", logger.Extra(map[string]any{
			"error": err.Error(),
			"query": query,
			"args":  args,
		}))
		return nil, err
	}
	var vehicleInfo types.Vehicles

	err = repo.db.QueryRowContext(ctx, query, args...).Scan(
		&vehicleInfo.ID,
		&vehicleInfo.Name,
		&vehicleInfo.Catagory,
		&vehicleInfo.Descriptions,
		&vehicleInfo.ProfilePictureUrl,
		&vehicleInfo.Rating,
		&vehicleInfo.VoteCnt,
		&vehicleInfo.CreatedAt,
		&vehicleInfo.UpdatedAt,
		&vehicleInfo.Isactive,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("No vehicle found with the provided ID", "id", id)
			return nil, nil
		}
		slog.Error("Failed to Execute query", logger.Extra(map[string]any{
			"error": err.Error(),
			"query": query,
			"args":  args,
		}))
		return nil, err
	}

	return &vehicleInfo, nil
}
