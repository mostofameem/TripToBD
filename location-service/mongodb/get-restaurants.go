package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type restaurants struct {
	Restaurants []int `bson:"restaurants"`
}

func (svc *LocationTypeRepo) GetRestaurants(
	ctx context.Context,
	locationId int,
) ([]int, error) {
	searchTerm := bson.M{"_id": GetKey(locationId)} // Use _id as the key for the search term

	var result restaurants
	err := svc.collection.FindOne(ctx, searchTerm, options.FindOne().SetProjection(bson.M{"restaurants": 1})).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch restaurants: %w", err)
	}

	return result.Restaurants, nil
}
