package location

import (
	"context"
	"log/slog"
	"post-service/logger"
	"post-service/mongodb"
	"time"
)

type Comment struct {
	LocationId int    `json:"location_id" validate:"required"`
	Userid     int    `json:"user_id" validate:"required"`
	UserName   string `json:"name" validate:"required"`
	Content    string `json:"content" validate:"required"`
	Created_at time.Time
	Updated_at time.Time
	Vote       int
}

func (svc *service) AddReviews(ctx context.Context, cmnt *Comment) error {
	err := svc.mdblocationTypeRepo.AddReviews(ctx, cmnt.LocationId, mongodb.Comment{
		Userid:     cmnt.Userid,
		UserName:   cmnt.UserName,
		Content:    cmnt.Content,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Vote:       0,
	})
	if err != nil {
		slog.Error("Failed to add review in mongodb", logger.Extra(map[string]any{
			"location_id": cmnt.LocationId,
			"user_id":     cmnt.UserName,
		}))
		return err
	}
	return nil
}
