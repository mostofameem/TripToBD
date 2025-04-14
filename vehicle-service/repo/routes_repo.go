package repo

import (
	"vehicles/vehicles"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type RoutesRepo interface {
	vehicles.RoutesRepo
}

type routesRepo struct {
	table string
	db    *sqlx.DB
	psql  sq.StatementBuilderType
}

func NewRoutesRepo(db *DB) RoutesRepo {
	return &routesRepo{
		table: "routes",
		db:    db.Db,
		psql:  db.psql,
	}
}

//migrate create -ext sql -dir migrations -seq create_table_routes
