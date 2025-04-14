package db

import (
	"database/sql"
	"post-service/config"
	"sync"
	"time"
)

type LocationTypeRepo struct {
	db    *sql.DB
	table string
}

var locationTypeRepo *LocationTypeRepo

var locationCntOnce = sync.Once{}

func NewLocationTypeRepo(cnf *config.DBConfig) *LocationTypeRepo {

	locationCntOnce.Do(func() {
		db := NewDB(cnf)
		locationTypeRepo = &LocationTypeRepo{
			db:    db,
			table: "locations",
		}
		MigrateDB(db)
	})
	return locationTypeRepo
}

type Location struct {
	ID         int       `db:"id"`
	Title      string    `db:"title" json:"title" validate:"required"`
	Content    string    `db:"content" json:"content"`
	BestTime   string    `db:"best_time" json:"best_time" validate:"required"`
	PictureUrl string    `db:"picture_url" json:"picture_url" validate:"required"`
	Rating     float32   `db:"rating" json:"rating"`
	VoteCount  int32     `db:"vote_count" json:"vote_count"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Isactive   bool      `db:"is_active"`
}
