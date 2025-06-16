package repo

import (
	"context"
	"location-service/entity"
	"location-service/location"
	"location-service/logger"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	LocationCollName = "location"
)

type LocationRepo interface {
	location.LocationRepo
}

type locationRepo struct {
	DB *mongo.Database
}

func NewLocationRepo(dB *mongo.Database) LocationRepo {
	return &locationRepo{
		DB: dB,
	}
}

func (repo *locationRepo) Create(
	ctx context.Context,
	params *location.AddLocationReq,
) (*string, error) {
	insertResult, err := repo.DB.Collection(LocationCollName).InsertOne(ctx,
		bson.M{
			"title":        params.Title,
			"descriptions": params.Descriptions,
			"best_time":    params.BestTime,
			"picture_url":  params.PictureUrl,
			"rating":       params.Rating,
			"vote_count":   params.Voted,
			"created_at":   params.CreatedAt,
			"created_by":   params.CreatedBy,
			"updatedAt":    params.UpdatedAt,
		},
	)
	if err != nil {
		slog.ErrorContext(ctx, "failed to add location", logger.Extra(map[string]any{
			"error":  err.Error(),
			"params": params,
		}))
		return nil, err
	}

	insertedOid := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return &insertedOid, nil
}

func (repo *locationRepo) GetOne(ctx context.Context, idStr string) (*entity.Location, error) {
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		slog.Error("Invalid ObjectID format", "error", err)
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	data := repo.DB.Collection(LocationCollName).FindOne(ctx, filter)
	if err := data.Err(); err != nil {
		slog.Error("Failed to get location", "error", err)
		return nil, err
	}

	var location entity.Location
	if err := data.Decode(&location); err != nil {
		slog.Error("Failed to decode location", "error", err)
		return nil, err
	}

	slog.Info("Data retrieved successfully")
	return &location, nil
}

func (repo *locationRepo) GetPage(ctx context.Context, params *location.GetPageWithFilter) (*[]location.LocationPageResponse, error) {
	builder := NewMongoQueryBuilder().
		AddRegex("title", *params.Title).
		AddRegex("best_time", *params.BestTime).
		SetPagination(params.Page, params.Limit).
		SetSorting(params.SortBy, params.SortOrder).
		SetProjection("_id", "title", "picture_url", "rating")

	filter, opts := builder.Build()

	cursor, err := repo.DB.Collection(LocationCollName).Find(ctx, filter, opts)
	if err != nil {
		slog.Error("Failed to execute query", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []location.LocationPageResponse
	if err = cursor.All(ctx, &results); err != nil {
		slog.Error("Failed to decode locations", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return nil, err
	}
	return &results, nil
}
