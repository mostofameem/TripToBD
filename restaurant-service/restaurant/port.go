package restaurant

import (
	"context"
	"restaurant-service/db"
	"restaurant-service/mongodb"
	"restaurant-service/web/utils"
)

type Service interface {
	AddRestaurent(ctx context.Context, location *Restaurant) error
	GetRestaurant(ctx context.Context, id int) (*mongodb.Restaurant, error)
	GetRestaurants(ctx context.Context, filter utils.PaginationParams) (*[]db.Restaurant, error)
	GetRestaurantsByLocation(ctx context.Context, locationId int, filter utils.PaginationParams) (*[]db.Restaurant, error)
	AddReviews(ctx context.Context, locationId int, cmnt Comment) error
	AddBranchs(ctx context.Context, locationId int, restaurantId int) error
	Search(ctx context.Context, title string, location string) (*[]db.Restaurant, error)
	//UpdateRestaurants(ctx context.Context, resInfo UpdateRequest) error
}
