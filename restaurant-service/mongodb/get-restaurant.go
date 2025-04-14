package mongodb

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *RestaurantTypeRepo) GetRestaurant(ctx context.Context, locationId int) (*Restaurant, error) {

	searchTerm := GetKey(fmt.Sprintf("%d", locationId))

	filter := bson.M{"_id": searchTerm}

	data := repo.collection.FindOne(ctx, filter)
	if err := data.Err(); err != nil {
		slog.Error("Failed to get location", "error", err)
		return nil, err
	}

	var restaurant Restaurant
	if err := data.Decode(&restaurant); err != nil {
		slog.Error("Failed to decode location", "error", err)
		return nil, err
	}

	slog.Info("Data retrieved successfully")
	return &restaurant, nil
}
