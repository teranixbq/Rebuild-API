package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TrashCategory struct {
	ID              string `gorm:"primary key"`
	TrashType       string `gorm:"not null;unique"`
	Point           int    `gorm:"not null"`
	Unit            string `gorm:"type:enum('barang', 'kilogram');not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (t *TrashCategory) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	t.ID = newUuid.String()
	return nil
}
