package db

import (
	"context"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

func (repo *LocationTypeRepo) GetLocationID(ctx context.Context, title string) (int, error) {
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
		slog.Error("Failed to get location ID", "error", err)
		return -1, err
	}
	return id, nil
}
