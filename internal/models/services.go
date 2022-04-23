package models

type Service struct {
	BaseModel
	Title           string           `json:"title"`
	Description     string           `json:"description"`
	UserID          string           `json:"user_id" gorm:"size:191"`
	ServiceVersions []ServiceVersion `json:"service_versions" gorm:"foreignKey:ServiceID"`
}
