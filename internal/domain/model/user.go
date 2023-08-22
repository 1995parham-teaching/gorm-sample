package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null; index"`
	Email    string `gorm:"not null; check:,email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'"`
	Birthday sql.NullTime
}
