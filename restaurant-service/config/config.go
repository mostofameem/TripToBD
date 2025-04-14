package config

import "sync"

var cntOnce = sync.Once{}

type Mode string

type DBConfig struct {
	Host                string `json:"host" validate:"required"`
	Port                int    `json:"port" validate:"required"`
	Name                string `json:"name" validate:"required"`
	User                string `json:"user" validate:"required"`
	Pass                string `json:"pass" validate:"required"`
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" validate:"required"`
	EnableSSLMode       bool   `json:"enable_ssl_mode"`
}
type GrpcUrlsConfig struct {
	UserUrl string `json:"user"`
	PostUrl string `json:"posts" validate:"required"`
}

type MongoDBConfig struct {
	Host                string `json:"host" validate:"required"`
	Port                int    `json:"port" validate:"required"`
	Name                string `json:"name" validate:"required"`
	User                string `json:"user" validate:"required"`
	Pass                string `json:"pass" validate:"required"`
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" validate:"required"`
	EnableSSLMode       bool   `json:"enable_ssl_mode"`
}

const DebugMode = Mode("debug")
const ReleaseMode = Mode("release")

type Config struct {
	Mode                   Mode           `json:"mode"  validate:"required"`
	ServiceName            string         `json:"service_name" validate:"required"`
	HttpPort               int            `json:"http_port"  validate:"required"`
	GrpcPort               int            `json:"grpc_port" validate:"required"`
	RabbitmqURL            string         `json:"rabbitmq_url" validate:"required"`
	DB                     DBConfig       `json:"db" validate:"required"`
	MongoDB                MongoDBConfig  `json:"mongodb" validate:"required"`
	GrpcUrls               GrpcUrlsConfig `json:"grpc_urls"`
	MigrationSource        string         `json:"migrations" validate:"required"`
	GrpcReqTimeOutInSecond int            `json:"grpc_req_timeout_in_second" validate:"required"`
	RmqReconnectDelay      int            `json:"rmq_reconnect_delay"              validate:"required"`
	RmqRetryInterval       int            `json:"rmq_retry_interval"               validate:"required"`
}

var config *Config

func GetConfig() *Config {
	cntOnce.Do(func() {
		LoadConfig()
	})
	return config
}
