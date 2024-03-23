package response

type AchievementResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	TargetPoint  int    `json:"target_point"`
	TotalClaimed int    `json:"total_claimed"`
}

type AchievementResponseUser struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	TargetPoint  int    `json:"target_point"`
}
