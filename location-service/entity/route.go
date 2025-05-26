package entity

type Route struct {
	From        string  `json:"from"          bson:"from"`
	To          string  `json:"to"            bson:"to"`
	VehicleType string  `json:"vehicleType"   bson:"vehicle_type"`
	CreatedBy   int     `json:"createdBy"     bson:"created_by"`
	Cost        float32 `json:"cost"          bson:"cost"`
}

type RouteInfo struct {
	LocationId int
	Route      []Route
}
