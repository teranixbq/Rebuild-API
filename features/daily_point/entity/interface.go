package entity

import user_entity "recything/features/user/entity"

type DailyPointRepositoryInterface interface {
	PostWeekly() error
	DailyClaim(userId string) error
	GetAllHistoryPoint(userID string) ([]map[string]interface{}, error)
	GetByIdHistoryPoint(userID,idTransaction string) (map[string]interface{}, error)
	GetAllClaimedDaily(userID string) ([]user_entity.UserDailyPointsCore, error)
}

type DailyPointServiceInterface interface {
	PostWeekly() error
	DailyClaim(userId string) error
	GetAllHistoryPoint(userID string) ([]map[string]interface{}, error)
	GetByIdHistoryPoint(userID,idTransaction string) (map[string]interface{}, error) 
	GetAllClaimedDaily(userID string) ([]user_entity.UserDailyPointsCore, error)
}
