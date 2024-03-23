package request

import "recything/features/achievement/entity"

func AchievementRequestToAchievementCore(data AchievementRequest) entity.AchievementCore {
	return entity.AchievementCore{
		TargetPoint: data.TargetPoint,
	}
}
