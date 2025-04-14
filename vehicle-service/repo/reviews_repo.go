package repo

import (
	"vehicles/vehicles"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ReviewsRepo interface {
	vehicles.ReviewsRepo
}

type reviewsRepo struct {
	table string
	db    *sqlx.DB
	psql  sq.StatementBuilderType
}

func NewReviewsRepo(db *DB) ReviewsRepo {
	return &reviewsRepo{
		table: "reviews",
		db:    db.Db,
		psql:  db.psql,
	}
}
