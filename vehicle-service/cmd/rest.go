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
		slog.Error("Failed to connect mongo database:", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}
	defer db.Db.Close()

	vehicleRepo := repo.NewVehiclesRepo(db)
	routeRepo := repo.NewRoutesRepo(db)
	picRepo := repo.NewPicsRepo(db)
	reviewRepo := repo.NewReviewsRepo(db)

	vehicleSvc := vehicles.NewService(cnf, vehicleRepo, routeRepo, picRepo, reviewRepo)

	handlers := handlers.NewHandlers(cnf, vehicleSvc)

	server := web.NewServer(cnf, handlers)
	server.Run()
	// //grpc := grpc.NewGRPC(cnf, vehicleSvc)
	// // grpc.Start()
	server.Wg.Wait()

	return nil
}
