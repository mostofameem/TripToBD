package repo

import (
	"vehicles/vehicles"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type VehiclesRepo interface {
	vehicles.VehiclesRepo
}

type vehiclesRepo struct {
	table string
	db    *sqlx.DB
	psql  sq.StatementBuilderType
}

func NewVehiclesRepo(db *DB) VehiclesRepo {
	return &vehiclesRepo{
		table: "vehicles",
		db:    db.Db,
		psql:  db.psql,
	}
}
