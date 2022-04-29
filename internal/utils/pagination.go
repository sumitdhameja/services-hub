package utils

import (
	"github.com/sumitdhameja/services-hub/internal/dto"
	"gorm.io/gorm"
)

var (
	Limit = "Limit"
	Page  = "Page"
)

func PaginateScope(p *dto.Pageable) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		searchStr, values := p.BuildSearchString()
		db.Where(searchStr, values...)
		return db.Offset(p.GetOffset()).Limit(p.GetLimit())

		// return db.Where(searchByteBuffer.String(), sValues...).Offset(p.GetOffset()).Limit(p.GetLimit())
	}
}
