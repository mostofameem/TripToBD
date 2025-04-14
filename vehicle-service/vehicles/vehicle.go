package vehicles

import (
	"context"
	"vehicles/types"
	"vehicles/web/utils"
)

func (svc *service) AddVehicle(ctx context.Context, params *types.Vehicles) error {
	err := svc.VehiclesRepo.AddVehicle(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) GetVehicleById(ctx context.Context, id int) (*types.Vehicles, error) {
	vehiche, err := svc.VehiclesRepo.GetVehicleById(ctx, id)
	if err != nil {
		return nil, err
	}
	return vehiche, nil
}

func (svc *service) GetVehicles(ctx context.Context, filter *utils.PaginationParams) (*[]types.Vehicles, error) {
	vehiches, err := svc.VehiclesRepo.GetVehicles(ctx, filter)
	if err != nil {
		return nil, err
	}
	return vehiches, nil
}

func (svc *service) UpdateVehicleById(ctx context.Context, params *types.Vehicles) error {

	return nil
}
