package clients

import (
	"context"
	"fmt"
	"log/slog"
	"post-service/config"
	"post-service/grpc/users"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClients struct {
	cnf    *config.Config
	conn   *grpc.ClientConn
	client users.UserServiceClient
}

var userClients *UserClients

var userClientCntOnce = sync.Once{}

func NewUserClient(conf *config.Config) *UserClients {
	userClientCntOnce.Do(func() {
		slog.Info(
			fmt.Sprintf("Initialize user clients. url: %s", conf.GrpcUrls.UserUrl),
		)
		conn, err := grpc.NewClient(
			conf.GrpcUrls.UserUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		userClients = &UserClients{
			cnf:    conf,
			conn:   conn,
			client: users.NewUserServiceClient(conn),
		}
	})
	return userClients
}
func (c *UserClients) GetUserName(ctx context.Context, req *users.GetUserNameReq) (*users.GetUserNameRes, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(c.cnf.GrpcReqTimeOutInSecond)*time.Second)
	defer cancel()

	slog.Info("asking username ")

	res, err := c.client.GetUserName(ctx, req)
	if err != nil {
		slog.Error("grpc res failed")
		return nil, err
	}

	return res, nil
}
