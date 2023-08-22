package config

import (
	"time"

	"github.com/1995parham-teaching/gorm-sample/internal/infra/db"
	"github.com/1995parham-teaching/gorm-sample/internal/infra/logger"
	"go.uber.org/fx"
)

// Default return default configuration.
// nolint: gomnd
func Default() Config {
	return Config{
		Out: fx.Out{},
		Database: db.Config{
			DSN:             "postgres://postgres:postgres@127.0.0.1:5432/pgsql?search_path=public&sslmode=disable",
			ConnMaxLifeTime: 10 * time.Second,
			ConnMaxIdleTime: 5 * time.Second,
			MaxOpenConns:    10,
			MaxIdleConns:    5,
		},
		Logger: logger.Config{
			Level: "debug",
		},
	}
}
