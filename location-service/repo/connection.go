package repo

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"location-service/config"
	"location-service/logger"
)

type MongoDBClient struct {
	Client *mongo.Client
}

func GetMongoURI(cnf config.MongoDBConfig) string {
	uri := cnf.GetURI()
	if uri == "" {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s/%s?authSource=%s&replicaSet=%s",
			cnf.GetUser(),
			cnf.GetPassword(),
			cnf.GetURI(),
			cnf.GetDatabase(),
			cnf.GetAuthSource(),
			cnf.GetReplicaSet(),
		)
	}
	return uri
}

func (m *MongoDBClient) Close() {
	if m.Client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		slog.ErrorContext(context.Background(), "MongoDB Disconnection Error", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return
	}

	slog.Info("MongoDB Connection Closed")
}

func ConnectMongoDB(cnf config.MongoDBConfig) (*MongoDBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := GetMongoURI(cnf)

	clientOptions := options.Client().
		ApplyURI(uri).
		SetAuth(options.Credential{
			Username:   cnf.GetUser(),
			Password:   cnf.GetPassword(),
			AuthSource: cnf.GetAuthSource(),
		}).
		SetMaxPoolSize(cnf.GetMaxPoolSize()).
		SetMinPoolSize(cnf.GetMinPoolSize()).
		SetMaxConnIdleTime(time.Duration(cnf.GetMaxConnIdleTimeInMs()) * time.Millisecond)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		slog.ErrorContext(context.Background(), "MongoDB Connection Error", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		slog.ErrorContext(context.Background(), "MongoDB Ping Failed", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return nil, err
	}

	slog.Info("Connected to MongoDB", logger.Extra(map[string]any{
		"database": cnf.GetDatabase(),
	}))

	return &MongoDBClient{Client: client}, nil
}
