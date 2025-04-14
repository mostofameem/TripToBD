package db

import (
	"log/slog"
	"restaurant-service/config"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB() {
	conf := config.GetConfig()

	migrations := &migrate.FileMigrationSource{
		Dir: conf.MigrationSource,
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	slog.Info("Successfully migrate database")
}
