package vehicles

import (
	"context"
	"vehicles/types"
	"vehicles/web/utils"
)

type Service interface {
	AddVehicle(ctx context.Context, params *types.Vehicles) error
	GetVehicleById(ctx context.Context, id int) (*types.Vehicles, error)
	GetVehicles(ctx context.Context, filterParams *utils.PaginationParams) (*[]types.Vehicles, error)
	UpdateVehicleById(ctx context.Context, params *types.Vehicles) error

	AddRoute(ctx context.Context, route *types.Routes) error
	GetRouteByVehicleId(ctx context.Context, filterParams *types.GetRoutesByVehicleParams) (*[]types.Routes, error)
	GetRoutesByLocation(ctx context.Context, filterParams *types.GetGetRoutesByLocationParams) (*[]types.Routes, error)

	AddReviews(ctx context.Context, review *types.Reviews) error
	GetReviewsById(ctx context.Context, id int) (*[]types.Reviews, error)

	AddPics(ctx context.Context, pics *types.Pictures) error
	GetPicsById(ctx context.Context, vehicleId int) (*[]string, error)
}

type VehiclesRepo interface {
	AddVehicle(ctx context.Context, params *types.Vehicles) error
	GetVehicleById(ctx context.Context, id int) (*types.Vehicles, error)
	GetVehicles(ctx context.Context, filterParams *utils.PaginationParams) (*[]types.Vehicles, error)
	// UpdateVehicleById(ctx context.Context, params *types.Vehicles) error
}

type RoutesRepo interface {
	AddRoute(ctx context.Context, route *types.Routes) error
	GetRoutesByVehicleId(ctx context.Context, filterParams *types.GetRoutesByVehicleParams) (*[]types.Routes, error)
	GetRoutesByLocation(ctx context.Context, filterParams *types.GetGetRoutesByLocationParams) (*[]types.Routes, error)
}

type PicsRepo interface {
	AddPics(ctx context.Context, params *types.Pictures) error
	GetPics(ctx context.Context, vechicle_id int) (*[]string, error)
}

type ReviewsRepo interface {
	AddReviews(ctx context.Context, review *types.Reviews) error
	GetReviewsByVehicleId(ctx context.Context, id int) (*[]types.Reviews, error)
}
