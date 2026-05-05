package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/1995parham-teaching/gorm-sample/internal/domain/model"
	"github.com/1995parham-teaching/gorm-sample/internal/infra/config"
	"github.com/1995parham-teaching/gorm-sample/internal/infra/db"
	"github.com/1995parham-teaching/gorm-sample/internal/infra/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	fx.New(
		fx.Provide(config.Provide),
		fx.Provide(logger.Provide),
		fx.Provide(db.Provide),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(start),
	).Run()
}

func start(shutdowner fx.Shutdowner, gdb *gorm.DB, logger *zap.Logger) {
	ctx := context.Background()

	// I really love the way Go has for describing the date/time formats.
	// you write down 2 Jan 2006 15:04:00 -0700 in your desired format
	// and it will figure your format out.
	birthday, err := time.Parse("2 January 2006 at 15:04 -0700", "12 October 1999 at 19:20 +0330")
	if err != nil {
		logger.Fatal("cannot parse datetime into given format", zap.Error(err))
	}

	// create user with gorm generic API.
	// please pay attention to time. you must provide the valid field when you are using
	// NullTime.
	// nolint: exhaustruct
	if err := gorm.G[model.User](gdb).Create(ctx, &model.User{
		Model:    gorm.Model{},
		ID:       1,
		Name:     "Elahe Dastan",
		Email:    "elahe.dstn@gmail.com",
		Birthday: sql.NullTime{Time: birthday, Valid: true},
	}); err != nil {
		logger.Fatal("cannot create user", zap.Error(err))
	}

	user, err := gorm.G[model.User](gdb).Where("id = ?", 1).First(ctx)
	if err != nil {
		logger.Fatal("cannot fetch first user", zap.Error(err))
	}

	logger.Info("first user from database", zap.Any("user", user))

	users, err := gorm.G[model.User](gdb).Find(ctx)
	if err != nil {
		logger.Fatal("cannot fetch users", zap.Error(err))
	}

	logger.Info("users from database", zap.Any("users", users))

	_ = shutdowner.Shutdown()
}
