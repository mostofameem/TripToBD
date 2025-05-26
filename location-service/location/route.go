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

type RouteInfo struct {
	LocationId int
	Route      []Route
}

func (svc *service) AddRoutes(ctx context.Context, req *RouteInfo) error {

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
