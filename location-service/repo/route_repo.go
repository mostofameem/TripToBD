package repo

import (
	"context"
	"location-service/entity"
	"location-service/location"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	routeCollectionName = "routes"
)

type RouteRepo interface {
	location.RouteRepo
}

type routeRepo struct {
	DB *mongo.Database
}

func NewRouteRepo(dB *mongo.Database) RouteRepo {
	return &routeRepo{
		DB: dB,
	}
}

func (repo *routeRepo) AddRoutes(ctx context.Context, locationId string, req *location.Route) error {
	objID, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		slog.Error("Invalid location ID", "error", err)
		return err
	}

	collection := repo.DB.Collection(routeCollectionName)

	filter := bson.M{"location_id": objID}
	update := bson.M{
		"$push": bson.M{
			"routes": req,
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		slog.Error("Failed to add route", "error", err)
		return err
	}

	slog.Info("Route added successfully", "location_id", locationId)
	return nil
}

func (repo *routeRepo) GetRoutes(ctx context.Context, locationId string) (*[]entity.Route, error) {
	objID, err := primitive.ObjectIDFromHex(locationId)
	if err != nil {
		slog.Error("Invalid location ID", "error", err)
		return nil, err
	}

	collection := repo.DB.Collection(routeCollectionName)

	filter := bson.M{"location_id": objID}

	var result struct {
		Routes []entity.Route `bson:"routes"`
	}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		slog.Error("Failed to fetch routes", "error", err)
		return nil, err
	}

	return &result.Routes, nil
}
