package service

import (
	"recything/features/daily_point/entity"
	user_entity "recything/features/user/entity"
)

type dailyPointService struct {
	DailyPointRepository entity.DailyPointRepositoryInterface
}

func NewDailyPointService(daily entity.DailyPointRepositoryInterface) entity.DailyPointServiceInterface {
	return &dailyPointService{
		DailyPointRepository: daily,
	}
}

// DailyClaim implements entity.DailyPointServiceInterface.
func (dailyS *dailyPointService) DailyClaim(userId string) error {
	tx := dailyS.DailyPointRepository.DailyClaim(userId)
	if tx != nil {
		return tx
	}

	return nil
}

// PostWeekly implements entity.DailyPointServiceInterface.
func (dailyS *dailyPointService) PostWeekly() error {
	err := dailyS.DailyPointRepository.PostWeekly()
	if err != nil {
		return err
	}

	return nil
}

func (dailyS *dailyPointService) GetAllHistoryPoint(userID string) ([]map[string]interface{}, error) {
	data, errDat := dailyS.DailyPointRepository.GetAllHistoryPoint(userID)
	if errDat != nil {
		return nil, errDat
	}

	return data, nil
}

func (dailyS *dailyPointService) GetByIdHistoryPoint(userID, idTransaction string) (map[string]interface{}, error) {
	data, errDat := dailyS.DailyPointRepository.GetByIdHistoryPoint(userID, idTransaction)
	if errDat != nil {
		return nil, errDat
	}

	return data, nil
}

func (dailyS *dailyPointService) GetAllClaimedDaily(userID string) ([]user_entity.UserDailyPointsCore, error) {
	dataDaily, errGet := dailyS.DailyPointRepository.GetAllClaimedDaily(userID)
	if errGet != nil {
		return []user_entity.UserDailyPointsCore{}, errGet
	}

	return dataDaily, nil
}
