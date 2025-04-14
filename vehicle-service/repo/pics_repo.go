package repo

import (
	"vehicles/vehicles"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PicsRepo interface {
	vehicles.PicsRepo
}

type picsRepo struct {
	table string
	db    *sqlx.DB
	psql  sq.StatementBuilderType
}

func NewPicsRepo(db *DB) PicsRepo {
	return &picsRepo{
		table: "pictures",
		db:    db.Db,
		psql:  db.psql,
	}
}
