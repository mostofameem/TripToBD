package location

import (
	"context"
	"location-service/config"
	"location-service/entity"
	"location-service/logger"
	"location-service/web/utils"
	"log/slog"
)

type service struct {
	cnf          *config.Config
	locationRepo LocationRepo
	routeRepo    RouteRepo
}

func NewService(cnf *config.Config, locationRepo LocationRepo, routeRepo RouteRepo) Service {
	return &service{
		cnf:          cnf,
		locationRepo: locationRepo,
		routeRepo:    routeRepo,
	}
}

func (svc *service) AddLocation(ctx context.Context, req *AddLocationReq) error {
	_, err := svc.locationRepo.Create(ctx, req)
	if err != nil {
		slog.Error("Failed to insert location data", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}

	return nil
}

func (svc *service) GetLocation(ctx context.Context, id string) (*entity.Location, error) {
	locations, err := svc.locationRepo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (svc *service) GetLocationPage(ctx context.Context, req utils.PaginationParams) (*[]entity.Location, error) {

	return nil, nil
}
