package cmd

import (
	"hotel-service/config"
	hotel "hotel-service/hotel"
	"hotel-service/logger"
	"hotel-service/repo"
	"hotel-service/web"
	"hotel-service/web/handlers"
	"hotel-service/web/utils"
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

	hotelRepo := repo.NewHotelRepo(db)

	picRepo := repo.NewPicsRepo(db)
	reviewRepo := repo.NewReviewsRepo(db)

	hotelvc := hotel.NewService(cnf, hotelRepo, picRepo, reviewRepo)

	handlers := handlers.NewHandlers(cnf, hotelvc)

	server := web.NewServer(cnf, handlers)
	server.Run()
	//grpc := grpc.NewGRPC(cnf, hotelvc)
	// grpc.Start()
	server.Wg.Wait()

	return nil
}
