package handlers

import (
	"restaurant-service/config"
	"restaurant-service/restaurant"
)

type Handlers struct {
	cnf     *config.Config
	restSvc restaurant.Service
}

func NewHandlers(cnf *config.Config, restSvc restaurant.Service) *Handlers {
	return &Handlers{
		cnf:     cnf,
		restSvc: restSvc,
	}
}
