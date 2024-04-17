package models

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Message   string    `gorm:"type:text" json:"message"`
	Path      string    `gorm:"type:text" json:"path"`
	Resource  string    `gorm:"index:idx_resource" json:"resource"`
	Timestamp time.Time `json:"timestamp" gorm:"index:idx_timestamp"`
}
