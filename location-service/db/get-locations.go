package db

import (
	"context"
	"log/slog"
	"post-service/web/utils"

	sq "github.com/Masterminds/squirrel"
)

func (repo *LocationTypeRepo) GetLocations(ctx context.Context, params utils.PaginationParams) ([]Location, error) {
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
		slog.Error("failed to execute query")
		return nil, err
	}
	defer rows.Close()

	var locations []Location
	for rows.Next() {
		var location Location
		if err := rows.Scan(
			&location.ID,
			&location.Title,
			&location.Content,
			&location.BestTime,
			&location.PictureUrl,
			&location.Rating,
			&location.VoteCount,
			&location.CreatedAt,
			&location.UpdatedAt,
			&location.Isactive,
		); err != nil {
			slog.Error("failed to scan row")
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		slog.Error("rows iteration error")
		return nil, err
	}

	return locations, nil
}

func ConfigPageSize(page, limit int) (int, int) {
	PageLimit := 20
	Offset := 0

	PageLimit = min(limit, PageLimit)
	Offset = PageLimit * page

	return PageLimit, Offset
}
