package entity

import (
	"time"

	"gorm.io/gorm"
)

type FaqCore struct {
	Id                uint
	Title             string
	Description       string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeleteAt          gorm.DeletedAt
}
