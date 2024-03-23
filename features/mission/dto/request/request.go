package request

type Mission struct {
	Title            string `form:"title" json:"title"`
	MissionImage     string `form:"mission_image" json:"mission_image"`
	Point            int    `form:"point" json:"point"`
	Description      string `form:"description" json:"description"`
	Start_Date       string `form:"start_date" json:"start_date"`
	End_Date         string `form:"end_date" json:"end_date"`
	TitleStage       string `form:"title_stage" json:"title_stage"`
	DescriptionStage string `form:"description_stage" json:"description_stage"`
	// MissionStages []MissionStage `form:"mission_stages" json:"mission_stages"`
}

type MissionStage struct {
	Name             string `form:"name" json:"name"`
	DescriptionStage string `form:"description_stage" json:"description_stage"`
}

// type AddMissionStage struct {
// 	MissionID string         `json:"mission_id"`
// 	Stages    []MissionStage `json:"stages"`
// }

type RequestMissionStage struct {
	MissionStage []UpdatedMissionStage `json:"mission_stage"`
}

type UpdatedMissionStage struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// type NewMissionStage struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// }

type StatusApproval struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type Claim struct {
	MissionID string `json:"mission_id"`
}

type UploadMissionTask struct {
	UserID      string
	MissionID   string `form:"mission_id"`
	Description string `form:"description"`
}

type UpdateUploadMissionTask struct {
	Description string `form:"description"`
}
