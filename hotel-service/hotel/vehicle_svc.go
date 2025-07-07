package hotel

import (
	"hotel-service/config"
)

type service struct {
	cnf         *config.Config
	hotelRepo   HotelRepo
	picsRepo    PicsRepo
	reviewsRepo ReviewsRepo
}

func NewService(cnf *config.Config, hotelRepo HotelRepo, picsRepo PicsRepo, reviewsRepo ReviewsRepo) Service {
	return &service{
		cnf:         cnf,
		hotelRepo:   hotelRepo,
		picsRepo:    picsRepo,
		reviewsRepo: reviewsRepo,
	}
}
