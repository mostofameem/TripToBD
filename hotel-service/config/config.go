package config

import "sync"

var cntOnce = sync.Once{}

type Mode string

const DebugMode = Mode("debug")
const ReleaseMode = Mode("release")

type DBConfig struct {
	Host                 string `mapstructure:"DB_HOST"                     validate:"required"`
	Port                 int    `mapstructure:"DB_PORT"                     validate:"required"`
	Name                 string `mapstructure:"DB_NAME"                     validate:"required"`
	User                 string `mapstructure:"DB_USER"                     validate:"required"`
	Password             string `mapstructure:"DB_PASS"                     validate:"required"`
	MaxIdleTimeInMinutes int    `mapstructure:"MAX_IDLE_TIME_IN_MINUTES"    validate:"required"`
	EnableSSLMode        bool   `mapstructure:"ENABLE_SSL_MODE"`
}
type GrpcUrlsConfig struct {
	UserUrl     string `mapstructure:"GRPC_USER_URL"     validate:"required"`
	LocationUrl string `mapstructure:"GRPC_LOCATION_URL" validate:"required"`
}

type Config struct {
	Mode                   Mode   `mapstructure:"MODE"                             validate:"required"`
	ServiceName            string `mapstructure:"SERVICE_NAME"                     validate:"required"`
	HttpPort               int    `mapstructure:"HTTP_PORT"                        validate:"required"`
	GrpcPort               int    `mapstructure:"GRPC_PORT"                        validate:"required"`
	JwtSecret              string `mapstructure:"JWT_SECREAT"                      validate:"required"`
	DB                     *DBConfig
	MigrationSource        string `mapstructure:"MIGRATION_SOURCE"                 validate:"required"`
	GrpcUrls               *GrpcUrlsConfig
	GrpcReqTimeOutInSecond int `mapstructure:"GRPC_REQ_TIMEOUT_IN_SECOND"`
}

var config *Config

func GetConfig() *Config {
	cntOnce.Do(func() {
		LoadConfig()
	})
	return config
}
