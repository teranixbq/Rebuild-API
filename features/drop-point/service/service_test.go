package service

import (
	"errors"
	"recything/features/drop-point/entity"
	"recything/mocks"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewDropPointService(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := NewDropPointService(mockRepository)
	assert.Equal(t, mockRepository, service.(*dropPointService).dropPointRepository)
}

// Test case for valid drop point creation
func TestCreateDropPointValidData(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "senin", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "selasa", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "rabu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "kamis", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "jumat", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "sabtu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "minggu", OpenTime: "08:00", CloseTime: "14:00"},
		},
	}

	mockRepository.On("CreateDropPoint", validData).Return(nil)

	err := service.CreateDropPoint(validData)

	assert.Nil(t, err)
	mockRepository.AssertExpectations(t)
}

// Test case for empty fields
func TestCreateDropPointEmptyFieldsDropPoints(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	emptyData := entity.DropPointsCore{}

	err := service.CreateDropPoint(emptyData)

	assert.Error(t, err)
	mockRepository.AssertNotCalled(t, "CreateDropPoint")
}

func TestCreateDropPointEmptyDataSchedules(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	emptyData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "", OpenTime: "", CloseTime: ""},
		},
	}

	err := service.CreateDropPoint(emptyData)

	assert.NotNil(t, err)
	mockRepository.AssertExpectations(t)
}

func TestCreateDropPointInvalidTime(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	invalidTimeData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "senin", OpenTime: "14:00", CloseTime: "08:00"}, // Invalid opening time
		},
	}

	err := service.CreateDropPoint(invalidTimeData)

	assert.NotNil(t, err)
	mockRepository.AssertExpectations(t)
}

// Test case for invalid latitude/longitude
func TestCreateDropPointInvalidLatLong(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	invalidLatLongData := entity.DropPointsCore{
		Name:      "Invalid Location",
		Address:   "456 Invalid Street",
		Latitude:  91.0,
		Longitude: 180.0,
	}

	err := service.CreateDropPoint(invalidLatLongData)

	assert.Error(t, err)
	mockRepository.AssertNotCalled(t, "CreateDropPoint")
}

// Test case for duplicate schedule days
func TestCreateDropPointDuplicateDays(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	duplicateDaysData := entity.DropPointsCore{
		Name:      "Duplicate Days",
		Address:   "789 Duplicate Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "senin", OpenTime: "08:00", CloseTime: "15:00"},
			{Day: "senin", OpenTime: "09:00", CloseTime: "16:00"},
		},
	}

	err := service.CreateDropPoint(duplicateDaysData)

	assert.Error(t, err)
	mockRepository.AssertNotCalled(t, "CreateDropPoint")
}

func TestCreateDropPointErrorDays(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	duplicateDaysData := entity.DropPointsCore{
		Name:      "Duplicate Days",
		Address:   "789 Duplicate Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "monday", OpenTime: "08:00", CloseTime: "15:00"},
			{Day: "kamis", OpenTime: "09:00", CloseTime: "16:00"},
		},
	}

	err := service.CreateDropPoint(duplicateDaysData)

	assert.Error(t, err)
	mockRepository.AssertNotCalled(t, "CreateDropPoint")
}

func TestCreateDropPointRepositoryError(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	data := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "senin", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "selasa", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "rabu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "kamis", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "jumat", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "sabtu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "minggu", OpenTime: "08:00", CloseTime: "14:00"},
		},
	}

	// Mengatur behavior mock repository
	expectedError := errors.New("Repository error")
	mockRepository.On("CreateDropPoint", data).Return(expectedError)

	err := service.CreateDropPoint(data)

	assert.EqualError(t, err, expectedError.Error())
	mockRepository.AssertExpectations(t)
}

func TestGetAllDropPoint_ValidData(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	page := 1
	limit := 10
	search := "example"

	dummyDropPoints := []entity.DropPointsCore{}
	dummyPageInfo := pagination.PageInfo{}
	dummyCount := 42

	mockRepository.On("GetAllDropPoint", page, limit, search).Return(dummyDropPoints, dummyPageInfo, dummyCount, nil)
	resultDropPoints, resultPageInfo, resultCount, err := service.GetAllDropPoint(page, limit, search)

	mockRepository.AssertCalled(t, "GetAllDropPoint", page, limit, search)
	assert.Nil(t, err)
	assert.Equal(t, dummyDropPoints, resultDropPoints)
	assert.Equal(t, dummyPageInfo, resultPageInfo)
	assert.Equal(t, dummyCount, resultCount)
}

