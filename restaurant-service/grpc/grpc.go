package grpc

import (
	"restaurant-service/config"
	"restaurant-service/restaurant"
	"sync"
)

type grpc struct {
	cnf     *config.Config
	restSvc restaurant.Service
	Wg      sync.WaitGroup
}

func NewGRPC(cnf *config.Config, restSvc restaurant.Service) *grpc {
	return &grpc{
		cnf:     cnf,
		restSvc: restSvc,
	}
}
