package grpc

import (
	"context"
	"log/slog"
	pb "post-service/grpc/posts"
	"post-service/logger"
	"post-service/mongodb"
)

type PostsService struct {
	pb.UnimplementedPostServiceServer
}

func NewPostsService() *PostsService {
	return &PostsService{}
}

func (r *PostsService) AddBranches(
	ctx context.Context,
	req *pb.AddbranchesReq,
) (*pb.AddbranchesRes, error) {
	err := mongodb.GetLocationTypeRepo().AddBranches(ctx, int(req.LocationId), int(req.RestaurantId))
	if err != nil {
		slog.Error("grpc mongodb add branches failed")
		return &pb.AddbranchesRes{
			Status: false,
		}, err
	}

	return &pb.AddbranchesRes{
		Status: true,
	}, nil
}
func (r *PostsService) GetRestaurantId(
	ctx context.Context,
	req *pb.GetRestaurantIdReq,
) (*pb.GetRestaurantIdRes, error) {
	restaurants, err := mongodb.GetLocationTypeRepo().GetRestaurants(ctx, int(req.LocationId))
	if err != nil {
		slog.Error("grpc mongodb GetRestaurantId failed", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return &pb.GetRestaurantIdRes{
			Status:        false,
			RestaurantIds: nil,
		}, err
	}

	restaurantIds := make([]int32, len(restaurants))
	for i, id := range restaurants {
		restaurantIds[i] = int32(id)
	}

	return &pb.GetRestaurantIdRes{
		Status:        true,
		RestaurantIds: restaurantIds,
	}, nil
}
