package db

import (
	"context"
	"log/slog"
	"restaurant-service/logger"
	"restaurant-service/web/utils"

	sq "github.com/Masterminds/squirrel"
)

func (repo *RestaurantTypeRepo) GetRestaurants(ctx context.Context, params utils.PaginationParams) (*[]Restaurant, error) {
	limit, offset := ConfigPageSize(params.Page, params.Limit)

	Query := NewQueryBuilder().
		Select("*").
		From(repo.table)
	for k, v := range params.Filters {
		Query = Query.Where(sq.Eq{k: v})
	}
	Query = Query.Limit(uint64(limit)).
		Offset(uint64(offset))
	query, args, err := Query.ToSql()
	if err != nil {
		slog.Error("failed to build query")
		return nil, err
	}

	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		slog.Error("failed to execute query", logger.Extra(map[string]any{
			"query": query,
			"args":  args,
			"err":   err,
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
			&restaurant.PictureUrl,
			&restaurant.LocationInfo,
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

func ConfigPageSize(page, limit int) (int, int) {
	PageLimit := 20
	Offset := 0

	PageLimit = min(limit, PageLimit)
	Offset = PageLimit * page

	return PageLimit, Offset
}

func (repo *RestaurantTypeRepo) GetRestaurantByLocation(
	ctx context.Context,
	restaurantIds []int,
	params utils.PaginationParams,
) (*[]Restaurant, error) {
	limit, offset := ConfigPageSize(params.Page, params.Limit)

	Query := NewQueryBuilder().
		Select("*").
		From(repo.table).
		Where(sq.Eq{"id": restaurantIds})

	for k, v := range params.Filters {
		Query = Query.Where(sq.Eq{k: v})
	}

	Query = Query.Limit(uint64(limit)).
		Offset(uint64(offset))

	query, args, err := Query.ToSql()
	if err != nil {
		slog.Error("failed to build query")
		return nil, err
	}

	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		slog.Error("failed to execute query")
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