func TestGetAllDropPoint_LimitExceeds(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	page := 1
	limit := 15
	search := "example"

	_, _, _, err := service.GetAllDropPoint(page, limit, search)
	assert.NotNil(t, err)
	assert.Equal(t, "limit tidak boleh lebih dari 10", err.Error())
}

func TestGetAllDropPointRepositoryError(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	page := 1
	limit := 10
	search := "example"

	expectedError := errors.New("repository error")
	mockRepository.On("GetAllDropPoint", page, limit, search).Return(nil, pagination.PageInfo{}, 0, expectedError)
	_, _, _, err := service.GetAllDropPoint(page, limit, search)

	mockRepository.AssertCalled(t, "GetAllDropPoint", page, limit, search)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}

func TestGetDropPointByIdInvalidID(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	invalidID := ""

	resultDropPoint, err := service.GetDropPointById(invalidID)
	mockRepository.AssertNotCalled(t, "GetDropPointById")

	assert.NotNil(t, err)
	assert.Equal(t, constanta.ERROR_ID_INVALID, err.Error())
	assert.Equal(t, entity.DropPointsCore{}, resultDropPoint)
}

func TestGetDropPointByIdRepositoryError(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validID := "validID"

	expectedError := errors.New("repository error")
	mockRepository.On("GetDropPointById", validID).Return(entity.DropPointsCore{}, expectedError)

	resultDropPoint, err := service.GetDropPointById(validID)
	mockRepository.AssertCalled(t, "GetDropPointById", validID)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, entity.DropPointsCore{}, resultDropPoint)
}

func TestGetDropPointByIdValidData(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validID := "validID"
	dummyDropPoint := entity.DropPointsCore{}

	mockRepository.On("GetDropPointById", validID).Return(dummyDropPoint, nil)
	resultDropPoint, err := service.GetDropPointById(validID)

	mockRepository.AssertCalled(t, "GetDropPointById", validID)

	assert.Nil(t, err)
	assert.Equal(t, dummyDropPoint, resultDropPoint)
}

func TestUpdateDropPointByIdEmptyData(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validID := "validID"
	emptyData := entity.DropPointsCore{}
	err := service.UpdateDropPointById(validID, emptyData)

	mockRepository.AssertNotCalled(t, "UpdateDropPointById")

	assert.NotNil(t, err)
	assert.Equal(t, "error : harap lengkapi data dengan benar", err.Error())
}

func TestUpdateDropPointByIdInvalidLatLong(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validID := "validID"
	invalidLatLongData := entity.DropPointsCore{
		Name:      "Invalid Location",
		Address:   "456 Invalid Street",
		Latitude:  91.0,
		Longitude: 180.0,
	}

	err := service.UpdateDropPointById(validID, invalidLatLongData)
	mockRepository.AssertNotCalled(t, "UpdateDropPointById")

	assert.NotNil(t, err)
	assert.Equal(t, "bukan latitude", err.Error())
}

func TestUpdateDropPointByIdInvalidDay(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validID := "validID"
	invalidDayData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "monday", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "invalid", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "rabu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "kamis", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "jumat", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "sabtu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "minggu", OpenTime: "08:00", CloseTime: "14:00"},
		},
	}

	mockRepository.AssertNotCalled(t, "CreateDropPoint")
	err := service.UpdateDropPointById(validID, invalidDayData)

	assert.Error(t, err)
	mockRepository.AssertNotCalled(t, "UpdateDropPointById")
}

func TestUpdateDropPointByIdEmptyDataSchedules(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validID := "validID"
	emptyData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "", OpenTime: "", CloseTime: ""},
		},
	}

	err := service.UpdateDropPointById(validID, emptyData)

	assert.NotNil(t, err)
	mockRepository.AssertExpectations(t)
}

func TestUpdateDropPointByIdInvalidTime(t *testing.T) {

	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	invalidData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "senin", OpenTime: "25:00", CloseTime: "14:00"},
		},
	}

	err := service.UpdateDropPointById("some_id", invalidData)
	assert.NotNil(t, err)
	assert.Equal(t, "format waktu buka tidak valid", err.Error())
}

