package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"-"`
	Password  string    `json:"-"`
	Admin     bool      `gorm:"default:false" json:"-"`
	Timestamp time.Time `json:"timestamp" gorm:"index:idx_timestamp"`
}
