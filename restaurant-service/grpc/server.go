package grpc

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	pb "restaurant-service/grpc/restaurants"

	gRPC "google.golang.org/grpc"
)

func (server *grpc) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server.cnf.GrpcPort))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	userService := NewRestaurantService(server.restSvc)

	gRPCServer := gRPC.NewServer()

	pb.RegisterRestaurantsServiceServer(gRPCServer, userService)

	server.Wg.Add(1)

	go func() {
		slog.Info(fmt.Sprintf("gRPC Server Listening at %v", server.cnf.GrpcPort))

		defer server.Wg.Done()

		if err := gRPCServer.Serve(lis); err != nil {
			slog.Error(err.Error())
		}
	}()
}
