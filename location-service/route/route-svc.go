package route

import (
	"post-service/mongodb"
)

type service struct {
	routeTypeRepo *mongodb.RouteTypeRepo
}

func NewRouteService() Service {
	repo := mongodb.GetRouteTypeRepo()
	return &service{
		routeTypeRepo: repo,
	}
}
