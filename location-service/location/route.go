package location

import (
	"context"
	"location-service/entity"
	"location-service/logger"
	"log/slog"
)

type Route struct {
	From        string  `json:"from"          bson:"from"`
	To          string  `json:"to"            bson:"to"`
	VehicleType string  `json:"vehicleType"   bson:"vehicle_type"`
	CreatedBy   int     `json:"createdBy"     bson:"created_by"`
	Cost        float32 `json:"cost"          bson:"cost"`
}

func (svc *service) AddRoutes(ctx context.Context, locationId string, route *Route) error {
	err := svc.routeRepo.AddRoutes(ctx, locationId, route)
	if err != nil {
		slog.Error("Failed to add routes", logger.Extra(map[string]any{
			"err":        err.Error(),
			"locationID": locationId,
			"route":      route,
		}))
		return err
	}

	return nil
}

func (svc *service) GetRoutes(ctx context.Context, locationId string) (*[]entity.RouteInfo, error) {
	_, err := svc.routeRepo.GetRoutes(ctx, locationId)
	if err != nil {
		slog.Error("Failed to get data from mongodb", logger.Extra(map[string]any{
			"req": locationId,
			"err": err.Error(),
		}))
		return nil, err
	}
	return nil, nil
}
