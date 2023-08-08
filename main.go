package main

import (
	"database/sql"
	"log"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null; index"`
	Email    string `gorm:"not null; check:,email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'"`
	Birthday sql.NullTime
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger = zap.NewNop()
	}

	dsn := "host=127.0.0.1 user=postgres password=postgres DB.name=pgsql port=5432 sslmode=disable"

	// nolint: exhaustruct
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: zapgorm2.New(logger),
	})
	if err != nil {
		log.Fatal(err)
	}

	// migrate the schema with gorm migrator manually
	if err := db.Migrator().DropTable(new(User)); err != nil {
		log.Fatal(err)
	}

	if err := db.Migrator().CreateTable(new(User)); err != nil {
		log.Fatal(err)
	}

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
	db.Create(&User{
		Model:    gorm.Model{},
		ID:       1,
		Name:     "Elahe Dastan",
		Email:    "elahe.dstn@gmail.com",
		Birthday: sql.NullTime{Time: birthday, Valid: true},
	})

	var user User

	db.First(&user, 1)
	logger.Info("first user from database", zap.Any("user", user))

	var users []User

	db.Find(&users)
	logger.Info("users from database", zap.Any("users", users))
}
