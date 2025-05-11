package model

import (
	"github.com/glebarez/sqlite"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBType string

const (
	DBTypeMySQL      DBType = "mysql"
	DBTypePostgreSQL DBType = "postgresql"
	DBTypeSQLite     DBType = "sqlite3"
)

func newDefaultDBConfig() *DBConfig {
	return &DBConfig{
		dbType: DBTypeSQLite,
		dbDsn:  "file::memory:?cache=shared",
	}
}

func NewDB(opts ...DBConfigOption) (db *gorm.DB, err error) {
	cfg := newDefaultDBConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	switch cfg.dbType {
	case DBTypeMySQL:
		db, err = newMySQLDB(cfg.dbDsn)
	case DBTypeSQLite:
		db, err = newSqlite3DB(cfg.dbDsn)
	case DBTypePostgreSQL:
		db, err = newPostgresDB(cfg.dbDsn)
	default:
		return nil, errors.Errorf("unknown db type: %s", cfg.dbType)
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}
	return db, nil
}

func NewMemorySqlite3() (*gorm.DB, error) {
	return NewDB(WithDBType(DBTypeSQLite), WithDBDsn("file::memory:?cache=shared"))
}

func newSqlite3DB(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
}

func newPostgresDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
}

func newMySQLDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
}
