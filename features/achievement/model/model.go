package model

import (
	"time"

	"gorm.io/gorm"
)

type Achievement struct {
	Id           int    `gorm:"primary key"`
	Name         string `gorm:"type:enum('platinum', 'gold', 'silver', 'bronze');not null;unique;index"`
	TargetPoint  int    `gorm:"not null"`
	TotalClaimed int    
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
