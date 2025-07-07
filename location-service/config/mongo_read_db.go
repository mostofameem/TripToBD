package config

type MongoDB struct {
	MONGO_URI           string `mapstructure:"MONGO_URI"                       validate:"required"`
	Database            string `mapstructure:"MONGO_DATABASE"                  validate:"required"`
	User                string `mapstructure:"MONGO_USER"`
	Password            string `mapstructure:"MONGO_PASSWORD"`
	ReplicaSet          string `mapstructure:"MONGO_REPLICA_SET"`
	AuthSource          string `mapstructure:"MONGO_AUTH_SOURCE"`
	SSL                 bool   `mapstructure:"MONGO_SSL"`
	MaxPoolSize         uint64 `mapstructure:"MONGO_MAX_POOL_SIZE"`
	MinPoolSize         uint64 `mapstructure:"MONGO_MIN_POOL_SIZE"`
	MaxConnIdleTimeInMs int    `mapstructure:"MONGO_MAX_CONN_IDLE_TIME_MS"`
}

func (db *MongoDB) GetURI() string              { return db.MONGO_URI }
func (db *MongoDB) GetDatabase() string         { return db.Database }
func (db *MongoDB) GetUser() string             { return db.User }
func (db *MongoDB) GetPassword() string         { return db.Password }
func (db *MongoDB) GetReplicaSet() string       { return db.ReplicaSet }
func (db *MongoDB) GetAuthSource() string       { return db.AuthSource }
func (db *MongoDB) GetSSL() bool                { return db.SSL }
func (db *MongoDB) GetMaxPoolSize() uint64      { return db.MaxPoolSize }
func (db *MongoDB) GetMinPoolSize() uint64      { return db.MinPoolSize }
func (db *MongoDB) GetMaxConnIdleTimeInMs() int { return db.MaxConnIdleTimeInMs }
