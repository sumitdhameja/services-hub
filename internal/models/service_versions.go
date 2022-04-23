package models

type ServiceVersion struct {
	BaseModel
	Version        string
	URL            string
	OtherCoolStuff string
	ServiceID      string `gorm:"size:191"`
	Service        Service
}
