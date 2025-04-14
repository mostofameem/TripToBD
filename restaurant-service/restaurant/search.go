package restaurant

import (
	"context"
	"restaurant-service/db"
)

func (svc *service) Search(ctx context.Context, title string, location string) (*[]db.Restaurant, error) {
	restaurants, err := svc.dbRestaurantTypeRepo.Search(ctx, title, location)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}
