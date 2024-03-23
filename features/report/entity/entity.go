package entity

import (
	"time"
)

type ReportCore struct {
	ID                   string
	ReportType           string
	UserId               string
	Longitude            float64
	Latitude             float64
	Location             string
	AddressPoint         string
	Status               string
	TrashType            string
	Description          string
	ScaleType            string
	InsidentDate         string
	InsidentTime         string
	CompanyName          string
	DangerousWaste       bool
	RejectionDescription string
	Images               []ImageCore
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type ImageCore struct {
	ID        string
	ReportID  string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
