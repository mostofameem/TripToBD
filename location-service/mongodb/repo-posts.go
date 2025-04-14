package mongodb

import (
	"post-service/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type LocationTypeRepo struct {
	schema     string
	collection *mongo.Collection
}

const (
	schema                 = "post"
	locationCollectionName = "locations"
)

var locationTypeRepo *LocationTypeRepo

var cntLocationOnce = sync.Once{}

func NewLocationTypeRepo(cnf *config.MongoDBConfig) *LocationTypeRepo {
	mongodb := NewMongoDB(cnf)

	return &LocationTypeRepo{
		schema:     schema,
		collection: mongodb.Database.Collection(locationCollectionName),
	}
}
func GetLocationTypeRepo() *LocationTypeRepo {
	cntLocationOnce.Do(func() {
		locationTypeRepo = NewLocationTypeRepo(&config.GetConfig().MongoDB)
	})
	return locationTypeRepo
}
