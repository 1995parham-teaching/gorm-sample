package simple

import (
	"fmt"

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
	db.Migrator().DropTable(&Owner{})
	db.Migrator().CreateTable(&Owner{})

	sample := Owner{
		FirstName: "Mohammad",
		LastName:  "Nasr",
	}

	db.Create(&sample)

	sample.LastName = "Nasr Esfahani"

	db.Debug().Save(&sample)

	var owner Owner
	db.Debug().First(&owner)
	fmt.Println(owner)

	if owner.FirstName != "" {
		db.Debug().Delete(&owner)
	}

	db.Debug().First(&owner)
	fmt.Println(owner)
}
