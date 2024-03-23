package model

import (
	"time"

	"gorm.io/gorm"
)

type Community struct {
	Id          string `gorm:"primary key"`
	Name        string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	Location    string `gorm:"not null"`
	Members     int    `gorm:"default:0"`
	MaxMembers  int    `gorm:"not null"`
	Image       string `gorm:"not null"`
	Events      []CommunityEvent `gorm:"foreignKey:CommunityId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type CommunityEvent struct {
	Id          string `gorm:"primary key"`
	CommunityId string `gorm:"index;foreignKey:Id"`
	Title       string
	Image       string
	Description string
	Location    string
	MapLink     string
	FormLink    string
	Quota       int
	Date        string
	Status      string `gorm:"type:enum('berjalan','belum berjalan','selesai');default:'belum berjalan'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
