package config

type MongoDBConfig interface {
	GetURI() string
	GetDatabase() string
	GetUser() string
	GetPassword() string
	GetReplicaSet() string
	GetAuthSource() string
	GetSSL() bool
	GetMaxPoolSize() uint64
	GetMinPoolSize() uint64
	GetMaxConnIdleTimeInMs() int
}
