package entity

import "time"

type RecybotCore struct {
	ID        string
	Category  string
	Question  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

