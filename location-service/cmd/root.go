package cmd

import (
	"post-service/config"
	"post-service/grpc"
	"post-service/location"
	"post-service/route"
	"post-service/web"
	"post-service/web/handlers"
)

func Main() {
	conf := config.GetConfig()
	locSvc := location.NewService(conf)
	routeSvc := route.NewRouteService()
	handlers := handlers.NewHandlers(conf, locSvc, routeSvc)
	grpc := grpc.NewGRPC(conf)

	server := web.NewServer(conf, handlers)
	server.Run()
	grpc.Start()
	server.Wg.Wait()
}
