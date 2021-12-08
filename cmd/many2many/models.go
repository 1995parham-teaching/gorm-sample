package many2many

import (
	"time"

	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	FirstName string
	LastName  string
	Books     []Book
}

type Book struct {
	gorm.Model
	Name        string
	PublishDate time.Time
	OwnerID     uint     `sql:"index"`
	Authors     []Author `gorm:"many2many:books_athors"`
}

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
}
