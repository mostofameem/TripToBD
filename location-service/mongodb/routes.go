package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"post-service/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Routes struct {
	From    string  `json:"from" bson:"from" validate:"required"`
	To      string  `json:"to" bson:"to" validate:"required"`
	Vehicle string  `json:"vehicle" bson:"vehicle" validate:"required"`
	Cost    float32 `json:"cost" bson:"cost" validate:"cost"`
}

type RouteInfo struct {
	Route [][]Routes `json:"routes" bson:"routes" validate:"required"`
	//Author     Author
	Reviews    int       `json:"reviews" bson:"reviews"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
	Updated_at time.Time `json:"updated_at" bson:"updated_at"`
	IsActive   bool      `json:"is_active" bson:"is_active"`
}

func (repo *RouteTypeRepo) AddRoutes(ctx context.Context, locationId int, req *RouteInfo) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mongoDoc, err := BsonM(req)
	if err != nil {
		slog.Error("Error converting struct to BSON", logger.Extra(map[string]any{
			"err": err.Error(),
		}))
		return err
	}

	// Get the routing key
	key := GetRouteKey(locationId)

	// Prepare the update with $push to add routes
	update := bson.M{
		"$push": bson.M{
			"routes": bson.M{
				"$each": mongoDoc["routes"],
			},
		},
	}

	// Filter to find the document with the specified routing key
	filter := bson.M{"_id": key}

	// Update or insert the document
	opts := options.Update().SetUpsert(true)
	_, err = repo.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		slog.Error("Error updating document", logger.Extra(map[string]any{
			"err": err.Error(),
		}))
		return err
	}

	return nil
}

func (repo *RouteTypeRepo) GetRoutes(ctx context.Context, locationId int) (*[]RouteInfo, error) {
	routesKey := GetRouteKey(locationId)

	filter := bson.M{"_id": routesKey}

	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		slog.Error("Error fetching routes", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return nil, err
	}
	defer cursor.Close(ctx)

	var routes []RouteInfo

	for cursor.Next(ctx) {
		var route RouteInfo
		if err := cursor.Decode(&route); err != nil {
			slog.Error("Error decoding route", logger.Extra(map[string]any{
				"error": err.Error(),
			}))
			return nil, err
		}
		routes = append(routes, route)
	}

	if err := cursor.Err(); err != nil {
		slog.Error("Cursor error while fetching routes", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return nil, err
	}

	return &routes, nil
}

func GetRouteKey(id int) string {
	return fmt.Sprintf("routes:%d", id)
}
