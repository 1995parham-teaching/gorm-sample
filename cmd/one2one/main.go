package one2one

import (
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Command(db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use: "one2one",
		Run: func(_ *cobra.Command, _ []string) { main(db) },
	}
}

func main(db *gorm.DB) {
	// internal.Run(db)
}
