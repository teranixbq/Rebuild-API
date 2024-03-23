package service

import (
	"errors"
	"log"
	"recything/features/achievement/entity"
	"recything/mocks"
	"recything/utils/constanta"

	"testing"

	"github.com/stretchr/testify/assert"
)

var dataAchievements = []entity.AchievementCore{
	{Id: 1, Name: "platinum", TargetPoint: 250000},
	{Id: 2, Name: "gold", TargetPoint: 100000},
	{Id: 3, Name: "silver", TargetPoint: 50000},
	{Id: 4, Name: "bronze", TargetPoint: 0},
}

var dataAchievement = entity.AchievementCore{
	Id: 1, Name: "platinum", TargetPoint: 250000,
}

func TestGetAllAchievement(t *testing.T) {

	t.Run("Succes GetAll", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		mockRepo.On("GetAllAchievement").Return(dataAchievements, nil)

		result, err := achievementService.GetAllAchievement()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Error", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		mockRepo.On("GetAllAchievement").Return(nil, errors.New("failed"))

		result, err := achievementService.GetAllAchievement()

		assert.Error(t, err)
		assert.Empty(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateById(t *testing.T) {
	t.Run("Sukses Update", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		achivementID := 1
		point := 249000

		expectedAchievement := dataAchievements[achivementID-1]
		mockRepo.On("FindById", achivementID).Return(expectedAchievement, nil)
		mockRepo.On("UpdateById", achivementID, point).Return(nil)

		err := achievementService.UpdateById(achivementID, point)

		assert.NoError(t, err)
		assert.Equal(t, achivementID, expectedAchievement.Id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		achivementID := 0
		point := 249000

		err := achievementService.UpdateById(achivementID, point)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		achivementID := 9
		point := 249000

		mockRepo.On("FindById", achivementID).Return(dataAchievement, errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		err := achievementService.UpdateById(achivementID, point)

		assert.Error(t, err)
		assert.NotEqual(t, achivementID, dataAchievement.Id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Point for Bronze", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		achivementID := 4
		point := 100

		mockRepo.On("FindById", achivementID).Return(dataAchievements[achivementID-1], nil)

		err := achievementService.UpdateById(achivementID, point)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Point for Silver", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		achivementID := 3
		point := 70000

		mockRepo.On("FindById", achivementID).Return(dataAchievements[achivementID-1], nil)

		err := achievementService.UpdateById(achivementID, point)
		log.Println("er", err)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Point for Gold", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		achivementID := 2
		point := 100

		mockRepo.On("FindById", achivementID).Return(dataAchievements[achivementID-1], nil)

		err := achievementService.UpdateById(achivementID, point)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Point for Platinum", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		achivementID := 1
		point := 500000

		mockRepo.On("FindById", achivementID).Return(dataAchievements[achivementID-1], nil)

		err := achievementService.UpdateById(achivementID, point)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Repository", func(t *testing.T) {
		mockRepo := new(mocks.AchievementRepositoryInterface)
		achievementService := NewAchievementService(mockRepo)

		achivementID := 1
		point := 249000

		mockRepo.On("FindById", achivementID).Return(dataAchievements[achivementID-1],nil)
		mockRepo.On("UpdateById", achivementID, point).Return(errors.New("failed")) 

		err := achievementService.UpdateById(achivementID, point)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

}
