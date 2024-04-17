package utils

import (
	"github.com/csivitu/csi-logger/schemas"
	"gorm.io/gorm"
)

func FilterLogs(db *gorm.DB, urlParams schemas.LogFetchSchema) *gorm.DB {
	query := db

	query = query.Limit(urlParams.Limit)

	query = query.Offset((urlParams.Page - 1) * 20)
	
	return query
}
