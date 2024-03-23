package service

import (
	"errors"
	entity "recything/features/recybot/entity"
	"recything/mocks"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"

	"github.com/stretchr/testify/assert"

	"testing"
)

var dataPrompt = []entity.RecybotCore{
	{ID: "1", Category: "informasi", Question: "jakarta ada di pulau jawa"},
	{ID: "2", Category: "batasan", Question: "tidak boleh berkata kasar"},
	{ID: "3", Category: "sampah organik", Question: "daun merupakan sampah organik"},
}

func TestCreateData(t *testing.T) {
	t.Run("Succes Create", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		requestBody := entity.RecybotCore{
			Category: "informasi",
			Question: "matahari terbit di timur",
		}
		mockRepo.On("Create", requestBody).Return(requestBody, nil)

		data, err := recybotService.CreateData(requestBody)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockRepo.AssertExpectations(t)

	})
	t.Run("Data Empty", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		requestBody := entity.RecybotCore{
			Category: "",
			Question: "",
		}

		data, err := recybotService.CreateData(requestBody)

		assert.Error(t, err)
		assert.Empty(t, data)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Equal Data Category", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		requestBody := entity.RecybotCore{
			Category: "robot",
			Question: "robot itu menggunakan baterai",
		}

		data, err := recybotService.CreateData(requestBody)

		assert.Error(t, err)
		assert.Empty(t, data)
		assert.NotEqualValues(t, []string{"sampah anorganik", "sampah organik", "informasi", "batasan"}, requestBody.Category)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		requestBody := entity.RecybotCore{
			Category: "informasi",
			Question: "matahari terbit di timur",
		}

		mockRepo.On("Create", requestBody).Return(entity.RecybotCore{}, errors.New("Repository error"))

		data, err := recybotService.CreateData(requestBody)

		assert.Error(t, err)
		assert.Empty(t, data)

		mockRepo.AssertExpectations(t)
	})
}

func TestFindAllData(t *testing.T) {

	t.Run("Success FindAll", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		mockRepo.On("FindAll", 1, 10, "", "").Return(dataPrompt, pagination.PageInfo{}, helper.CountPrompt{}, nil)

		data, pageInfo, count, err := recybotService.FindAllData("", "", "", "")

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.NotNil(t, count)
		assert.NotNil(t, pageInfo)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Wrong Limit", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		data, _, _, err := recybotService.FindAllData("", "", "", "50")

		assert.Error(t, err)
		assert.Empty(t, data)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		mockRepo.On("FindAll", 1, 10, "", "").Return(nil, pagination.PageInfo{}, helper.CountPrompt{}, errors.New("Repository error"))

		data, _, _, err := recybotService.FindAllData("", "", "1", "10")

		assert.Error(t, err)
		assert.Empty(t, data)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Succes GetById", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		recybotID := "1"

		mockRepo.On("GetById", recybotID).Return(dataPrompt[0], nil)

		result, err := recybotService.GetById(recybotID)

		assert.NoError(t, err)
		assert.Equal(t, recybotID, dataPrompt[0].ID)
		assert.NotNil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		recybotID := "2"
		mockRepo.On("GetById", recybotID).Return(entity.RecybotCore{}, errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		result, err := recybotService.GetById(recybotID)

		assert.Error(t, err)
		assert.NotEqual(t, recybotID, dataPrompt[0].ID)
		assert.Empty(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Succes DeleteData", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		recybotID := "1"

		mockRepo.On("Delete", recybotID).Return(nil)

		err := recybotService.DeleteData(recybotID)

		assert.NoError(t, err)
		assert.Equal(t, recybotID, dataPrompt[0].ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		recybotID := "2"
		mockRepo.On("Delete", recybotID).Return(errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		err := recybotService.DeleteData(recybotID)

		assert.Error(t, err)
		assert.NotEqual(t, recybotID, dataPrompt[0].ID)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	t.Run("Succes UpdateData", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		requestBody := entity.RecybotCore{
			Category: "informasi",
			Question: "mobil beroda empat",
		}

		recybotID := "2"
		mockRepo.On("Update", recybotID, requestBody).Return(requestBody, nil)

		data, err := recybotService.UpdateData(recybotID, requestBody)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, recybotID, dataPrompt[1].ID)
		mockRepo.AssertExpectations(t)

	})
	t.Run("Data Empty", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		requestBody := entity.RecybotCore{
			Category: "",
			Question: "",
		}
		recybotID := "2"
		data, err := recybotService.UpdateData(recybotID, requestBody)

		assert.Error(t, err)
		assert.Empty(t, data)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Equal Data Category", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		requestBody := entity.RecybotCore{
			Category: "robot",
			Question: "robot itu menggunakan baterai",
		}

		recybotID := "2"
		data, err := recybotService.UpdateData(recybotID, requestBody)

		assert.Error(t, err)
		assert.Empty(t, data)
		assert.NotEqualValues(t, []string{"sampah anorganik", "sampah organik", "informasi", "batasan"}, requestBody.Category)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		requestBody := entity.RecybotCore{
			Category: "informasi",
			Question: "sampah",
		}
		recybotID := "2"
		mockRepo.On("Update", recybotID, requestBody).Return(entity.RecybotCore{}, errors.New("Repository error"))

		data, err := recybotService.UpdateData(recybotID, requestBody)

		assert.Error(t, err)
		assert.Empty(t, data)

		mockRepo.AssertExpectations(t)
	})
}

func TestGetPrompt(t *testing.T) {
	t.Run("Error API Key", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		question := "contoh 3 sampah organik"
		mockRepo.On("GetAll").Return(dataPrompt, nil)

		result, err := recybotService.GetPrompt(question)

		assert.Error(t, err)
		assert.Empty(t,result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Get All Prompt", func(t *testing.T) {
		mockRepo := new(mocks.RecybotRepositoryInterface)
		recybotService := NewRecybotService(mockRepo)

		question := "contoh 3 sampah organik"
		mockRepo.On("GetAll").Return(dataPrompt, errors.New("Failed Get Data"))

		result, err := recybotService.GetPrompt(question)

		assert.Error(t, err)
		assert.Empty(t, result)
		mockRepo.AssertExpectations(t)
	})
}
