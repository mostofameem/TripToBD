package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"restaurant-service/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(db *config.MongoDBConfig) *MongoDB {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	mongodbSource := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/",
		db.User,
		db.Pass,
		db.Host,
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbSource))
	if err != nil {
		slog.Info("Mongodb connection Failed")
		return nil
	}

	database := client.Database(db.Name)

	return &MongoDB{
		Client:   client,
		Database: database,
	}
}
