package models

import (
	"time"

	"github.com/google/uuid"
)

type Resource struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	HostedURL string    `gorm:"unique;not null" json:"hosted_url"`
	APIKey    string    `gorm:"unique;not null" json:"api_key"`
	Timestamp time.Time `json:"timestamp" gorm:"index:idx_timestamp"`
}
