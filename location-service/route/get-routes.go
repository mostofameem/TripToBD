package route

import (
	"context"
	"log/slog"
	"post-service/logger"
	"post-service/mongodb"
)

func (svc *service) GetRoutes(ctx context.Context, locationId int) (*[]mongodb.RouteInfo, error) {
	routes, err := svc.routeTypeRepo.GetRoutes(ctx, locationId)
	if err != nil {
		slog.Error("Failed to get data from mongodb", logger.Extra(map[string]any{
			"req": locationId,
			"err": err.Error(),
		}))
		return nil, err
	}
	return routes, nil
}
