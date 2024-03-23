package model

import (
	dropPoint "recything/features/drop-point/model"
	trashCategory "recything/features/trash_category/model"
	"time"

	"gorm.io/gorm"
)

type TrashExchange struct {
	Id            string  `gorm:"primary key"`
	Name          string  `gorm:"not null"`
	EmailUser     string  `gorm:"index"`
	DropPointName string  `gorm:"not null"`
	TotalIncome   int     `gorm:"not null"`
	TotalPoint    int     `gorm:"not null"`
	TotalUnit     float64 `gorm:"type:decimal(10,1);not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	DropPointId          string                `gorm:"index"`
	DropPoint            dropPoint.DropPoints  `gorm:"foreignKey:DropPointId"`
	TrashExchangeDetails []TrashExchangeDetail `gorm:"foreignKey:TrashExchangeId;constraint:OnDelete:CASCADE;"`
}

type TrashExchangeDetail struct {
	Id              string  `gorm:"primary key"`
	TrashExchangeId string  `gorm:"index"`
	Amount          float64 `gorm:"not null"`
	Unit            string  `gorm:"not null"`
	TotalPoints     int     `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt              `gorm:"index"`
	Type            string                      `gorm:"index"`
	TrashCategory   trashCategory.TrashCategory `gorm:"foreignKey:Type;references:TrashType"`
}
