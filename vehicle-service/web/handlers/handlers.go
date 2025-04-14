package handlers

import (
	"vehicles/config"
	"vehicles/vehicles"
)

type Handlers struct {
	cnf        *config.Config
	vehicleSvc vehicles.Service
}

func NewHandlers(cnf *config.Config, vehicleSvc vehicles.Service) *Handlers {
	return &Handlers{
		cnf:        cnf,
		vehicleSvc: vehicleSvc,
	}
}
