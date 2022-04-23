package models

import (
	"github.com/sumitdhameja/services-hub/internal/logger"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	logger.Info("Migrating model")
	// TODO: auto migrate models
}
