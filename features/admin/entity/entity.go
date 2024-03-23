package entity

import "time"

type AdminCore struct {
	Id              string
	Fullname        string
	Image           string
	Role            string
	Email           string
	Password        string
	ConfirmPassword string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
