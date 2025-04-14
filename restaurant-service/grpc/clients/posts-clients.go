package clients

import (
	"fmt"
	"log/slog"
	"restaurant-service/config"
	"restaurant-service/grpc/posts"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PostsClients struct {
	conn   *grpc.ClientConn
	client posts.PostServiceClient
}

var postConfig *PostsClients
var cntOnce = sync.Once{}

func NewPostsClient(conf *config.Config) *PostsClients {
	slog.Info(
		fmt.Sprintf("Initialize user clients. url: %s", conf.GrpcUrls.PostUrl),
	)
	conn, err := grpc.NewClient(
		conf.GrpcUrls.PostUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		slog.Error(err.Error())
		return &PostsClients{}
	}

	postConfig = &PostsClients{
		conn:   conn,
		client: posts.NewPostServiceClient(conn),
	}
	return postConfig
}
func GetPostsClient() *PostsClients {
	cntOnce.Do(func() {
		NewPostsClient(config.GetConfig())
	})
	return postConfig
}