func TestUpdateDropPointByIdDuplicateDay(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	invalidData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "senin", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "senin", OpenTime: "09:00", CloseTime: "15:00"},
		},
	}

	err := service.UpdateDropPointById("some_id", invalidData)

	assert.NotNil(t, err)
	assert.Equal(t, "hari tidak boleh sama", err.Error())
}

func TestUpdateDropPointByIdSuccessfulUpdate(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "senin", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "selasa", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "rabu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "kamis", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "jumat", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "sabtu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "minggu", OpenTime: "08:00", CloseTime: "14:00"},
		},
	}

	mockRepository.On("UpdateDropPointById", "some_id", validData).Return(nil)
	err := service.UpdateDropPointById("some_id", validData)

	assert.Nil(t, err)
	mockRepository.AssertExpectations(t)
}

func TestUpdateDropPointById_RepositoryError(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	validData := entity.DropPointsCore{
		Name:      "Test Drop Point",
		Address:   "123 Main Street",
		Latitude:  12.345678,
		Longitude: 98.765432,
		Schedule: []entity.ScheduleCore{
			{Day: "senin", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "selasa", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "rabu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "kamis", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "jumat", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "sabtu", OpenTime: "08:00", CloseTime: "14:00"},
			{Day: "minggu", OpenTime: "08:00", CloseTime: "14:00"},
		},
	}

	expectedError := errors.New("repository error")

	mockRepository.On("UpdateDropPointById", "some_id", validData).Return(expectedError)
	err := service.UpdateDropPointById("some_id", validData)

	assert.Equal(t, expectedError, err)
	mockRepository.AssertExpectations(t)
}

func TestDeleteDropPointById_SuccessfulDelete(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	mockRepository.On("DeleteDropPointById", "some_id").Return(nil)
	err := service.DeleteDropPointById("some_id")

	assert.Nil(t, err)
	mockRepository.AssertExpectations(t)
}

func TestDeleteDropPointByIdRepositoryError(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	expectedError := errors.New("repository error")
	mockRepository.On("DeleteDropPointById", "some_id").Return(expectedError)
	err := service.DeleteDropPointById("some_id")

	assert.Equal(t, expectedError, err)
	mockRepository.AssertExpectations(t)
}

func TestDeleteDropPointById_InvalidID(t *testing.T) {
	mockRepository := mocks.NewDropPointRepositoryInterface(t)
	service := dropPointService{dropPointRepository: mockRepository}

	err := service.DeleteDropPointById("")

	assert.Equal(t, constanta.ERROR_ID_INVALID, err.Error())
}

func TestCreateDropPointAddNewDayScheduleSuccess(t *testing.T) {
    mockRepository := mocks.NewDropPointRepositoryInterface(t)
    service := dropPointService{dropPointRepository: mockRepository}

    data := entity.DropPointsCore{
        Name:      "Test Drop Point",
        Address:   "123 Main Street",
        Latitude:  12.345678,
        Longitude: 98.765432,
        Schedule: []entity.ScheduleCore{
            {Day: "senin", OpenTime: "08:00", CloseTime: "15:00"},
            {Day: "selasa", OpenTime: "08:00", CloseTime: "15:00"},
            {Day: "kamis", OpenTime: "08:00", CloseTime: "15:00"},
            {Day: "jumat", OpenTime: "08:00", CloseTime: "15:00"},
            {Day: "sabtu", OpenTime: "08:00", CloseTime: "15:00"},
            {Day: "minggu", OpenTime: "08:00", CloseTime: "15:00"},
        },
    }

    mockRepository.On("CreateDropPoint", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
        arg := args.Get(0).(entity.DropPointsCore)
        found := false
        for _, schedule := range arg.Schedule {
            if schedule.Day == "rabu" {
                found = true
                assert.True(t, schedule.Closed, "Expected Closed to be true for Wednesday, but got false")
            }
        }
        assert.True(t, found, "Expected Schedule for Wednesday not found")
    })

    err := service.CreateDropPoint(data)

    assert.Nil(t, err)
    mockRepository.AssertExpectations(t)
}










