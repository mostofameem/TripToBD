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

	if locations == nil {
		slog.Error("No location found", logger.Extra(map[string]any{
			"id": id,
		}))
	}

	return locations, nil
}

func (svc *service) GetLocationPage(ctx context.Context, req utils.PaginationParams) (*[]LocationPageResponse, error) {
	var (
		title    string
		bestTime string
	)
	if req.Filters != nil {
		if v, ok := req.Filters["title"]; ok && v != nil {
			if s, ok := v.(string); ok {
				title = s
			}
		}
		if v, ok := req.Filters["bestTime"]; ok && v != nil {
			if s, ok := v.(string); ok {
				bestTime = s
			}
		}
	}

	locs, err := svc.locationRepo.GetPage(ctx, &GetPageWithFilter{
		Title:     &title,
		BestTime:  &bestTime,
		Page:      req.Page,
		Limit:     req.Limit,
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	})
	if err != nil {
		slog.Error("failed to get locations", logger.Extra(map[string]any{
			"error": err.Error(),
			"req":   req,
		}))
		return nil, err
	}

	return locs, nil
}
