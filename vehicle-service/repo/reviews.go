package repo

import (
	"context"
	"log/slog"
	"time"
	"vehicles/types"

	sq "github.com/Masterminds/squirrel"
)

func (repo *reviewsRepo) AddReviews(ctx context.Context, review *types.Reviews) error {
	columns := map[string]interface{}{
		"vehicle_id":  review.VehicleId,
		"review":      review.Review,
		"reviewer_id": review.ReviewerId,
		"reviewed_at": time.Now(),
		"is_active":   true,
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

func (repo *reviewsRepo) GetReviewsByVehicleId(ctx context.Context, vehicleId int) (*[]types.Reviews, error) {
	Query := NewQueryBuilder().
		Select("*").
		From(repo.table)
	Query = Query.Where(sq.Eq{"vehicle_id": vehicleId})
	Query = Query.Limit(uint64(20))
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

	var reviews []types.Reviews
	for rows.Next() {
		var review types.Reviews
		if err := rows.Scan(
			&review.Id,
			&review.VehicleId,
			&review.Review,
			&review.ReviewerId,
			&review.ReviewedAt,
			&review.Isactive,
		); err != nil {
			slog.Error("failed to scan row", "error", err)
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		slog.Error("row iteration error", "error", err)
		return nil, err
	}

	return &reviews, nil
}
