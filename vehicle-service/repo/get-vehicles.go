package repo

import (
	"context"
	"log/slog"
	"vehicles/types"
	"vehicles/web/utils"

	sq "github.com/Masterminds/squirrel"
)

func (repo *vehiclesRepo) GetVehicles(ctx context.Context, params *utils.PaginationParams) (*[]types.Vehicles, error) {
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
