package grpc

import (
	"location-service/config"
	"sync"
)

type grpc struct {
	cnf *config.Config
	Wg  sync.WaitGroup
}

func NewGRPC(cnf *config.Config) *grpc {
	return &grpc{
		cnf: cnf,
	}
}
