package entity

import "time"

type RecybotCore struct {
	ID        string
	Category  string
	Question  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecybbotHistories struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

