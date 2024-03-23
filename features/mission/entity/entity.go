package entity

import (
	"time"
)

type Mission struct {
	ID               string
	Title            string
	Creator          string
	Status           string
	AdminID          string
	MissionImage     string
	Point            int
	Description      string
	StartDate        string
	EndDate          string
	TitleStage       string
	DescriptionStage string
	// MissionStages []MissionStage
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MissionStage struct {
	MissionID   string
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ClaimedMission struct {
	ID        string
	UserID    string
	MissionID string
	Claimed   bool
	CreatedAt time.Time
}

// User Upload

type UploadMissionTaskCore struct {
	ID          string
	UserID      string
	User        string
	MissionID   string
	MissionName string
	Description string
	Reason      string
	Images      []ImageUploadMissionCore
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ImageUploadMissionCore struct {
	ID                  string
	UploadMissionTaskID string
	Image               string
	CreatedAt           time.Time
}

// histories

type MissionHistories struct {
	MissionID        string
	ClaimedID        string
	TransactionID    string
	Title            string
	StatusApproval   string
	StatusMission    string
	MissionImage     string
	Reason           string
	Point            int
	Description      string
	StartDate        string
	EndDate          string
	TitleStage       string
	DescriptionStage string
	// MissionStages  []MissionStage
	CreatedAt time.Time
	UpdatedAt time.Time
}
