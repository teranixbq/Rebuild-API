package model

import (
	"time"

	"gorm.io/gorm"
)

type Faq struct {
	Id                uint         `gorm:"primaryKey;not null" json:"id"`
	Title             string         `gorm:"not null" json:"title"`
	Description       string         `gorm:"not null" json:"description"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
}
