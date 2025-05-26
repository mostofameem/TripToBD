package cmd

import (
	"location-service/config"
	"location-service/location"
	"location-service/logger"
	"location-service/repo"
	"location-service/web"
	"location-service/web/handlers"
	"location-service/web/utils"
	"log/slog"

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

	mongoDB, err := repo.ConnectMongoDB(cnf.MongoDB)
	if err != nil {
		slog.Error("Failed to connect mongo database:", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}
	defer mongoDB.Close()

	db := mongoDB.Client.Database(cnf.MongoDB.GetDatabase())

	locationRepo := repo.NewLocationRepo(db)
	routeRepo := repo.NewRouteRepo(db)

	locSvc := location.NewService(cnf, locationRepo, routeRepo)

	handlers := handlers.NewHandlers(cnf, locSvc)

	server := web.NewServer(cnf, handlers)
	server.Run()
	//grpc := grpc.NewGRPC(cnf, vehicleSvc)
	// grpc.Start()
	server.Wg.Wait()

	return nil
}
