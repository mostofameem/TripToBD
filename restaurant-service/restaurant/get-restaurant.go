package restaurant

import (
	"context"
	"log/slog"
	"restaurant-service/db"
	"restaurant-service/logger"
	"restaurant-service/mongodb"
	"restaurant-service/web/utils"
)

func (svc *service) GetRestaurant(ctx context.Context, id int) (*mongodb.Restaurant, error) {
	// locationId, err := svc.dbRestaurantTypeRepo.GetRestaurantID(ctx, title)
	// if err != nil {
	// 	return nil, err
	// }

	locations, err := svc.mdbRestaurantTypeRepo.GetRestaurant(ctx, id)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (svc *service) GetRestaurants(ctx context.Context, filter utils.PaginationParams) (*[]db.Restaurant, error) {
	locations, err := svc.dbRestaurantTypeRepo.GetRestaurants(ctx, filter)
	if err != nil {
		slog.Error("faild to get data from psql db", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return nil, err
	}
	return locations, nil
}
