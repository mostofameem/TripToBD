package handlers

import (
	"post-service/config"
	"post-service/location"
	"post-service/route"
)

type Handlers struct {
	cnf      *config.Config
	locSvc   location.Service
	routeSvc route.Service
}

func NewHandlers(cnf *config.Config, locSvc location.Service, routeSvc route.Service) *Handlers {
	return &Handlers{
		cnf:      cnf,
		locSvc:   locSvc,
		routeSvc: routeSvc,
	}
}
