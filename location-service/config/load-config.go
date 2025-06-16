package config

import (
	"log/slog"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() error {

	exit := func(err error) {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		slog.Warn(".env not found, that's okay!")
	}

	viper.AutomaticEnv()

	config = &Config{
		Mode:            Mode(viper.GetString("MODE")),
		ServiceName:     viper.GetString("SERVICE_NAME"),
		HttpPort:        viper.GetInt("HTTP_PORT"),
		GrpcPort:        viper.GetInt("GRPC_PORT"),
		JwtSecret:       viper.GetString("JWT_SECRET"),
		//MigrationSource: viper.GetString("MIGRATION_SOURCE"),
		// PgDB: &PgDB{
		// 	DbHost:                 viper.GetString("PG_DB_HOST"),
		// 	DbPort:                 viper.GetInt("PG_DB_PORT"),
		// 	DbName:                 viper.GetString("PG_DB_NAME"),
		// 	DbUser:                 viper.GetString("PG_DB_USER"),
		// 	DbPassword:             viper.GetString("PG_DB_PASS"),
		// 	DbMaxIdleTimeInMinutes: viper.GetInt("PG_DB_MAX_IDLE_TIME_IN_MINUTES"),
		// 	DbMaxOpenConns:         viper.GetInt("PG_DB_MAX_OPEN_CONNS"),
		// 	DbMaxIdleConns:         viper.GetInt("PG_DB_MAX_IDLE_CONNS"),
		// 	DbEnableSSLMode:        viper.GetBool("PG_DB_ENABLE_SSL_MODE"),
		// },
		MongoDB: &MongoDB{
			MONGO_URI:           viper.GetString("MONGO_URI"),
			Database:            viper.GetString("MONGO_DATABASE"),
			User:                viper.GetString("MONGO_USER"),
			Password:            viper.GetString("MONGO_PASSWORD"),
			ReplicaSet:          viper.GetString("MONGO_REPLICA_SET"),
			AuthSource:          viper.GetString("MONGO_AUTH_SOURCE"),
			SSL:                 viper.GetBool("MONGO_SSL"),
			MaxPoolSize:         viper.GetUint64("MONGO_MAX_POOL_SIZE"),
			MinPoolSize:         viper.GetUint64("MONGO_MIN_POOL_SIZE"),
			MaxConnIdleTimeInMs: viper.GetInt("MONGO_MAX_CONN_IDLE_TIME_MS"),
		},
		GrpcReqTimeOutInSecond: viper.GetInt("GRPC_REQ_TIMEOUT_IN_SECOND"),
	}
	v := validator.New()
	if err = v.Struct(config); err != nil {
		exit(err)
	}

	return nil
}
