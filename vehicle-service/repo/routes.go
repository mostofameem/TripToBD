package repo

import (
	"context"
	"log/slog"
	"vehicles/types"
	"vehicles/web/utils"

	sq "github.com/Masterminds/squirrel"
)

func (repo *routesRepo) AddRoute(ctx context.Context, route *types.Routes) error {
	columns := map[string]interface{}{
		"vehicle_id": route.VehicleId,
		"src":        "dhaka",
		"dest":       route.Dest,
		"catagory":   route.Catagory,
		"cost":       route.Cost,
		"is_active":  false,
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

func (repo *routesRepo) GetRoutesByVehicleId(
	ctx context.Context,
	params *types.GetRoutesByVehicleParams,
) (*[]types.Routes, error) {
	limit, offset := utils.ConfigPageSize(params.Page, params.Limit)

	Query := NewQueryBuilder().
		Select("*").
		From(repo.table).
		Where(sq.Eq{"vehicle_id": params.Vehicle_id})
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

	var routes []types.Routes
	for rows.Next() {
		var route types.Routes
		if err := rows.Scan(
			&route.Id,
			&route.VehicleId,
			&route.Src,
			&route.Dest,
			&route.Catagory,
			&route.Cost,
			&route.Isactive,
		); err != nil {
			slog.Error("failed to scan row", "error", err)
			return nil, err
		}
		routes = append(routes, route)
	}

	if err = rows.Err(); err != nil {
		slog.Error("row iteration error", "error", err)
		return nil, err
	}

	return &routes, nil
}

func (repo *routesRepo) GetRoutesByLocation(
	ctx context.Context,
	params *types.GetGetRoutesByLocationParams,
) (*[]types.Routes, error) {
	limit, offset := utils.ConfigPageSize(params.Page, params.Limit)

	Query := NewQueryBuilder().
		Select("*").
		From(repo.table)
	Query = Query.Where(sq.Eq{"src": "dhaka"})
	if params.Dest != "" {
		Query = Query.Where(sq.Eq{"dest": params.Dest})
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

	var routes []types.Routes
	for rows.Next() {
		var route types.Routes
		if err := rows.Scan(
			&route.Id,
			&route.VehicleId,
			&route.Src,
			&route.Dest,
			&route.Catagory,
			&route.Cost,
			&route.Isactive,
		); err != nil {
			slog.Error("failed to scan row", "error", err)
			return nil, err
		}
		routes = append(routes, route)
	}

	if err = rows.Err(); err != nil {
		slog.Error("row iteration error", "error", err)
		return nil, err
	}

	return &routes, nil
}
