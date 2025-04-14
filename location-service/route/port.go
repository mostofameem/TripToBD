package route

import (
	"context"
	"post-service/mongodb"
)

type Service interface {
	GetRoutes(ctx context.Context, locationId int) (*[]mongodb.RouteInfo, error)
	AddRoutes(ctx context.Context, req *RouteInfo) error
}
