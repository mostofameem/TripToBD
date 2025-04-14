package types

import "time"

type Vehicles struct {
	ID                int       `db:"id"           json:"id"`
	Name              string    `db:"name"         json:"name"        validate:"required"`
	Catagory          string    `db:"catagory"     json:"catagory"    validate:"required"`
	Descriptions      string    `db:"description"  json:"description" validate:"required"`
	ProfilePictureUrl string    `db:"picture_url"  json:"picture_url"`
	Rating            float32   `db:"rating"       json:"rating"`
	VoteCnt           int       `db:"vote_count"   json:"vote_count"`
	CreatedAt         time.Time `db:"created_at"   json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"   json:"updated_at"`
	Isactive          bool      `db:"is_active"    json:"is_active"`
}

type Routes struct {
	Id        int    `db:"id"                   json:"id"`
	VehicleId int    `db:"vehicle_id"           json:"vehicle_id"   validate:"required"`
	Src       string `db:"src"                  json:"src"          validate:"required"`
	Dest      string `db:"dest"                 json:"dest"         validate:"required"`
	Catagory  string `db:"catagory"             json:"catagory"     validate:"required"`
	Cost      int    `db:"cost"                 json:"cost"         validate:"required"`
	Isactive  bool   `db:"is_active"            json:"is_active"`
}

type Pictures struct {
	Id        int    `db:"id"                   json:"id"`
	VehicleId int    `db:"vehicle_id"           json:"vehicleId"    validate:"required"`
	Url       string `db:"url"                  json:"url"`
}

type Reviews struct {
	Id         int       `db:"id"                   json:"id"`
	VehicleId  int       `db:"vehicle_id"           json:"vehicle_id"    validate:"required"`
	Review     string    `db:"review"               json:"review"        validate:"required"`
	ReviewerId int       `db:"reviewer_id"          json:"reviewer_id"   validate:"required"`
	ReviewedAt time.Time `db:"reviewed_at"          json:"reviewed_at"`
	Isactive   bool      `db:"is_active"            json:"is_active"`
}

type GetRoutesByVehicleParams struct {
	Vehicle_id int `json:"id"    validate:"required"`
	Limit      int `json:"limit"`
	Page       int `json:"page"`
}

type GetGetRoutesByLocationParams struct {
	Dest  string `json:"dest"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}
