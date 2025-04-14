package cmd

import (
	"restaurant-service/config"
	"restaurant-service/grpc"
	"restaurant-service/restaurant"
	"restaurant-service/web"
	"restaurant-service/web/handlers"
)

func Cmd() {
	conf := config.GetConfig()
	locSvc := restaurant.NewService(conf)
	handlers := handlers.NewHandlers(conf, locSvc)
	grpc := grpc.NewGRPC(conf, locSvc)

	server := web.NewServer(conf, handlers)
	server.Run()
	grpc.Start()
	server.Wg.Wait()
}
