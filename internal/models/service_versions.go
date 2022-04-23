package models

type ServiceVersion struct {
	BaseModel
	ServiceID      string `json:"service_id" gorm:"size:191"`
	Version        string `json:"version"`
	URL            string `json:"url"`
	OtherCoolStuff string `json:"other_cool_stuff"`
}
