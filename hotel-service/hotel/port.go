package hotel

import (
	"context"
	"hotel-service/types"
	"hotel-service/web/utils"
)

type Service interface {
	AddHotel(ctx context.Context, params *types.Vehicles) error
	GetHotelById(ctx context.Context, id int) (*types.Vehicles, error)
	GetHotels(ctx context.Context, filterParams *utils.PaginationParams) (*[]types.Vehicles, error)
	UpdateHotelById(ctx context.Context, params *types.Vehicles) error

	AddReview(ctx context.Context, review *types.Reviews) error
	GetReviewsById(ctx context.Context, id int) (*[]types.Reviews, error)

	AddPics(ctx context.Context, pics *types.Pictures) error
	GetPicsById(ctx context.Context, vehicleId int) (*[]string, error)
}

type HotelRepo interface {
	AddHotel(ctx context.Context, params *types.Vehicles) error
	GetHotelById(ctx context.Context, id int) (*types.Vehicles, error)
	GetHotels(ctx context.Context, filterParams *utils.PaginationParams) (*[]types.Vehicles, error)
	// UpdateVehicleById(ctx context.Context, params *types.Vehicles) error
}

type PicsRepo interface {
	AddPics(ctx context.Context, params *types.Pictures) error
	GetPics(ctx context.Context, vechicle_id int) (*[]string, error)
}

type ReviewsRepo interface {
	AddReview(ctx context.Context, review *types.Reviews) error
	GetReviews(ctx context.Context, id int) (*[]types.Reviews, error)
}
