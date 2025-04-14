package restaurant

import (
	"context"
	"log/slog"
	"restaurant-service/grpc/clients"
	"restaurant-service/grpc/posts"
	"restaurant-service/logger"
)

func (svc *service) AddBranchs(ctx context.Context, locationId int, restaurantId int) error {
	res, err := clients.GetPostsClient().AddBranches(ctx, &posts.AddbranchesReq{
		LocationId:   int32(locationId),
		RestaurantId: int32(restaurantId),
	})
	if err != nil || !res.Status {
		slog.Error("grpc call error", logger.Extra(map[string]any{
			"location id":   locationId,
			"restaurant id": restaurantId,
			"response":      res,
		}))
		return err
	}
	return nil
}
