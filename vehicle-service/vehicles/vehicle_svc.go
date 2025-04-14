package vehicles

import (
	"vehicles/config"
)

type service struct {
	cnf          *config.Config
	VehiclesRepo VehiclesRepo
	routesRepo   RoutesRepo
	picsRepo     PicsRepo
	reviewsRepo  ReviewsRepo
}

func NewService(cnf *config.Config, VehiclesRepo VehiclesRepo, routesRepo RoutesRepo, picsRepo PicsRepo, reviewsRepo ReviewsRepo) Service {
	return &service{
		cnf:          cnf,
		VehiclesRepo: VehiclesRepo,
		routesRepo:   routesRepo,
		picsRepo:     picsRepo,
		reviewsRepo:  reviewsRepo,
	}
}
