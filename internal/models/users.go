package models

type User struct {
	BaseModel
	Email    string    `json:"-"`
	Name     string    `json:"-"`
	Services []Service `json:"services"`
}
