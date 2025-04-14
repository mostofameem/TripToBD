package restaurant

import (
	"context"
	"restaurant-service/mongodb"
	"strconv"
	"time"
)

type Comment struct {
	Userid     int
	UserName   string
	Content    string `json:"content" validation:"required"`
	Created_at time.Time
	Updated_at time.Time
	Vote       int
}

func (svc *service) AddReviews(ctx context.Context, restaurantId int, cmnt Comment) error {
	// userInfo, err := svc.grpcUserClient.GetUserName(ctx, &users.GetUserNameReq{
	// 	UserId: int32(id),
	// })
	// if err != nil {
	// 	slog.Error("failed to get username from grpc")
	// 	return err
	// }

	err := svc.mdbRestaurantTypeRepo.AddReviews(ctx, strconv.Itoa(restaurantId), mongodb.Comment{
		Userid:     cmnt.Userid,
		UserName:   "mostofa", //"userInfo.Name",
		Content:    cmnt.Content,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Vote:       0,
	})
	return err
}
