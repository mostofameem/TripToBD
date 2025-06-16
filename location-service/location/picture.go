package location

import "context"

type Picture struct {
	LocationId int      `json:"locationId"   bson:"location_id"`
	Urls       []string `json:"urls"         bson:"urls"`
}

func (svc *service) AddPictures(ctx context.Context, req Picture) error {

	return nil
}

func (svc *service) GetPictures(ctx context.Context, req Picture) (*Picture, error) {

	return nil, nil
}
