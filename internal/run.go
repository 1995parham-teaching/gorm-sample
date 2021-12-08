package internal

import (
	"test-gorm/internal/models"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.Statement.Migrator().DropTable(&models.Owner{})
	db.Migrator().CreateTable(&models.Owner{})

	sample := models.Owner{
		FirstName: "Mohammad",
		LastName:  "Nasr",
	}

	db.Create(&sample)
}
