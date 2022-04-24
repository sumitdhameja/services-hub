package models

import (
	"fmt"

	"github.com/sumitdhameja/services-hub/internal/logger"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	logger.Info("Migrating model")
	if err := db.AutoMigrate(&User{}, &Service{}, &ServiceVersion{}); err != nil {
		logger.Error(fmt.Sprintf("Can't automigrate schema %s", err))
	}
}
