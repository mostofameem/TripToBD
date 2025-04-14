package grpc

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	pb "post-service/grpc/posts"

	gRPC "google.golang.org/grpc"
)

func (server *grpc) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server.cnf.GrpcPort))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	postSvc := NewPostsService()

	gRPCServer := gRPC.NewServer()

	pb.RegisterPostServiceServer(gRPCServer, postSvc)

	server.Wg.Add(1)

	go func() {
		log.Println(fmt.Sprintf("gRPC Server Listening at %v", server.cnf.GrpcPort))

		defer server.Wg.Done()

		if err := gRPCServer.Serve(lis); err != nil {
			slog.Error(err.Error())
		}
	}()
}
