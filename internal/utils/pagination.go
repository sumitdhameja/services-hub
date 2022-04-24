package utils

import (
	"bytes"
	"fmt"

	"github.com/sumitdhameja/services-hub/internal/dto"
	"gorm.io/gorm"
)

var (
	Limit = "Limit"
	Page  = "Page"
)

func PaginateScope(p *dto.Pageable) func(db *gorm.DB) *gorm.DB {

	var searchByteBuffer bytes.Buffer
	sOptions := p.GetSearchOptions()
	sValues := []interface{}{}

	for i, s := range sOptions {
		searchByteBuffer.WriteString(fmt.Sprintf(" %s %s ? ", s.Column, s.Operator))
		sValues = append(sValues, fmt.Sprintf("%%%v%%", s.SearchString))
		if i < len(sOptions)-1 {
			searchByteBuffer.WriteString("OR")
		}
	}

	return func(db *gorm.DB) *gorm.DB {
		db.Where(searchByteBuffer.String(), sValues...)
		return db.Offset(p.GetOffset()).Limit(p.GetLimit())

		// return db.Where(searchByteBuffer.String(), sValues...).Offset(p.GetOffset()).Limit(p.GetLimit())
	}
}
