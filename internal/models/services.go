package models

type Service struct {
	BaseModel
	Name   string
	UserID int
	User   User
}
