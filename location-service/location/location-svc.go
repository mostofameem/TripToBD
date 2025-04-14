package location

import (
	"post-service/config"
	"post-service/db"
	"post-service/mongodb"
)

type service struct {
	cnf                 *config.Config
	dblocationTypeRepo  *db.LocationTypeRepo
	mdblocationTypeRepo *mongodb.LocationTypeRepo
}

func NewService(cnf *config.Config) Service {
	locationTypeRepo := db.NewLocationTypeRepo(&cnf.DB)
	mongoSvc := mongodb.NewLocationTypeRepo(&cnf.MongoDB)
	return &service{
		cnf:                 cnf,
		dblocationTypeRepo:  locationTypeRepo,
		mdblocationTypeRepo: mongoSvc,
	}
}
