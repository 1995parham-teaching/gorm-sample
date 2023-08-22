package db

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type Config struct {
	DSN string

	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifeTime time.Duration
}

func Provide(cfg Config, logger *zap.Logger) *gorm.DB {
	// nolint: exhaustruct
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
		Logger: zapgorm2.New(logger),
	})
	if err != nil {
		logger.Fatal("database connection failed", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("acquiring sql database failed", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifeTime)

	return db
}
