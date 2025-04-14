package clients

import (
	"context"
	"log/slog"
	"restaurant-service/grpc/posts"
	"restaurant-service/logger"
	"time"
)

func (p *PostsClients) AddBranches(ctx context.Context, req *posts.AddbranchesReq) (*posts.AddbranchesRes, error) {
	slog.Info("Adding resturents to locations", logger.Extra(map[string]any{
		"payload": req,
	}))

	ctx, cancle := context.WithTimeout(ctx, time.Duration(20*time.Second))
	defer cancle()

	resp, err := p.client.AddBranches(ctx, req)
	if err != nil {
		slog.Error(err.Error(), logger.Extra(map[string]any{
			"payload":  req,
			"respomse": resp,
		}))
		return nil, err
	}

	return resp, nil
}

func (p *PostsClients) GetRestaurantId(
	ctx context.Context,
	req *posts.GetRestaurantIdReq,
) (*posts.GetRestaurantIdRes, error) {
	slog.Info("Get resturents id", logger.Extra(map[string]any{
		"payload": req,
	}))

	ctx, cancle := context.WithTimeout(ctx, time.Duration(20*time.Second))
	defer cancle()

	resp, err := p.client.GetRestaurantId(ctx, req)
	if err != nil {
		slog.Error(err.Error(), logger.Extra(map[string]any{
			"payload":  req,
			"respomse": resp,
		}))
		return nil, err
	}

	return resp, nil
}
