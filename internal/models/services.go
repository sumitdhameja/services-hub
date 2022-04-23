package models

type Service struct {
	BaseModel
	Name        string
	Title       string
	Description string
	UserID      string `gorm:"size:191"`
	User        User
}
