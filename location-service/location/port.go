package location

import (
	"context"
	"location-service/entity"
	"location-service/web/utils"
)

type Service interface {
	AddLocation(ctx context.Context, req *AddLocationReq) error
	GetLocation(ctx context.Context, id string) (*entity.Location, error)
	GetLocationPage(ctx context.Context, filter utils.PaginationParams) (*[]entity.Location, error)
	//AddReviews(ctx context.Context, cmnt *Comment) error
	GetRoutes(ctx context.Context, locationId string) (*[]entity.RouteInfo, error)
	AddRoutes(ctx context.Context, req *RouteInfo) error
}

type LocationRepo interface {
	Create(ctx context.Context, req *AddLocationReq) (*string, error)
	GetPage(ctx context.Context, params *GetPageWithFilter) ([]*entity.Location, error)
	GetOne(ctx context.Context, idStr string) (*entity.Location, error)
}

type CommentRepo interface {
}

type RouteRepo interface {
	GetRoutes(ctx context.Context, locationId string) (*[]entity.Route, error)
	//AddRoutes(ctx context.Context, req *RouteInfo) error
	AddRoutes(ctx context.Context, locationId string, req *Route) error
}
