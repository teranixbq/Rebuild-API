package entity

import (
	"time"
)

type CommunityCore struct {
	Id          string
	Name        string
	Description string
	Location    string
	Members     int
	MaxMembers  int
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CommunityEventCore struct {
	Id          string
	CommunityId string
	Title       string
	Image       string
	Description string
	Location    string
	MapLink     string
	FormLink    string
	Quota       int
	Date        string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
