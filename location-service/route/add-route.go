package route

import (
	"context"
	"log/slog"
	"post-service/logger"
	"post-service/mongodb"
	"time"
)

type Routes struct {
	From    string  `json:"from" validate:"required"`
	To      string  `json:"to" validate:"required"`
	Vehicle string  `json:"vehicle" validate:"required"`
	Cost    float32 `json:"cost" validate:"cost"`
}
type RouteInfo struct {
	LocationId int
	Route      [][]Routes
	//Author     mongodb.Author
	Reviews    int
	Created_at time.Time
	Updated_at time.Time
	IsActive   bool
}

func (svc *service) AddRoutes(ctx context.Context, req *RouteInfo) error {
	//req.Author.UserID = 1
	//req.Author.Username = "mostofa"
	// req.Created_at = time.Now()
	// req.Updated_at = time.Now()
	// req.Reviews = 0
	// req.IsActive = false

	var convertedRoutes [][]mongodb.Routes
	for _, routeGroup := range req.Route {
		var convertedGroup []mongodb.Routes
		for _, route := range routeGroup {
			convertedGroup = append(convertedGroup, mongodb.Routes{
				From:    route.From,
				To:      route.To,
				Vehicle: route.Vehicle,
				Cost:    route.Cost,
			})
		}
		convertedRoutes = append(convertedRoutes, convertedGroup)
	}

	err := svc.routeTypeRepo.AddRoutes(ctx, req.LocationId, &mongodb.RouteInfo{
		Route: convertedRoutes,
		//Author:     req.Author,
		// Reviews:    req.Reviews,
		// Created_at: req.Created_at,
		// Updated_at: req.Updated_at,
		// IsActive:   req.IsActive,
	})
	if err != nil {
		slog.Error("Failed to insert routes in db", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}
	return nil
}
