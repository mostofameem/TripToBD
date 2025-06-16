package repo

import (
	"context"
	"database/sql"
	"errors"
	"hotel-service/hotel"
	"hotel-service/logger"
	"hotel-service/types"
	"hotel-service/web/utils"
	"log/slog"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type HotelsRepo interface {
	hotel.HotelRepo
}

type hotelRepo struct {
	table string
	db    *sqlx.DB
	psql  sq.StatementBuilderType
}

func NewHotelRepo(db *DB) HotelsRepo {
	return &hotelRepo{
		table: "hotels",
		db:    db.Db,
		psql:  db.psql,
	}
}

func (repo *hotelRepo) AddHotel(ctx context.Context, params *types.Vehicles) error {
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

func (repo *hotelRepo) GetHotelById(ctx context.Context, id int) (*types.Vehicles, error) {
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

func (repo *hotelRepo) GetHotels(ctx context.Context, params *utils.PaginationParams) (*[]types.Vehicles, error) {
	limit, offset := utils.ConfigPageSize(params.Page, params.Limit)

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

	var vehicles []types.Vehicles
	for rows.Next() {
		var vehicle types.Vehicles
		if err := rows.Scan(
			&vehicle.ID,
			&vehicle.Name,
			&vehicle.Catagory,
			&vehicle.Descriptions,
			&vehicle.ProfilePictureUrl,
			&vehicle.Rating,
			&vehicle.VoteCnt,
			&vehicle.CreatedAt,
			&vehicle.UpdatedAt,
			&vehicle.Isactive,
		); err != nil {
			slog.Error("failed to scan row", "error", err)
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}

	if err = rows.Err(); err != nil {
		slog.Error("row iteration error", "error", err)
		return nil, err
	}

	return &vehicles, nil
}

// func (repo *BusTypeRepo) GetbusByLocation(
// 	ctx context.Context,
// 	busIds []int,
// 	params utils.PaginationParams,
// ) (*[]Bus, error) {
// 	limit, offset := ConfigPageSize(params.Page, params.Limit)

// 	Query := NewQueryBuilder().
// 		Select("*").
// 		From(repo.table).
// 		Where(sq.Eq{"id": busIds})

// 	for k, v := range params.Filters {
// 		Query = Query.Where(sq.Eq{k: v})
// 	}

// 	Query = Query.Limit(uint64(limit)).
// 		Offset(uint64(offset))

// 	query, args, err := Query.ToSql()
// 	if err != nil {
// 		slog.Error("failed to build query")
// 		return nil, err
// 	}

// 	rows, err := repo.db.QueryContext(ctx, query, args...)
// 	if err != nil {
// 		slog.Error("failed to execute query")
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var buses []Bus
// 	for rows.Next() {
// 		var bus Bus
// 		if err := rows.Scan(
// 			&bus.ID,
// 			&bus.Title,
// 			&bus.LocationInfo,
// 			&bus.PictureUrl,
// 			&bus.Rating,
// 			&bus.VoteCnt,
// 			&bus.CreatedAt,
// 			&bus.UpdatedAt,
// 			&bus.Isactive,
// 		); err != nil {
// 			slog.Error("failed to scan row", "error", err)
// 			return nil, err
// 		}
// 		buses = append(buses, bus)
// 	}

// 	if err = rows.Err(); err != nil {
// 		slog.Error("row iteration error", "error", err)
// 		return nil, err
// 	}

// 	return &buses, nil
// }
