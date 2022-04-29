package models

// gorm:"index:search_member" creates a composite key for lookups
type Service struct {
	BaseModel
	Title           string           `gorm:"index:,class:FULLTEXT" json:"title"` // FULLTEXT index algo
	Description     string           `gorm:"type:text;" json:"description"`
	UserID          string           `json:"user_id" gorm:"size:191"`
	ServiceVersions []ServiceVersion `json:"service_versions" gorm:"foreignKey:ServiceID"`
}

// Custom scanner for GetServiceRawMySQL
// func (s *ServiceVersions) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
// 	}
// 	return json.Unmarshal(bytes, &s)

// }
