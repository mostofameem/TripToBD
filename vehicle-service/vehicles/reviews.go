package vehicles

import (
	"context"
	"vehicles/types"
)

func (svc *service) AddReviews(ctx context.Context, review *types.Reviews) error {
	err := svc.reviewsRepo.AddReviews(ctx, review)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) GetReviewsById(ctx context.Context, id int) (*[]types.Reviews, error) {
	reviews, err := svc.reviewsRepo.GetReviewsByVehicleId(ctx, id)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
