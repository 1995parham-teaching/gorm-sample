package internal

import (
	"test-gorm/internal/models"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	// db.AutoMigrate(&User{}, &CreditCard{})
	db.Statement.Migrator().DropTable(&models.Owner{})
	db.Migrator().CreateTable(&models.Owner{})
}
