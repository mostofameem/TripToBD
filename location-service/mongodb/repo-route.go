package mongodb

import (
	"post-service/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type RouteTypeRepo struct {
	schema     string
	collection *mongo.Collection
}

const (
	routeCollectionName = "routes"
)

var routeTypeRepo *RouteTypeRepo

var cntRouteOnce = sync.Once{}

func NewRouteTypeRepo(cnf *config.MongoDBConfig) *RouteTypeRepo {
	mongodb := NewMongoDB(cnf)

	return &RouteTypeRepo{
		schema:     schema,
		collection: mongodb.Database.Collection(routeCollectionName),
	}
}
func GetRouteTypeRepo() *RouteTypeRepo {
	cntRouteOnce.Do(func() {
		routeTypeRepo = NewRouteTypeRepo(&config.GetConfig().MongoDB)
	})
	return routeTypeRepo
}
