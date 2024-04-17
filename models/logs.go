package models

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Message    string    `gorm:"type:text" json:"message"`
	Level      string    `gorm:"index:idx_level" json:"level"`
	Path       string    `gorm:"type:text" json:"path"`
	ResourceID uuid.UUID `gorm:"type:uuid;not null" json:"resource_id"`
	Resource   Resource  `gorm:"" json:"resource"`
	Timestamp  time.Time `json:"timestamp" gorm:"index:idx_timestamp"`
}
