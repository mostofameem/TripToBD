package db

import (
	"context"
	"log/slog"
	"restaurant-service/logger"

	sq "github.com/Masterminds/squirrel"
)

func (repo *RestaurantTypeRepo) GetRestaurantID(ctx context.Context, title string) (int, error) {
	query, args, err := NewQueryBuilder().
		Select("id").
		From(repo.table).
		Where(sq.Eq{"title": title}).ToSql()
	if err != nil {
		slog.Error("Failed to create fetch query")
		return -1, err
	}

	var id int
	err = repo.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		slog.Error("failed to execute query", logger.Extra(map[string]any{
			"query": query,
			"args":  args,
			"err":   err,
		}))
		return -1, err
	}
	return id, nil
}
