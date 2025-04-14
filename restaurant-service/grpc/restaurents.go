package grpc

import (
	pb "restaurant-service/grpc/restaurants"
	"restaurant-service/restaurant"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantsServiceServer
}

func NewRestaurantService(locSvc restaurant.Service) *RestaurantService {
	return &RestaurantService{}
}
