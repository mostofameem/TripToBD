package restaurant

import (
	"context"
	"log/slog"
	"restaurant-service/db"
	"restaurant-service/mongodb"
	"strconv"
	"time"
)

type Restaurant struct {
	ID           int
	Title        string
	LocationInfo string
	Descriptions string
	PictureUrl   string
	Rating       float32
	VoteCnt      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Isactive     bool
}

func (svc *service) AddRestaurent(ctx context.Context, restaurantReq *Restaurant) error {
	restaurantId, err := svc.dbRestaurantTypeRepo.AddRestaurent(&db.Restaurant{
		Title:        restaurantReq.Title,
		LocationInfo: restaurantReq.LocationInfo,
		PictureUrl:   restaurantReq.PictureUrl,
	})
	if err != nil {
		slog.Error("Failed to insert location data")
		return err
	}

	// userInfo, err := svc.grpcUserClient.GetUserName(ctx, &users.GetUserNameReq{
	// 	UserId: int32(id),
	// })
	// if err != nil {
	// 	slog.Error("failed to get username from grpc")
	// 	return err
	// }
	err = svc.mdbRestaurantTypeRepo.AddRestaurent(&mongodb.Restaurant{
		ID:           strconv.Itoa(restaurantId),
		Title:        restaurantReq.Title,
		Descriptions: restaurantReq.Descriptions,
		Author: mongodb.Author{
			UserID:   1,
			Username: "mostofa",
		},
		PictureUrl: restaurantReq.PictureUrl,
	})
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
