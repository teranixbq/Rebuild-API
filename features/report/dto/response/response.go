package response

import (
	"time"
)

type ReportCreateResponse struct {
	Id                   string          `json:"Id,omitempty"`
	ReportType           string          `json:"report_type,omitempty"`
	Longitude            float64         `json:"longitude,omitempty"`
	Latitude             float64         `json:"latitude,omitempty"`
	Location             string          `json:"location,omitempty"`
	Description          string          `json:"description,omitempty"`
	AddressPoint         string          `json:"address_point,omitempty"`
	Status               string          `json:"status,omitempty"`
	Images               []ImageResponse `json:"images,omitempty"`
	TrashType            string          `json:"trash_type,omitempty"`
	ScaleType            string          `json:"scale_type,omitempty"`
	InsidentDate         string          `json:"insident_date,omitempty"`
	InsidentTime         string          `json:"insident_time,omitempty"`
	DangerousWaste       bool            `json:"dangerous_waste,omitempty"`
	RejectionDescription string          `json:"rejection_description,omitempty"`
	CompanyName          string          `json:"company_name,omitempty"`
	CreatedAt            time.Time       `json:"created_at,omitempty"`
	UpdatedAt            time.Time       `json:"updated_at,omitempty"`
}

type ImageResponse struct {
	ID        string    `json:"id,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ReportDetails struct {
	Id           string `json:"id,omitempty"`
	ReportType   string `json:"report_type,omitempty"`
	Fullname     string `json:"name,omitempty"`
	Location     string `json:"location,omitempty"`
	InsidentDate string `json:"insident_date,omitempty"`
	Status       string `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ReportDetailsById struct {
	Id                   string          `json:"Id,omitempty"`
	ReportType           string          `json:"report_type,omitempty"`
	Longitude            float64         `json:"longitude,omitempty"`
	Latitude             float64         `json:"latitude,omitempty"`
	Location             string          `json:"location,omitempty"`
	Description          string          `json:"description,omitempty"`
	AddressPoint         string          `json:"address_point,omitempty"`
	Status               string          `json:"status,omitempty"`
	Images               []ImageResponse `json:"images,omitempty"`
	TrashType            string          `json:"trash_type,omitempty"`
	ScaleType            string          `json:"scale_type,omitempty"`
	InsidentDate         string          `json:"insident_date,omitempty"`
	InsidentTime         string          `json:"insident_time,omitempty"`
	DangerousWaste       bool            `json:"dangerous_waste,omitempty"`
	RejectionDescription string          `json:"rejection_description,omitempty"`
	CompanyName          string          `json:"company_name,omitempty"`
	Fullname             string          `json:"fullname"`
}