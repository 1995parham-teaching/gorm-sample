package gorm

// TODO Add database options e.g. ReadTimeout, WriteTimeout, ... .
type Config struct {
	Username string `default:"admin"`
	Password string `default:"admin"`
	Host     string `default:"127.0.0.1"`
	Port     int    `default:"3306"`
	Database string `default:"bell" split_words:"true"`
}
