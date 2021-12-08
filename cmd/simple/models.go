package simple

type Owner struct {
	Id        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
}

func (o *Owner) TableName() string {
	return "owners"
}
