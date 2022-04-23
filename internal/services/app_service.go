package services

import (
	"errors"

	"github.com/sumitdhameja/services-hub/internal/dto"
	"github.com/sumitdhameja/services-hub/internal/models"
	"gorm.io/gorm"
)

// appService is a service private
type appService struct {
	db *gorm.DB
}

//AppService interface
type AppService interface {
	GetAllService(userID string, page dto.Pageable) (*[]models.Service, error)
	GetService(userID string, id string) (*models.Service, error)
}

// NewAppService create appService
func NewAppService(db *gorm.DB) AppService {
	return &appService{db}
}

// GetAllService return all AppService
func (p appService) GetAllService(userID string, pageable dto.Pageable) (*[]models.Service, error) {
	user := new(models.User)

	if err := p.db.Preload("Services.ServiceVersions").First(&user, &models.User{BaseModel: models.BaseModel{ID: userID}}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no user found")
		}
		return nil, errors.New("unknown error")
	}

	return &user.Services, nil
}

// Getservice return only one appservice
func (p appService) GetService(userID string, id string) (*models.Service, error) {
	service := &models.Service{}

	if err := p.db.Preload("ServiceVersions").Where(&models.Service{UserID: userID}).First(&service, &models.Service{BaseModel: models.BaseModel{ID: id}}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no user found")
		}
		return nil, errors.New("unknown error")
	}

	return service, nil
}
