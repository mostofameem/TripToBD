package restaurant

import (
	"context"
	"log/slog"
	"restaurant-service/db"
	"restaurant-service/grpc/clients"
	"restaurant-service/grpc/posts"
	"restaurant-service/logger"
	"restaurant-service/web/utils"
)

func (svc *service) GetRestaurantsByLocation(
	ctx context.Context,
	locationId int,
	filter utils.PaginationParams,
) (*[]db.Restaurant, error) {
	response, err := clients.GetPostsClient().GetRestaurantId(ctx, &posts.GetRestaurantIdReq{
		LocationId: int32(locationId),
	})
	if err != nil {
		slog.Error("grpc call error", logger.Extra(map[string]any{
			"location id": locationId,
			"response":    response,
		}))
		return nil, err
	}

	restaurantIds := make([]int, len(response.RestaurantIds))
	for i, id := range response.RestaurantIds {
		restaurantIds[i] = int(id)
	}

	restaurants, err := svc.dbRestaurantTypeRepo.GetRestaurantByLocation(ctx, restaurantIds, filter)
	if err != nil {
		slog.Error("GetRestaurantByLocation failed", logger.Extra(map[string]any{
			"location id": locationId,
			"response":    response,
		}))
		return nil, err
	}

	return restaurants, nil
}
