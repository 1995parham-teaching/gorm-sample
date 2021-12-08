package one2many

import "time"

type Owner struct {
	Id        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	Books     []Book
}

type Book struct {
	Id          uint `gorm:"primarykey"`
	Name        string
	PublishDate time.Time
	OwnerID     uint `sql:"index"`
}
