package db

import (
	"database/sql"
	"log/slog"
	"post-service/logger"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		slog.Error("failed to migrate ", logger.Extra(map[string]any{
			"dir ":       "./migrations",
			"err":        err.Error(),
			"migrations": migrations,
		}))
		return
	}
	slog.Info("Successfully migrate database")
}
