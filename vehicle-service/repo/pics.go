package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"vehicles/logger"
	"vehicles/types"

	sq "github.com/Masterminds/squirrel"
)

func (repo *picsRepo) AddPics(ctx context.Context, params *types.Pictures) error {
	columns := map[string]interface{}{
		"vehicle_id": params.VehicleId,
		"url":        params.Url,
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

func (repo *picsRepo) GetPics(ctx context.Context, vehicleId int) (*[]string, error) {
	query, args, err := repo.psql.Select("url").
		From(repo.table).
		Where(sq.Eq{"vehicle_id": vehicleId}).ToSql()
	if err != nil {
		slog.Error("failed to build query")
		return nil, err
	}
	var pics []string
	err = repo.db.SelectContext(ctx, &pics, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Info(fmt.Sprintf("No Row Found for vehichle_id %d", vehicleId))
			return nil, nil
		}
		slog.Error("failed to execute query", logger.Extra(map[string]any{
			"err":   err.Error(),
			"query": query,
			"args":  args,
		}))
		return nil, err
	}
	return &pics, nil
}
