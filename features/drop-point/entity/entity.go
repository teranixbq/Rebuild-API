package entity

import (
	"time"

	"gorm.io/gorm"
)

type DropPointsCore struct {
	Id        string
	Name      string
	Address   string
	Latitude  float64
	Longitude float64
	Schedule  []ScheduleCore
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type ScheduleCore struct {
	Id           string
	DropPointsID string
	Day          string
	OpenTime     string
	CloseTime    string
	Closed       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
