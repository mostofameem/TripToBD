package mongodb

import (
	"restaurant-service/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantTypeRepo struct {
	schema     string
	collection *mongo.Collection
}

const restaurantCollectionName = "restaurants"

var restaurantTypeRepo *RestaurantTypeRepo

var restaurantCntOnce = sync.Once{}

func NewLocationTypeRepo(cnf *config.MongoDBConfig) *RestaurantTypeRepo {
	restaurantCntOnce.Do(func() {
		mongodb := NewMongoDB(cnf)

		restaurantTypeRepo = &RestaurantTypeRepo{
			schema:     "restaurant-service",
			collection: mongodb.Database.Collection(restaurantCollectionName),
		}
	})
	return restaurantTypeRepo
}
