package config

import "sync"

var cntOnce = sync.Once{}

type Mode string

const DebugMode = Mode("debug")
const ReleaseMode = Mode("release")

type GrpcUrlsConfig struct {
	UserUrl       string `mapstructure:"GRPC_USER_URL"       `
	RestaurantUrl string `mapstructure:"GRPC_RESTAURANT_URL" `
}

type Config struct {
	Mode        Mode   `mapstructure:"MODE"                             validate:"required"`
	ServiceName string `mapstructure:"SERVICE_NAME"                     validate:"required"`
	HttpPort    int    `mapstructure:"HTTP_PORT"                        validate:"required"`
	GrpcPort    int    `mapstructure:"GRPC_PORT"                        validate:"required"`
	JwtSecret   string `mapstructure:"JWT_SECRET"                      validate:"required"`
	//PgDB                   DBConfig
	MongoDB MongoDBConfig
	//MigrationSource        string `mapstructure:"MIGRATION_SOURCE"                 validate:"required"`
	GrpcUrls               GrpcUrlsConfig
	GrpcReqTimeOutInSecond int `mapstructure:"GRPC_REQ_TIMEOUT_IN_SECOND"`
}

var config *Config

func GetConfig() *Config {
	cntOnce.Do(func() {
		LoadConfig()
	})
	return config
}
