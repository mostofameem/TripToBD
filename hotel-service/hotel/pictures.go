package hotel

import (
	"context"
	"hotel-service/types"
)

func (svc *service) AddPics(ctx context.Context, params *types.Pictures) error {
	err := svc.picsRepo.AddPics(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
func (svc *service) GetPicsById(ctx context.Context, vehicleId int) (*[]string, error) {
	pics, err := svc.picsRepo.GetPics(ctx, vehicleId)
	if err != nil {
		return nil, nil
	}
	return pics, nil
}
