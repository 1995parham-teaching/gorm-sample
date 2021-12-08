package main

import (
	"fmt"

	"github.com/cng-by-example/gorm-sample/cmd/many2many"
	"github.com/cng-by-example/gorm-sample/cmd/one2many"
	"github.com/cng-by-example/gorm-sample/cmd/one2one"
	"github.com/cng-by-example/gorm-sample/cmd/simple"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	short = "short description"
	long  = `long description`
)

// TODO Add database options e.g. ReadTimeout, WriteTimeout, ... .
type Config struct {
	Username string `default:"admin"`
	Password string `default:"admin"`
	Host     string `default:"127.0.0.1"`
	Port     int    `default:"3306"`
	Database string `default:"bell" split_words:"true"`
}

func NewMysql(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	dbConfig := Config{
		Username: "root",
		Password: "pass",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "gorm",
	}

	db, err := NewMysql(&dbConfig)
	if err != nil {
		panic(err)
	}

	cmd := &cobra.Command{Short: short, Long: long}
	cmd.AddCommand(simple.Command(db), one2one.Command(), one2many.Command(), many2many.Command())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		panic(map[string]interface{}{"err": err, "msg": errExecuteCMD})
	}
}
