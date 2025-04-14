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

func (repo *LocationTypeRepo) AddReviews(ctx context.Context, locationId int, cmnt Comment) error {
	postKey := GetKey(locationId)

	filter := bson.M{"_id": postKey}

	update := bson.M{
		"$push": bson.M{"comments": cmnt},
	}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		slog.Error("Error adding review")
		return err
	}

	return nil
}

// func (repo *LocationTypeRepo) AddLike(ctx context.Context, locationId int, cmntId Comment) error {
// 	postKey := GetKey(locationId)

// 	filter := bson.M{"_id": postKey, "comments._id": cmntId.Userid}

// 	update := bson.M{
// 		"$push": bson.M{"comments": cmntId},
// 		"$inc":  bson.M{"comments.$.vote": 1},
// 	}

// 	_, err := repo.collection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return err
// 	}

// 	slog.Info("Comment added and vote incremented successfully")
// 	return nil
// }
