package entity

import (
	"time"
)

type AchievementCore struct {
	Id           int
	Name         string
	TargetPoint  int
	TotalClaimed int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
