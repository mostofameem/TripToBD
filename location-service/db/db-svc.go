package db

import (
	"database/sql"
	"post-service/config"
	"sync"
)

var cntOnce = sync.Once{}

var db *sql.DB

func NewDB(dbCnf *config.DBConfig) *sql.DB {
	cntOnce.Do(func() {
		db = connect(dbCnf)
	})
	return db
}
