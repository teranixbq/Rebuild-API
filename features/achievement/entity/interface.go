package entity

type AchievementRepositoryInterface interface {
	GetAllAchievement() ([]AchievementCore, error)
	UpdateById(id int, point int) error 
	FindById(id int) (AchievementCore, error)
}

type AchievementServiceInterface interface {
	GetAllAchievement() ([]AchievementCore, error)
	UpdateById(id int, point int) error 
}