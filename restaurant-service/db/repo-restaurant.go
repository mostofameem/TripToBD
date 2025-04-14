package db

import (
	"database/sql"
	"restaurant-service/config"
	"sync"
	"time"
)

type RestaurantTypeRepo struct {
	db    *sql.DB
	table string
}

var restaurantTypeRepo *RestaurantTypeRepo

var restaurantCntOnce = sync.Once{}

func NewLocationTypeRepo(cnf *config.DBConfig) *RestaurantTypeRepo {

	restaurantCntOnce.Do(func() {
		db := NewDB(cnf)
		restaurantTypeRepo = &RestaurantTypeRepo{
			db:    db,
			table: "restaurants",
		}
	})
	return restaurantTypeRepo
}

type Restaurant struct {
	ID           int       `db:"id" json:"id"`
	Title        string    `db:"title" json:"title"`
	LocationInfo string    `db:"location_info" json:"location_info"`
	PictureUrl   string    `db:"picture_url" json:"picture_url"`
	Rating       float32   `db:"rating" json:"rating"`
	VoteCnt      int       `db:"vote_count" json:"vote_count"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Isactive     bool      `db:"is_active" json:"is_active"`
}
