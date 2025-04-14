package db

import (
	"context"
	"fmt"
	"log/slog"
	"restaurant-service/logger"

	sq "github.com/Masterminds/squirrel"
)

func (repo *RestaurantTypeRepo) Search(ctx context.Context, title string, location string) (*[]Restaurant, error) {
	queryBuilder := sq.Select("*").From(repo.table).PlaceholderFormat(sq.Dollar)

	if title != "" {
		queryBuilder = queryBuilder.Where(sq.ILike{"title": fmt.Sprintf("%%%s%%", title)})
	}
	if location != "" {
		queryBuilder = queryBuilder.Where(sq.ILike{"location_info": fmt.Sprintf("%%%s%%", location)})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		slog.Error("failed to build query", "error", err)
		return nil, err
	}

	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		slog.Error("failed to execute query", "error", logger.Extra(map[string]any{
			"query":    query,
			"args":     args,
			"title":    title,
			"location": location,
		}))
		return nil, err
	}

	defer rows.Close()

	var restaurants []Restaurant
	for rows.Next() {
		var restaurant Restaurant
		if err := rows.Scan(
			&restaurant.ID,
			&restaurant.Title,
			&restaurant.LocationInfo,
			&restaurant.PictureUrl,
			&restaurant.Rating,
			&restaurant.VoteCnt,
			&restaurant.CreatedAt,
			&restaurant.UpdatedAt,
			&restaurant.Isactive,
		); err != nil {
			slog.Error("failed to scan row", "error", err)
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}

	if err = rows.Err(); err != nil {
		slog.Error("row iteration error", "error", err)
		return nil, err
	}

	return &restaurants, nil
}
