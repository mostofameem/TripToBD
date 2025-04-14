package mongodb

import (
	"log"
	"os"
	"post-service/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var cntOnce = sync.Once{}

var mongoDb *MongoDB

func NewMongoDB(mongoCnf *config.MongoDBConfig) *MongoDB {
	cntOnce.Do(func() {
		mongoDb = Connet(mongoCnf)
	})

	return mongoDb
}

func Connet(mongoCnf *config.MongoDBConfig) *MongoDB {
	mongoDB := connect(mongoCnf)
	if mongoDB == nil {
		log.Println("MongoDb Connection is nil")
		os.Exit(1)
	}

	log.Println("Connected to Mongo Database")
	return mongoDB
}
