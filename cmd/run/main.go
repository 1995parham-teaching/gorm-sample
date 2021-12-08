package run

import (
	"test-gorm/internal"
	"test-gorm/internal/gorm"

	"github.com/spf13/cobra"
)

const (
	use   = "run"
	short = "run sample application"
)

func Command() *cobra.Command {
	return &cobra.Command{Use: use, Short: short, Run: main}
}

func main(cmd *cobra.Command, _ []string) {
	dbConfig := gorm.Config{
		Username: "root",
		Password: "hp590mt",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "gorm",
	}

	db, err := gorm.NewMysql(&dbConfig)
	if err != nil {
		panic(err)
	}
	internal.Run(db)
}
