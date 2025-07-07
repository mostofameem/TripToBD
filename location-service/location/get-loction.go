package location

// import (
// 	"context"
// 	"location-service/db"
// 	"location-service/mongodb"
// 	"location-service/web/utils"
// )

// func (svc *service) GetLocation(ctx context.Context, id int) (*mongodb.Location, error) {
// 	locations, err := svc.mdblocationTypeRepo.GetLocation(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return locations, nil
// }

// func (svc *service) GetLocations(ctx context.Context, filter utils.PaginationParams) (*[]db.Location, error) {
// 	locations, err := svc.dblocationTypeRepo.GetLocations(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &locations, nil
// }
