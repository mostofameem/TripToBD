package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"restaurant-service/config"
	"time"

	_ "github.com/lib/pq"
)

func GetConnectionString(dbConf *config.DBConfig) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		dbConf.User,
		dbConf.Pass,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)
	if dbConf.EnableSSLMode {
		connectionString += " sslmode=require"
	} else {
		connectionString += " sslmode=disable"
	}
	
	return connectionString
}

func connect(dbConf *config.DBConfig) *sql.DB {
	dbSource := GetConnectionString(dbConf)

	dbCon, err := sql.Open("postgres", dbSource)
	if err != nil {
		slog.Error(fmt.Sprintf("Connection error %v", err))
		os.Exit(1)
	}

	if err := dbCon.Ping(); err != nil {
		slog.Error(fmt.Sprintf("DB ping error %v", err))
		os.Exit(1)
	}

	dbCon.SetConnMaxIdleTime(
		time.Duration(dbConf.MaxIdleTimeInMinute * int(time.Minute)),
	)

	slog.Info("Connected to postgres Database")
	return dbCon
}
