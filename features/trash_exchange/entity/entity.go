package entity

import (
	"time"
)

type TrashExchangeCore struct {
	Id                   string                    `json:"id"`
	Name                 string                    `json:"name"`
	EmailUser            string                    `json:"email"`
	DropPointName        string                    `json:"drop_point_name"`
	DropPointId          string                    `json:"drop_point_id"`
	TotalIncome          int                       `json:"total_income"`
	TotalPoint           int                       `json:"total_point"`
	TotalUnit            float64                   `json:"total_unit"`
	CreatedAt            time.Time                 `json:"created_at"`
	UpdatedAt            time.Time                 `json:"updated_at"`
	DeletedAt            *time.Time                `json:"deleted_at,omitempty"`
	TrashExchangeDetails []TrashExchangeDetailCore `json:"trash_exchange_details"`
}

type TrashExchangeDetailCore struct {
	Id              string     `json:"id"`
	TrashExchangeId string     `json:"trash_exchange_id"`
	TrashType       string     `json:"trash_type"`
	Amount          float64    `json:"amount"`
	Unit            string     `json:"unit"`
	TotalPoints     int        `json:"total_points"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}
