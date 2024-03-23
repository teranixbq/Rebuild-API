package service

import (
	"errors"
	"recything/features/achievement/entity"
	"recything/utils/constanta"
)

type achievementService struct {
	achievementRepository entity.AchievementRepositoryInterface
}

func NewAchievementService(achievement entity.AchievementRepositoryInterface) entity.AchievementServiceInterface {
	return &achievementService{
		achievementRepository: achievement,
	}
}

// GetAllAchievement implements entity.AchievementServiceInterface.
func (as *achievementService) GetAllAchievement() ([]entity.AchievementCore, error) {
	achievement, err := as.achievementRepository.GetAllAchievement()
	if err != nil {
		return nil, err
	}

	return achievement, nil
}

// UpdateById implements entity.AchievementServiceInterface.
func (as *achievementService) UpdateById(id int, point int) error {
	if id == 0 {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	dataAchievement, errFind := as.achievementRepository.FindById(id)
	if errFind != nil {
		return errFind
	}

	if dataAchievement.Name == "bronze" {
	
		if point > 0 {
			return errors.New("error: target point lencana bronze tidak boleh lebih dari 0")
		}
	} else if dataAchievement.Name == "silver" {
		if point > 50000 || point <= 0 {
			return errors.New("error: target point lencana silver tidak boleh lebih dari 50000 atau kurang dari 0")
		}
	} else if dataAchievement.Name == "gold" {
		if point > 100000 || point <= 50000 {
			return errors.New("error: target point lencana gold tidak boleh lebih dari 100000 atau kurang dari lencana sebelumnya")
		}
	} else if dataAchievement.Name == "platinum" {
		if point > 250000 || point <= 100000 {
			return errors.New("error: target point lencana platinum tidak boleh lebih dari 250000 atau kurang dari lencana sebelumnya")
		}
	}

	if point == dataAchievement.TargetPoint {

	} else {
		errUpdate := as.achievementRepository.UpdateById(id, point)
		if errUpdate != nil {
			return errUpdate
		}
	}

	return nil
}
