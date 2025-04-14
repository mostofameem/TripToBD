package location

import (
	"context"
	"post-service/db"
	"post-service/mongodb"
	"post-service/web/utils"
)

type Service interface {
	AddLocation(ctx context.Context, location *Location) error
	GetLocation(ctx context.Context, id int) (*mongodb.Location, error)
	GetLocations(ctx context.Context, filter utils.PaginationParams) (*[]db.Location, error)
	AddReviews(ctx context.Context, cmnt *Comment) error
}
