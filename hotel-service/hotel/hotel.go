package hotel

import (
	"context"
	"hotel-service/types"
	"hotel-service/web/utils"
)

func (svc *service) AddHotel(ctx context.Context, params *types.Vehicles) error {
	err := svc.hotelRepo.AddHotel(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) GetHotelById(ctx context.Context, id int) (*types.Vehicles, error) {
	vehiche, err := svc.hotelRepo.GetHotelById(ctx, id)
	if err != nil {
		return nil, err
	}
	return vehiche, nil
}

func (svc *service) GetHotels(ctx context.Context, filter *utils.PaginationParams) (*[]types.Vehicles, error) {
	vehiches, err := svc.hotelRepo.GetHotels(ctx, filter)
	if err != nil {
		return nil, err
	}
	return vehiches, nil
}

func (svc *service) UpdateHotelById(ctx context.Context, params *types.Vehicles) error {

	return nil
}
