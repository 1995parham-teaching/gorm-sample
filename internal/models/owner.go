package models

type Owner struct {
	FirstName string
	LastName  string
}

func (o *Owner) TableName() string {
	return "owners"
}
