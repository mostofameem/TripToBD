package mongodb

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Comment struct {
	Userid     int
	UserName   string
	Content    string
	Created_at time.Time
	Updated_at time.Time
	Vote       int
}

func (repo *RestaurantTypeRepo) AddReviews(ctx context.Context, restaurantId string, cmnt Comment) error {
	postKey := GetKey(restaurantId)

	filter := bson.M{"_id": postKey}

	update := bson.M{
		"$push": bson.M{"comments": cmnt},
	}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Error("Error adding comment")
		return err
	}

	return nil
}
