package response

import (
	"time"
)

type Mission struct {
	ID           string `json:"id"`
	Title        string `json:"name"`
	Creator      string `json:"creator"`
	Status       string `json:"status"`
	MissionImage string `json:"mission_image"`
	Point        int    `json:"point"`
	Description  string `json:"description"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	// MissionStages []MissionStage `json:"mission_stages"`
	TitleStage       string    `json:"title_stage,omitempty"`
	DescriptionStage string    `json:"description_stage,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type MissionStage struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UploadMissionTask struct {
	ID          string               `json:"id"`
	UserID      string               `json:"user_id"`
	User        string               `json:"user"`
	MissionID   string               `json:"mission_id"`
	MissionName string               `json:"mission_name"`
	Description string               `json:"description"`
	Reason      string               `json:"reason,omitempty"`
	Images      []ImageUploadMission `json:"images"`
	Status      string               `json:"status"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type ImageUploadMission struct {
	ID                  string    `json:"id"`
	UploadMissionTaskID string    `json:"upload_mission_task_id"`
	Image               string    `json:"image"`
	CreatedAt           time.Time `json:"created_at"`
}

type Proof struct {
	ID   string `json:"id"`
	File string `json:"file"`
}

type UploadMissionTaskResponse struct {
	ID        string `json:"id"`
}

type MissionHistories struct {
	MissionID      string `json:"mission_id"`
	ClaimedID      string `json:"claimed_id,omitempty"`
	TransactionID  string `json:"transaction_id,omitempty"`
	Title          string `json:"title"`
	StatusApproval string `json:"status_approval,omitempty"`
	StatusMission  string `json:"status_mission,omitempty"`
	MissionImage   string `json:"mission_image"`
	// Reason         string `json:"reason,omitempty"`
	Point       int    `json:"point"`
	Description string `json:"description"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	// MissionStages  []MissionStage `json:"mission_stages,omitempty"`
	TitleStage       string    `json:"title_stage,omitempty"`
	DescriptionStage string    `json:"description_stage,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}
