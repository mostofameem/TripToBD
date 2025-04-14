package mongodb

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *LocationTypeRepo) GetLocation(ctx context.Context, locationId int) (*Location, error) {
	
	searchTerm := GetKey(locationId)

	filter := bson.M{"_id": searchTerm}

	data := repo.collection.FindOne(ctx, filter)
	if err := data.Err(); err != nil {
		slog.Error("Failed to get location", "error", err)
		return nil, err
	}

	var location Location
	if err := data.Decode(&location); err != nil {
		slog.Error("Failed to decode location", "error", err)
		return nil, err
	}

	slog.Info("Data retrieved successfully")
	return &location, nil
}
