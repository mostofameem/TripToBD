package restaurant

import (
	"restaurant-service/config"
	"restaurant-service/db"
	"restaurant-service/mongodb"
)

type service struct {
	cnf                   *config.Config
	dbRestaurantTypeRepo  *db.RestaurantTypeRepo
	mdbRestaurantTypeRepo *mongodb.RestaurantTypeRepo
}

func NewService(cnf *config.Config) Service {
	restaurantTypeRepo := db.NewLocationTypeRepo(&cnf.DB)
	mongoSvc := mongodb.NewLocationTypeRepo(&cnf.MongoDB)
	return &service{
		cnf:                   cnf,
		dbRestaurantTypeRepo:  restaurantTypeRepo,
		mdbRestaurantTypeRepo: mongoSvc,
	}
}
