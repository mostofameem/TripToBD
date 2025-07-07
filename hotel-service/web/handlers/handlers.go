package handlers

import (
	"hotel-service/config"
	"hotel-service/hotel"
)

type Handlers struct {
	cnf      *config.Config
	hotelSvc hotel.Service
}

func NewHandlers(cnf *config.Config, hotelSvc hotel.Service) *Handlers {
	return &Handlers{
		cnf:      cnf,
		hotelSvc: hotelSvc,
	}
}
