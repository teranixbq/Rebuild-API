package model

import (
	"time"

	"gorm.io/gorm"
)

type DropPoints struct {
	Id        string      `gorm:"primary key;type:varchar(191)"`
	Name      string      `gorm:"not null"`
	Address   string      `gorm:"not null"`
	Latitude  float64     `gorm:"not null"`
	Longitude float64     `gorm:"not null"`
	Schedule  []Schedules `gorm:"foreignKey:DropPointsID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Schedules struct {
	Id           string `gorm:"primary key;type:varchar(191)"`
	DropPointsID string
	Day          string
	OpenTime     string
	CloseTime    string
	Closed       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
