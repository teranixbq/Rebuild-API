package service

import (
	"errors"
	"math"
	"recything/features/trash_category/entity"
	"recything/mocks"
	"recything/utils/pagination"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCategory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		testData := entity.TrashCategoryCore{
			Unit:      "barang",
			TrashType: "TestTrashType",
			Point:     42,
		}

		mockRepo.On("Create", testData).Return(nil)

		err := trashCategoryService.CreateCategory(testData)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("data empty", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)
		request := entity.TrashCategoryCore{
			Unit:      "",
			TrashType: "",
			Point:     0,
		}
		err := trashCategoryService.CreateCategory(request)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

	t.Run("equal data", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		request := entity.TrashCategoryCore{
			Unit:      "pcs",
			TrashType: "plastik",
			Point:     10,
		}

		err := trashCategoryService.CreateCategory(request)
		assert.NotEqualValues(t, []string{"barang", "kilogram"}, request.Unit)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

}

func TestGetAllCategory_Success(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		expectedData := []entity.TrashCategoryCore{
			{
				ID:        "1",
				TrashType: "Plastic",
				Point:     5,
				Unit:      "Kg",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		expectedPagination := pagination.PageInfo{
			Limit:       10,
			CurrentPage: 1,
			LastPage:    int(math.Ceil(float64(len(expectedData)) / 10)),
		}

		expectedCount := len(expectedData)
		mockRepo.On("FindAll", mock.Anything, mock.Anything, mock.Anything).Return(expectedData, expectedPagination, expectedCount, nil)

		data, paginationInfo, count, err := trashCategoryService.GetAllCategory("1", "10", "")

		assert.NoError(t, err)
		assert.Equal(t, expectedData, data)
		assert.Equal(t, expectedPagination, paginationInfo)
		assert.Equal(t, expectedCount, count)

		mockRepo.AssertExpectations(t)
	})

	t.Run("limit error", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)
		_, _, _, err := trashCategoryService.GetAllCategory("1", "A", "")
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("count error", func(t *testing.T) {
		expectedData := []entity.TrashCategoryCore{
			{
				ID:        "1",
				TrashType: "Plastic",
				Point:     5,
				Unit:      "Kg",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		expectedPagination := pagination.PageInfo{
			Limit:       10,
			CurrentPage: 1,
			LastPage:    int(math.Ceil(float64(len(expectedData)) / 10)),
		}
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		expectedCount := len(expectedData)
		mockRepo.On("FindAll", mock.Anything, mock.Anything, mock.Anything).Return(expectedData, expectedPagination, expectedCount, errors.New("invalid count"))

		_, _, count, err := trashCategoryService.GetAllCategory("1", "2", "")
		assert.Error(t, err)
		assert.NotEqual(t, count, expectedCount)
	})

}

func TestGetById_Success(t *testing.T) {

	t.Run("succes", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		expectedData := entity.TrashCategoryCore{
			ID:        "1",
			TrashType: "Plastik",
			Point:     50,
			Unit:      "kilogram",
		}

		mockRepo.On("GetById", "1").Return(expectedData, nil)
		data, err := trashCategoryService.GetById("1")
		assert.NoError(t, err)
		assert.EqualValues(t, expectedData, data)
		mockRepo.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		expectedData := entity.TrashCategoryCore{
			ID:        "1",
			TrashType: "Plastik",
			Point:     50,
			Unit:      "kilogram",
		}

		mockRepo.On("GetById", "2").Return(expectedData, errors.New("not found"))
		_, err := trashCategoryService.GetById("2")
		assert.Error(t, err)
		assert.NotEqual(t, expectedData.ID, "2")
		mockRepo.AssertExpectations(t)
	})

}

func TestDeleteCategory_Success(t *testing.T) {
	data := entity.TrashCategoryCore{
		ID:        "1",
		TrashType: "Plastik",
		Point:     5,
		Unit:      "kilogram",
	}
	t.Run("succes", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		mockRepo.On("Delete", "1").Return(nil)
		err := trashCategoryService.DeleteCategory("1")

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		categoryId := "2"
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)
		mockRepo.On("Delete", categoryId).Return(errors.New("data id tidak ditemukan"))

		err := trashCategoryService.DeleteCategory(categoryId)
		assert.Error(t, err)
		assert.NotEqual(t, data.ID, categoryId)
		mockRepo.AssertExpectations(t)
	})

}

func TestUpdateCategory(t *testing.T) {

	data := entity.TrashCategoryCore{
		ID:        "1",
		TrashType: "Plastik",
		Point:     5,
		Unit:      "kilogram",
	}
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)
		dataToUpdate := entity.TrashCategoryCore{
			ID:        "1",
			TrashType: "Plastik",
			Point:     5,
			Unit:      "kilogram",
		}
		mockRepo.On("Update", "1", dataToUpdate).Return(dataToUpdate, nil)

		updatedData, err := trashCategoryService.UpdateCategory("1", dataToUpdate)

		assert.NoError(t, err)
		assert.Equal(t, dataToUpdate.ID, updatedData.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("ValidationFailure", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)
		request := entity.TrashCategoryCore{
			TrashType: "",
			Unit:      "",
		}
		categoryID := "1"

		_, err := trashCategoryService.UpdateCategory(data.ID, request)

		assert.Error(t, err)
		assert.Equal(t, categoryID, data.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("RepoError", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)
		dataToUpdate := entity.TrashCategoryCore{
			ID:        "1",
			TrashType: "Plastik",
			Point:     5,
			Unit:      "Kg",
		}
		mockRepo.On("Update", "1", dataToUpdate).Return(entity.TrashCategoryCore{}, errors.New("repository error"))
		updatedData, err := trashCategoryService.UpdateCategory("1", dataToUpdate)
		assert.Error(t, err)
		assert.Equal(t, entity.TrashCategoryCore{}, updatedData)
		mockRepo.AssertExpectations(t)
	})
}

func TestFindAllFetch_Success(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		expectedData := []entity.TrashCategoryCore{
			{
				ID:        "1",
				TrashType: "Plastic",
				Point:     5,
				Unit:      "Kg",
			},
		}

		mockRepo.On("FindAllFetch").Return(expectedData, nil)
		data, err := trashCategoryService.FindAllFetch()

		assert.NoError(t, err)
		assert.Equal(t, expectedData, data)
		mockRepo.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		mockRepo := new(mocks.TrashCategoryRepositoryInterface)
		trashCategoryService := NewTrashCategoryService(mockRepo)

		mockRepo.On("FindAllFetch").Return([]entity.TrashCategoryCore{}, errors.New("data tidak ditemukan"))
		data, err := trashCategoryService.FindAllFetch()

		assert.Error(t, err)
		assert.Equal(t, []entity.TrashCategoryCore{}, data)
		mockRepo.AssertExpectations(t)
	})
}

