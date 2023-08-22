package main

import (
	"database/sql"
	"log"
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
		fx.Provide(db.Provide),
		fx.Provide(logger.Provide),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(start),
	).Run()
}

func start(shutdowner fx.Shutdowner, db *gorm.DB, logger *zap.Logger) {
	// I really love the way Go has for describing the date/time formats.
	// you write down 2 Jan 2006 15:04:00 -0700 in your desired format
	// and it will figure your format out.
	birthday, err := time.Parse("2 January 2006 at 15:04 -0700", "12 October 1999 at 19:20 +0330")
	if err != nil {
		log.Fatal(err)
	}

	// create user with gorm.
	// please pay attention to time. you must provide the valid field when you are using
	// NullTime.
	// nolint: exhaustruct
	db.Create(&model.User{
		Model:    gorm.Model{},
		ID:       1,
		Name:     "Elahe Dastan",
		Email:    "elahe.dstn@gmail.com",
		Birthday: sql.NullTime{Time: birthday, Valid: true},
	})

	var user model.User

	db.First(&user, 1)
	logger.Info("first user from database", zap.Any("user", user))

	var users []model.User

	db.Find(&users)
	logger.Info("users from database", zap.Any("users", users))

	_ = shutdowner.Shutdown()
}
