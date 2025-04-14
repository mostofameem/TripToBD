package cmd

import (
	"log/slog"
	"vehicles/config"
	"vehicles/logger"
	"vehicles/repo"
	"vehicles/vehicles"
	"vehicles/web"
	"vehicles/web/handlers"
	"vehicles/web/utils"

	"github.com/spf13/cobra"
)

var serveRestCmd = &cobra.Command{
	Use:   "serve-rest",
	Short: "start a rest server",
	RunE:  serveRest,
}

func serveRest(cmd *cobra.Command, args []string) error {
	cnf := config.GetConfig()

	utils.InitValidator()
	logger.SetupLogger(cnf.ServiceName)

	db, err := repo.NewDB(cnf.DB)
	if err != nil {
		slog.Error("Failed to Connect with Database:", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}
	defer repo.CloseDB(db)

	err = repo.MigrateDB(db.Db, cnf.MigrationSource)
	if err != nil {
		slog.Error("Failed to Migrate Database:", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}

	vehiclesRepo := repo.NewVehiclesRepo(db)
	routesRepo := repo.NewRoutesRepo(db)
	picRepo := repo.NewPicsRepo(db)
	reviewRepo := repo.NewReviewsRepo(db)

	vehicleSvc := vehicles.NewService(cnf, vehiclesRepo, routesRepo, picRepo, reviewRepo)

	handlers := handlers.NewHandlers(cnf, vehicleSvc)

	server := web.NewServer(cnf, handlers)
	server.Run()
	//grpc := grpc.NewGRPC(cnf, vehicleSvc)
	// grpc.Start()
	server.Wg.Wait()

	return nil
}
