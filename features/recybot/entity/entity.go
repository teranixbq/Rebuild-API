package entity

import "time"

type RecybotCore struct {
	ID        string
	Category  string
	Question  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecybotHistories struct {
	ID        string 
	Question  string 
	UserId    string
	Answer    string
	CreatedAt time.Time
}

