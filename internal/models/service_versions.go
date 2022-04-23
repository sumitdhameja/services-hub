package models

type ServiceVersion struct {
	BaseModel
	Name      string
	ServiceID int
	Service   Service
}
