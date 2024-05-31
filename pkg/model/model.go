package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Rating struct {
	gorm.Model
	UserUUID uuid.UUID `gorm:"not null" json:"user_uuid"`
	Rating   uint8     `gorm:"not null" json:"rating"`
	Text     string    `json:"text"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
