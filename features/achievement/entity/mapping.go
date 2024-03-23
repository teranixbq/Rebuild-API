package entity

import "recything/features/achievement/model"

func AchievementModelToAchievementCore(data model.Achievement) AchievementCore {
	return AchievementCore{
		Id:           data.Id,
		Name:         data.Name,
		TargetPoint:  data.TargetPoint,
		TotalClaimed: data.TotalClaimed,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}

}

func ListAchievementModelToAchievementCore(data []model.Achievement) []AchievementCore {
	listAchievement := []AchievementCore{}
	for _, achievement := range data {
		AchievementCore := AchievementModelToAchievementCore(achievement)
		listAchievement = append(listAchievement, AchievementCore)
	}
	return listAchievement
}

func AchievementCoreToAchievementModel(data AchievementCore) model.Achievement {
	return model.Achievement{
		Id:           data.Id,
		Name:         data.Name,
		TargetPoint:  data.TargetPoint,
		TotalClaimed: data.TotalClaimed,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ListAdminCoreToAdminModel(data []AchievementCore) []model.Achievement {
	listAchievement := []model.Achievement{}
	for _, achievement := range data {
		achievementModel := AchievementCoreToAchievementModel(achievement)
		listAchievement = append(listAchievement, achievementModel)
	}
	return listAchievement
}
