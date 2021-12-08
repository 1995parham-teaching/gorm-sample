package simple

import (
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
	// dbConfig := gorm.Config{
	// 	Username: "root",
	// 	Password: "hp590mt",
	// 	Host:     "127.0.0.1",
	// 	Port:     3306,
	// 	Database: "gorm",
	// }

	// db, err := gorm.NewMysql(&dbConfig)
	// if err != nil {
	// 	panic(err)
	// }
	// internal.Run(db)
}
