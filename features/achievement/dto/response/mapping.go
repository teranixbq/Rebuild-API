package response

import "recything/features/achievement/entity"

func AchievementCoreToAchievementResponse(data entity.AchievementCore) AchievementResponse {
	return AchievementResponse{
		Id:           data.Id,
		Name:         data.Name,
		TargetPoint:  data.TargetPoint,
		TotalClaimed: data.TotalClaimed,
	}
}

func ListAchievementCoreToAchievementResponse(data []entity.AchievementCore) []AchievementResponse {
	dataAchievement := []AchievementResponse{}
	for _, achievement := range data {
		achievementRespon := AchievementCoreToAchievementResponse(achievement)
		dataAchievement = append(dataAchievement, achievementRespon)
	}
	return dataAchievement
}


func AchievementCoreToAchievementResponseUser(data entity.AchievementCore) AchievementResponseUser {
	return AchievementResponseUser{
		Id:           data.Id,
		Name:         data.Name,
		TargetPoint:  data.TargetPoint,
		
	}
}

func ListAchievementCoreToAchievementResponseUser(data []entity.AchievementCore) []AchievementResponseUser {
	dataAchievement := []AchievementResponseUser{}
	for _, achievement := range data {
		achievementRespon := AchievementCoreToAchievementResponseUser(achievement)
		dataAchievement = append(dataAchievement, achievementRespon)
	}
	return dataAchievement
}
