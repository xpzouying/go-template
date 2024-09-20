package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type DBConfig struct {
	dbType DBType
	dbDsn  string
}

type DBConfigOption func(config *DBConfig)

func WithDBType(dbType DBType) DBConfigOption {
	return func(config *DBConfig) {
		config.dbType = dbType
	}
}

func WithDBDsn(dbDsn string) DBConfigOption {
	return func(config *DBConfig) {
		config.dbDsn = dbDsn
	}
}

func InitDB(dbType DBType, dbDsn string) error {
	var err error
	db, err = NewDB(WithDBType(dbType), WithDBDsn(dbDsn))
	if err != nil {
		return errors.Wrap(err, "failed to initialize database")
	}

	// --- Auto Migrate ---
	for _, model := range []interface{}{
		&User{},
	} {
		if err := db.AutoMigrate(model); err != nil {
			return errors.Wrap(err, "failed to auto migrate")
		}
	}

	return nil
}
