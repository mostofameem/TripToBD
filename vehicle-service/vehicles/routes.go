package vehicles

import (
	"context"
	"vehicles/types"
)

func (svc *service) AddRoute(ctx context.Context, route *types.Routes) error {
	err := svc.routesRepo.AddRoute(ctx, route)
	if err != nil {
		return err
	}

	return nil
}

func (svc *service) GetRouteByVehicleId(ctx context.Context, filterParams *types.GetRoutesByVehicleParams) (*[]types.Routes, error) {
	routes, err := svc.routesRepo.GetRoutesByVehicleId(ctx, filterParams)
	if err != nil {
		return routes, err
	}

	return routes, nil
}

func (svc *service) GetRoutesByLocation(ctx context.Context, filterParams *types.GetGetRoutesByLocationParams) (*[]types.Routes, error) {
	routes, err := svc.routesRepo.GetRoutesByLocation(ctx, filterParams)
	if err != nil {
		return routes, err
	}

	return routes, nil
}
