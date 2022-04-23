package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID string `json:"id" gorm:"primary_key"`
	// ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedOn time.Time      `json:"created_on"`
	UpdatedOn time.Time      `json:"updated_on"`
	DeletedOn gorm.DeletedAt `json:"-" sql:"index"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.NewV4().String()
	return
}
