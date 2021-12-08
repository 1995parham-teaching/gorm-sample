package one2many

import (
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Command(db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use: "one2many",
		Run: func(_ *cobra.Command, _ []string) { main(db) },
	}
}

func main(db *gorm.DB) {
	db.Migrator().DropTable(&Owner{}, &Book{})
	db.Migrator().CreateTable(&Owner{}, &Book{})
}
