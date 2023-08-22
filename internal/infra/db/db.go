package db

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type Config struct {
	DSN string `json:"dsn,omitempty" koanf:"dsn"`

	MaxIdleConns    int           `json:"max_idle_conns,omitempty"     koanf:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns,omitempty"     koanf:"max_open_conns"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time,omitempty" koanf:"conn_max_idle_time"`
	ConnMaxLifeTime time.Duration `json:"conn_max_life_time,omitempty" koanf:"conn_max_life_time"`
}

func Provide(cfg Config, logger *zap.Logger) (*gorm.DB, error) {
	// nolint: exhaustruct
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
		Logger: zapgorm2.New(logger),
	})
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	db.Debug()

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("acquiring sql database failed: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifeTime)

	return db, nil
}
