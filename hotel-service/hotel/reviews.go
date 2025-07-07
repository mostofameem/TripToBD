package hotel

import (
	"context"
	"hotel-service/types"
)

func (svc *service) AddReview(ctx context.Context, review *types.Reviews) error {
	err := svc.reviewsRepo.AddReview(ctx, review)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) GetReviewsById(ctx context.Context, id int) (*[]types.Reviews, error) {
	reviews, err := svc.reviewsRepo.GetReviews(ctx, id)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
