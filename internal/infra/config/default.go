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
			DSN:             "host=127.0.0.1 user=postgres password=postgres DB.name=pgsql port=5432 sslmode=disable",
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