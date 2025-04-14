package location

import (
	"context"
	"log/slog"
	"post-service/db"
	"post-service/mongodb"
	"time"
)

type Location struct {
	ID         int       `db:"id"`
	Title      string    `db:"title" json:"title" validate:"required"`
	BestTime   string    `db:"best_time" json:"best_time" validate:"required"`
	Content    string    `db:"content" json:"content"`
	PictureUrl string    `db:"picture_url" json:"picture_url" validate:"required"`
	Rating     float32   `db:"rating" json:"rating"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Isactive   bool      `db:"is_active"`
}

func (svc *service) AddLocation(ctx context.Context, locationReq *Location) error {
	locationId, err := svc.dblocationTypeRepo.AddLocation(&db.Location{
		Title:      locationReq.Title,
		Content:    locationReq.Content,
		BestTime:   locationReq.BestTime,
		PictureUrl: locationReq.PictureUrl,
	})
	if err != nil {
		slog.Error("Failed to insert location data")
		return err
	}

	// userInfo, err := grpc.NewGRPC(svc.cnf).GetUserName(ctx, &users.GetUserNameReq{
	// 	UserId: int32(id),
	// })
	// if err != nil {
	// 	slog.Error("failed to get username from grpc")
	// 	return err
	// }
	err = svc.mdblocationTypeRepo.AddLocation(&mongodb.Location{
		ID:           locationId,
		Title:        locationReq.Title,
		Descriptions: locationReq.Content,
		BestTime:     locationReq.BestTime,
		Author: mongodb.Author{
			UserID:   1,
			Username: "userInfo.Name",
		},
		PictureUrl: locationReq.PictureUrl,
	})
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
