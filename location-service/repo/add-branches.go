package repo

// func (repo *locationRepo) AddBranches(
// 	ctx context.Context,
// 	locationId int,
// 	restaurentId int,
// ) error {
// 	postKey := GetKey(locationId)
// 	filter := bson.M{"_id": postKey}

// 	update := bson.M{
// 		"$push": bson.M{"restaurants": restaurentId},
// 	}

// 	_, err := repo.DB.Collection("restaurants").UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		slog.Error("Error adding comment")
// 		return err
// 	}

// 	return nil
// }
