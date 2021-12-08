package simple

import (
	"github.com/cng-by-example/gorm-sample/internal/models"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Command(db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use: "simple",
		Run: func(_ *cobra.Command, _ []string) { main(db) },
	}
}

func main(db *gorm.DB) {
	db.Migrator().DropTable(&models.Owner{})
	db.Migrator().CreateTable(&models.Owner{})

	sample := models.Owner{
		FirstName: "Mohammad",
		LastName:  "Nasr",
	}

	db.Create(&sample)

	sample.FirstName = "Parham"
	sample.LastName = "Alvani"

	db.Debug().Save(&sample)
}
