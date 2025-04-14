package mongodb

import (
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *LocationTypeRepo) AddBranches(
	ctx context.Context,
	locationId int,
	restaurentId int,
) error {
	postKey := GetKey(locationId)
	filter := bson.M{"_id": postKey}

	update := bson.M{
		"$push": bson.M{"restaurants": restaurentId},
	}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Error("Error adding comment")
		return err
	}

	slog.Info("Comment added successfully")
	return nil
}
