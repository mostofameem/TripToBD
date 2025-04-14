package db

import (
	"log/slog"
	"restaurant-service/logger"
	"time"
)

func (repo *RestaurantTypeRepo) AddRestaurent(
	restaurant *Restaurant,
) (int, error) {
	columns := map[string]interface{}{
		"title":         restaurant.Title,
		"picture_url":   restaurant.PictureUrl,
		"location_info": restaurant.LocationInfo,
		"rating":        0.0,
		"vote_count":    0,
		"created_at":    time.Now(),
		"updated_at":    time.Now(),
		"is_active":     false,
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
		return -1, err
	}

	index := -1
	err = repo.db.QueryRow(query, args...).Scan(&index)
	if err != nil {
		slog.Error("failed to execute query", logger.Extra(map[string]any{
			"query": query,
			"args":  args,
			"err":   err,
		}))
		return -1, err
	}

	return index, nil
}
