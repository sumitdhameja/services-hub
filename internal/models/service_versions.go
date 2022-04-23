package models

type ServiceVersion struct {
	BaseModel
	ServiceID      string `gorm:"size:191"`
	Version        string
	URL            string
	OtherCoolStuff string
}
