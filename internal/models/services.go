package models

type Service struct {
	BaseModel
	Title           string           `json:"title"`
	Description     string           `json:"description"`
	UserID          string           `gorm:"size:191"`
	ServiceVersions []ServiceVersion `gorm:"foreignKey:ServiceID"`
}
