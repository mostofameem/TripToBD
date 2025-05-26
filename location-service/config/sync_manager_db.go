package config

type PgDB struct {
	DbHost                 string `mapstructure:"SYNC_MANAGER_DB_HOST"                        validate:"required"`
	DbPort                 int    `mapstructure:"SYNC_MANAGER_DB_PORT"                        validate:"required"`
	DbName                 string `mapstructure:"SYNC_MANAGER_DB_NAME"                        validate:"required"`
	DbUser                 string `mapstructure:"SYNC_MANAGER_DB_USER"                        validate:"required"`
	DbPassword             string `mapstructure:"SYNC_MANAGER_DB_PASS"                        validate:"required"`
	DbMaxIdleTimeInMinutes int    `mapstructure:"SYNC_MANAGER_DB_MAX_IDLE_TIME_IN_MINUTES"    validate:"required"`
	DbMaxOpenConns         int    `mapstructure:"SYNC_MANAGER_DB_MAX_OPEN_CONNS"             validate:"required"`
	DbMaxIdleConns         int    `mapstructure:"SYNC_MANAGER_DB_MAX_IDLE_CONNS"             validate:"required"`
	DbEnableSSLMode        bool   `mapstructure:"SYNC_MANAGER_DB_ENABLE_SSL_MODE"`
}

func (db *PgDB) User() string              { return db.DbUser }
func (db *PgDB) Password() string          { return db.DbPassword }
func (db *PgDB) Host() string              { return db.DbHost }
func (db *PgDB) Port() int                 { return db.DbPort }
func (db *PgDB) Name() string              { return db.DbName }
func (db *PgDB) EnableSSLMode() bool       { return db.DbEnableSSLMode }
func (db *PgDB) MaxIdleTimeInMinutes() int { return db.DbMaxIdleTimeInMinutes }
func (db *PgDB) MaxOpenConns() int         { return db.DbMaxOpenConns }
func (db *PgDB) MaxIdleConns() int         { return db.DbMaxIdleConns }
