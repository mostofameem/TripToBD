package repo

import (
	"context"
	"location-service/location"
	"location-service/logger"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	pictureCollectionName = "pictures"
)

type PictureRepo interface {
	location.PictureRepo
}

type pictureRepo struct {
	DB *mongo.Database
}

func NewPictureRepo(dB *mongo.Database) PictureRepo {
	return &pictureRepo{
		DB: dB,
	}
}

func (r *pictureRepo) Add(ctx context.Context, locationId string, urls []string) (*string, error) {
	res, err := r.DB.Collection(pictureCollectionName).InsertOne(
		ctx,
		bson.M{
			"location_id": locationId,
			"ulrs":        urls,
		},
	)
	if err != nil {
		slog.ErrorContext(ctx, "failed to add location", logger.Extra(map[string]any{
			"error":      err.Error(),
			"locationId": locationId,
		}))
		return nil, err
	}

	insertedOid := res.InsertedID.(primitive.ObjectID).Hex()
	return &insertedOid, nil
}
func (r *pictureRepo) Get(ctx context.Context, locationId string) (*[]string, error) {
	return nil, nil
}
