package service

import (
	"errors"
	"mime/multipart"
	"recything/features/community/entity"
	"recything/mocks"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCommunitySuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockImage := &multipart.FileHeader{}
	mockCommunityData := entity.CommunityCore{
		Name:        "Sample",
		Description: "This is a test community",
		Location:    "Test Location",
		MaxMembers:  100,
	}

	// Set up the mock behavior for GetByName to return an error, indicating that the name is not taken
	repoData.On("GetByName", mockCommunityData.Name).Return(entity.CommunityCore{}, errors.New("not found"))

	// Set up the mock behavior for CreateCommunity to succeed
	repoData.On("CreateCommunity", mockImage, mockCommunityData).Return(nil)

	// Call method
	err := communityService.CreateCommunity(mockImage, mockCommunityData)

	// Assertions
	assert.NoError(t, err)
}

func TestCreateCommunityNameTaken(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockImage := &multipart.FileHeader{}
	mockCommunityData := entity.CommunityCore{
		Name:        "Sample Community",
		Description: "This is a test community",
		Location:    "Test Location",
		MaxMembers:  100,
	}

	// Set up the mock behavior for GetByName to return no error, indicating that the name is already taken
	repoData.On("GetByName", mockCommunityData.Name).Return(entity.CommunityCore{}, nil)

	// Call the method
	err := communityService.CreateCommunity(mockImage, mockCommunityData)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "nama community sudah digunakan")
}

func TestCreateCommunityEmptyFields(t *testing.T) {
	communityService := NewCommunityService(nil)

	// Call the method with empty fields
	err := communityService.CreateCommunity(nil, entity.CommunityCore{})

	// Assertions
	assert.Error(t, err)
}

func TestCreateCommunityRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockImage := &multipart.FileHeader{}
	mockCommunityData := entity.CommunityCore{
		Name:        "Sample",
		Description: "This is a test community",
		Location:    "Test Location",
		MaxMembers:  100,
	}

	// Set up the mock behavior for GetByName to return an error, indicating that the name is not taken
	repoData.On("GetByName", mockCommunityData.Name).Return(entity.CommunityCore{}, errors.New("not found"))

	// Set up the mock behavior for CreateCommunity to return an error
	expectedError := errors.New("failed to create community")
	repoData.On("CreateCommunity", mockImage, mockCommunityData).Return(expectedError)

	// Call the method
	err := communityService.CreateCommunity(mockImage, mockCommunityData)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
}

func TestDeleteCommunityByIdSuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	communityID := "community123"

	// Mock DeleteCommunityById to return nil
	repoData.On("DeleteCommunityById", communityID).Return(nil)

	err := communityService.DeleteCommunityById(communityID)

	assert.NoError(t, err)
}

func TestDeleteCommunityByIdInvalidID(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	communityID := ""

	err := communityService.DeleteCommunityById(communityID)

	assert.Error(t, err)
	assert.Equal(t, constanta.ERROR_ID_INVALID, err.Error())
}

func TestDeleteCommunityByIdRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	communityID := "community123"

	// Mock DeleteCommunityById to return an error
	repoData.On("DeleteCommunityById", communityID).Return(errors.New("repository error"))

	err := communityService.DeleteCommunityById(communityID)

	assert.Error(t, err)
	assert.Equal(t, "repository error", err.Error())
}

func TestGetAllCommunitySuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	page := "1"
	limit := "10"
	search := "community"

	expectedCommunityCores := []entity.CommunityCore{
		{Id: "1", Name: "Community1", Description: "Description1", Location: "Location1", MaxMembers: 100},
		{Id: "2", Name: "Community2", Description: "Description2", Location: "Location2", MaxMembers: 150},
		{Id: "3", Name: "Community3", Description: "Description3", Location: "Location3", MaxMembers: 100},
		{Id: "4", Name: "Community4", Description: "Description4", Location: "Location4", MaxMembers: 150},
		{Id: "5", Name: "Community5", Description: "Description5", Location: "Location5", MaxMembers: 100},
	}

	expectedPageInfo := pagination.PageInfo{
		CurrentPage: 1,
		Limit:       10,
		LastPage:    4,
	}

	expectedCount := 15

	repoData.On("GetAllCommunity", 1, 10, "community").Return(expectedCommunityCores, expectedPageInfo, expectedCount, nil)

	communityCores, pageInfo, count, err := communityService.GetAllCommunity(page, limit, search)

	assert.NoError(t, err)
	assert.Equal(t, expectedCommunityCores, communityCores)
	assert.Equal(t, expectedPageInfo, pageInfo)
	assert.Equal(t, expectedCount, count)
}

func TestGetAllCommunityInvalidPaginationParameters(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	page := "invalid"
	limit := "10"
	search := "community"

	communityCores, pageInfo, count, err := communityService.GetAllCommunity(page, limit, search)

	assert.Error(t, err)
	assert.Nil(t, communityCores)
	assert.Equal(t, pagination.PageInfo{}, pageInfo)
	assert.Equal(t, 0, count)
}

func TestGetAllCommunityRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	page := "1"
	limit := "10"
	search := "community"

	repoData.On("GetAllCommunity", 1, 10, "community").Return(nil, pagination.PageInfo{}, 0, errors.New("repository error"))

	communityCores, pageInfo, count, err := communityService.GetAllCommunity(page, limit, search)

	assert.Error(t, err)
	assert.Nil(t, communityCores)
	assert.Equal(t, pagination.PageInfo{}, pageInfo)
	assert.Equal(t, 0, count)
	assert.Equal(t, "repository error", err.Error())
}

func TestGetCommunityByIdSuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockID := "community123"

	// Mock data for the community returned by the repository
	mockCommunity := entity.CommunityCore{
		Id:          mockID,
		Name:        "Sample Community",
		Description: "This is a test community",
		Location:    "Test Location",
		Members:     50,
		MaxMembers:  100,
		Image:       "community_image.jpg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Set up the mock behavior
	repoData.On("GetCommunityById", mockID).Return(mockCommunity, nil)

	// Call the method
	resultCommunity, err := communityService.GetCommunityById(mockID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockCommunity, resultCommunity)
}

func TestGetCommunityByIdInvalidID(t *testing.T) {
	communityService := NewCommunityService(nil)

	invalidID := ""

	// Call the method with an invalid ID
	resultCommunity, err := communityService.GetCommunityById(invalidID)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, entity.CommunityCore{}, resultCommunity)
	assert.EqualError(t, err, constanta.ERROR_ID_INVALID)
}

func TestGetCommunityByIdRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockID := "community123"

	// Set up the mock behavior for repository error
	repoData.On("GetCommunityById", mockID).Return(entity.CommunityCore{}, errors.New("repository error"))

	// Call the method
	resultCommunity, err := communityService.GetCommunityById(mockID)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, entity.CommunityCore{}, resultCommunity)
}

func TestUpdateCommunityByIdSuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockID := "community123"
	mockImage := &multipart.FileHeader{
		Filename: "community_image.jpg",
		Size:     1024,
	}

	mockData := entity.CommunityCore{
		Name:        "Updated Community",
		Description: "Updated description",
		Location:    "Updated Location",
		Members:     100,
		MaxMembers:  150,
	}

	// Set up the mock behavior for repository
	repoData.On("GetByName", mockData.Name).Return(entity.CommunityCore{}, errors.New("not found"))
	repoData.On("UpdateCommunityById", mockID, mockImage, mockData).Return(nil)

	// Call the method
	err := communityService.UpdateCommunityById(mockID, mockImage, mockData)

	// Assertions
	assert.NoError(t, err)
}

func TestUpdateCommunityByIdEmptyFields(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockID := "community123"
	mockImage := &multipart.FileHeader{
		Filename: "community_image.jpg",
		Size:     1024,
	}

	mockData := entity.CommunityCore{} // Empty fields

	// Call the method with empty fields
	err := communityService.UpdateCommunityById(mockID, mockImage, mockData)

	// Assertions
	assert.Error(t, err)
}

func TestUpdateCommunityByIdDuplicateName(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockID := "community123"
	mockImage := &multipart.FileHeader{
		Filename: "community_image.jpg",
		Size:     1024,
	}

	mockData := entity.CommunityCore{
		Name:        "Updated Community",
		Description: "Updated description",
		Location:    "Updated Location",
		Members:     100,
		MaxMembers:  150,
	}

	// Set up the mock behavior for repository with duplicate name
	repoData.On("GetByName", mockData.Name).Return(entity.CommunityCore{}, nil)

	// Call the method
	err := communityService.UpdateCommunityById(mockID, mockImage, mockData)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "nama community sudah digunakan")
}

func TestUpdateCommunityByIdRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockID := "community123"
	mockImage := &multipart.FileHeader{
		Filename: "community_image.jpg",
		Size:     1024,
	}

	mockData := entity.CommunityCore{
		Name:        "Updated Community",
		Description: "Updated description",
		Location:    "Updated Location",
		Members:     100,
		MaxMembers:  150,
		Id: "123de",
	}

	// Set up the mock behavior for GetByName to return no error
	repoData.On("GetByName", mockData.Name).Return(entity.CommunityCore{}, errors.New("not found"))

	// Set up the mock behavior for UpdateCommunityById with repository error
	repoData.On("UpdateCommunityById", mockID, mockImage, mockData).Return(errors.New("repository error"))

	// Call the method
	err := communityService.UpdateCommunityById(mockID, mockImage, mockData)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")
}

func TestCreateEventSuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventInput := entity.CommunityEventCore{
		Title:       "Sample Event",
		Description: "This is a test event",
		Date:        "2023/12/25",
		Quota:       50,
		Location:    "Test Location",
		Status:      "berjalan",
		MapLink:     "https://map-link.com",
		FormLink:    "https://form-link.com",
	}

	mockImage := &multipart.FileHeader{
		Filename: "event_image.jpg",
		Size:     1024,
	}

	// Set up the mock behavior for repository
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)
	repoData.On("CreateEvent", mockCommunityID, mockEventInput, mockImage).Return(nil)

	// Call the method
	err := communityService.CreateEvent(mockCommunityID, mockEventInput, mockImage)

	// Assertions
	assert.NoError(t, err)
}

func TestCreateEventCommunityNotFound(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "nonexistentCommunity"
	mockEventInput := entity.CommunityEventCore{
		Title:       "Sample Event",
		Description: "This is a test event",
		Date:        "2023/12/25",
		Quota:       50,
		Location:    "Test Location",
		Status:      "berjalan",
		MapLink:     "https://map-link.com",
		FormLink:    "https://form-link.com",
	}

	// Set up the mock behavior for community not found
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, errors.New("community not found"))

	// Call the method
	err := communityService.CreateEvent(mockCommunityID, mockEventInput, nil)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "community not found")
}

func TestCreateEventEmptyFields(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventInput := entity.CommunityEventCore{} // Empty fields

	// Set up the expected behavior for GetCommunityById
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Call the method with empty fields
	err := communityService.CreateEvent(mockCommunityID, mockEventInput, nil)

	// Assertions
	assert.Error(t, err)
}

func TestCreateEventInvalidDateFormat(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventInput := entity.CommunityEventCore{
		Title:       "Sample Event",
		Description: "This is a test event",
		Date:        "2023-12-25",
		Quota:       50,
		Location:    "Test Location",
		Status:      "berjalan",
		MapLink:     "https://map-link.com",
		FormLink:    "https://form-link.com",
	}

	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Call the method with invalid date format
	err := communityService.CreateEvent(mockCommunityID, mockEventInput, nil)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "error, tanggal harus dalam format 'yyyy/mm/dd'")
}

func TestCreateEventExceedFileSizeLimit(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventInput := entity.CommunityEventCore{
		Title:       "Sample Event",
		Description: "This is a test event",
		Date:        "2023/12/25",
		Quota:       50,
		Location:    "Test Location",
		Status:      "berjalan",
		MapLink:     "https://map-link.com",
		FormLink:    "https://form-link.com",
	}

	mockImage := &multipart.FileHeader{
		Filename: "event_image.jpg",
		Size:     6 * 1024 * 1024,
	}

	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Call the method with an image exceeding the file size limit
	err := communityService.CreateEvent(mockCommunityID, mockEventInput, mockImage)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "ukuran file tidak boleh lebih dari 5 MB")
}

func TestCreateEventRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventInput := entity.CommunityEventCore{
		Title:       "Sample Event",
		Description: "This is a test event",
		Date:        "2023/12/25",
		Quota:       50,
		Location:    "Test Location",
		Status:      "berjalan",
		MapLink:     "https://map-link.com",
		FormLink:    "https://form-link.com",
	}

	mockImage := &multipart.FileHeader{
		Filename: "event_image.jpg",
		Size:     1024,
	}

	// Set up the mock behavior for repository error
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)
	repoData.On("CreateEvent", mockCommunityID, mockEventInput, mockImage).Return(errors.New("repository error"))

	// Call the method
	err := communityService.CreateEvent(mockCommunityID, mockEventInput, mockImage)

	// Assertions
	assert.Error(t, err)
}

func TestDeleteEventSuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"

	// Set up the expected behavior for DeleteEvent
	repoData.On("DeleteEvent", mockCommunityID, mockEventID).Return(nil)

	// Call the method
	err := communityService.DeleteEvent(mockCommunityID, mockEventID)

	// Assertions
	assert.NoError(t, err)
}

func TestDeleteEventEmptyEventID(t *testing.T) {
	communityService := NewCommunityService(nil) // No need for a repository mock in this case

	mockCommunityID := "community123"
	mockEmptyEventID := ""

	// Call the method with an empty event ID
	err := communityService.DeleteEvent(mockCommunityID, mockEmptyEventID)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "id event tidak ditemukan")
}

func TestDeleteEventRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"

	// Set up the expected behavior for DeleteEvent with repository error
	repoData.On("DeleteEvent", mockCommunityID, mockEventID).Return(errors.New("repository error"))

	// Call the method
	err := communityService.DeleteEvent(mockCommunityID, mockEventID)

	// Assertions
	assert.Error(t, err)
}

func TestReadAllEventSuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockStatus := "berjalan"
	mockPage := "1"
	mockLimit := "10"
	mockSearch := "Sample Search"
	mockCommunityID := "community123"

	// Set up the expected behavior for ReadAllEvent
	repoData.On("ReadAllEvent", mockStatus, 1, 10, mockSearch, mockCommunityID).Return(
		[]entity.CommunityEventCore{
			{
				Id:          "event123",
				Title:       "Sample Event",
				Description: "This is a test event",
				Location:    "Test Location",
				MapLink:     "https://maps.example.com",
				FormLink:    "https://form.example.com",
				Quota:       50,
				Date:        "2023/12/31",
				Status:      "berjalan",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
		pagination.PageInfo{Limit: 10, CurrentPage: 1, LastPage: 5},
		pagination.CountEventInfo{TotalCount: 50, CountBerjalan: 20, CountBelumBerjalan: 15, CountSelesai: 15},
		nil,
	)

	// Call the method
	data, pageInfo, count, err := communityService.ReadAllEvent(mockStatus, mockPage, mockLimit, mockSearch, mockCommunityID)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, "Sample Event", data[0].Title)
	assert.Equal(t, 10, pageInfo.Limit)
	assert.Equal(t, 1, pageInfo.CurrentPage)
	assert.Equal(t, 5, pageInfo.LastPage)
	assert.Equal(t, 50, count.TotalCount)
	assert.Equal(t, 20, count.CountBerjalan)
	assert.Equal(t, 15, count.CountBelumBerjalan)
	assert.Equal(t, 15, count.CountSelesai)
}

func TestReadAllEventInvalidStatus(t *testing.T) {
	communityService := NewCommunityService(nil) // No need for a repository mock in this case

	mockInvalidStatus := "invalid"
	mockPage := "1"
	mockLimit := "10"
	mockSearch := "Sample Search"
	mockCommunityID := "community123"

	// Call the method with an invalid status
	data, pageInfo, count, err := communityService.ReadAllEvent(mockInvalidStatus, mockPage, mockLimit, mockSearch, mockCommunityID)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.EqualError(t, err, "status tidak valid")
	assert.Equal(t, pagination.PageInfo{}, pageInfo)
	assert.Equal(t, pagination.CountEventInfo{}, count)
}

func TestReadAllEventValidationErr(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockStatus := "berjalan"
	mockPage := "1"
	mockLimit := "limit"
	mockSearch := "Sample Search"
	mockCommunityID := "community123"

	// Call the method
	data, pageInfo, count, err := communityService.ReadAllEvent(mockStatus, mockPage, mockLimit, mockSearch, mockCommunityID)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.EqualError(t, err, "limit harus berupa angka")
	assert.Equal(t, pagination.PageInfo{}, pageInfo)
	assert.Equal(t, pagination.CountEventInfo{}, count)
}

func TestReadAllEventRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockStatus := "berjalan"
	mockPage := "1"
	mockLimit := "10"
	mockSearch := "Sample Search"
	mockCommunityID := "community123"

	// Set up the expected behavior for ReadAllEvent with repository error
	repoData.On("ReadAllEvent", mockStatus, 1, 10, mockSearch, mockCommunityID).Return(
		nil, pagination.PageInfo{}, pagination.CountEventInfo{}, errors.New("repository error"),
	)

	// Call the method
	data, pageInfo, count, err := communityService.ReadAllEvent(mockStatus, mockPage, mockLimit, mockSearch, mockCommunityID)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.EqualError(t, err, "repository error")
	assert.Equal(t, pagination.PageInfo{}, pageInfo)
	assert.Equal(t, pagination.CountEventInfo{}, count)
}

func TestReadEventSuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"

	// Set up the expected behavior for ReadEvent
	repoData.On("ReadEvent", mockCommunityID, mockEventID).Return(
		entity.CommunityEventCore{
			Id:          mockEventID,
			Title:       "Sample Event",
			Description: "This is a test event",
			Location:    "Test Location",
			MapLink:     "https://maps.example.com",
			FormLink:    "https://form.example.com",
			Quota:       50,
			Date:        "2023/12/31",
			Status:      "berjalan",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		nil,
	)

	// Call the method
	eventData, err := communityService.ReadEvent(mockCommunityID, mockEventID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockEventID, eventData.Id)
	assert.Equal(t, "Sample Event", eventData.Title)
}

func TestReadEventInvalidEventID(t *testing.T) {
	communityService := NewCommunityService(nil) // No need for a repository mock in this case

	mockCommunityID := "community123"
	mockInvalidEventID := ""

	// Call the method with an invalid event ID
	eventData, err := communityService.ReadEvent(mockCommunityID, mockInvalidEventID)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, entity.CommunityEventCore{}, eventData)
	assert.EqualError(t, err, "event tidak ditemukan")
}

func TestReadEventRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"

	// Set up the expected behavior for ReadEvent with repository error
	repoData.On("ReadEvent", mockCommunityID, mockEventID).Return(
		entity.CommunityEventCore{}, errors.New("repository error"),
	)

	// Call the method
	eventData, err := communityService.ReadEvent(mockCommunityID, mockEventID)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, entity.CommunityEventCore{}, eventData)
	assert.EqualError(t, err, "repository error")
}

func TestUpdateEventSuccess(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"

	mockImage := &multipart.FileHeader{
		Filename: "event_image.jpg",
		Size:     1024,
	}

	mockEventInput := entity.CommunityEventCore{
		Title:       "Updated Event Title",
		Description: "Updated event description",
		Location:    "Updated Location",
		MapLink:     "https://updated-maps.example.com",
		FormLink:    "https://updated-form.example.com",
		Quota:       75,
		Date:        "2024/01/15",
		Status:      "selesai",
	}

	// Set up the mock behavior for GetCommunityById
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Set up the mock behavior for UpdateEvent
	repoData.On("UpdateEvent", mockCommunityID, mockEventID, mockEventInput, mockImage).Return(nil)

	// Call the method
	err := communityService.UpdateEvent(mockCommunityID, mockEventID, mockEventInput, mockImage)

	// Assertions
	assert.NoError(t, err)
}

func TestUpdateEventInvalidEventID(t *testing.T) {
	communityService := NewCommunityService(nil) // No need for a repository mock in this case

	mockCommunityID := "community123"
	mockInvalidEventID := ""

	// Call the method with an invalid event ID
	err := communityService.UpdateEvent(mockCommunityID, mockInvalidEventID, entity.CommunityEventCore{}, nil)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "event tidak ditemukan")
}

func TestUpdateEventInvalidStatus(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"

	mockImage := &multipart.FileHeader{
		Filename: "event_image.jpg",
		Size:     1024,
	}

	mockEventInput := entity.CommunityEventCore{
		Status: "invalid_status",
	}

	// Set up the mock behavior for GetCommunityById
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Call the method with an invalid status
	err := communityService.UpdateEvent(mockCommunityID, mockEventID, mockEventInput, mockImage)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "error : status input tidak valid")
}

func TestUpdateEventEmptyFields(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"
	mockImage := &multipart.FileHeader{
		Filename: "event_image.jpg",
		Size:     1024,
	}

	mockEventInput := entity.CommunityEventCore{
		Description: "Updated event description",
		Location:    "Updated Location",
		MapLink:     "https://updated-maps.example.com",
		FormLink:    "https://updated-form.example.com",
		Quota:       75,
		Date:        "2024/01/15",
		Status:      "selesai",
	}

	// Set up the mock behavior for GetCommunityById
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Call the method with empty fields
	err := communityService.UpdateEvent(mockCommunityID, mockEventID, mockEventInput, mockImage)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "error : harap lengkapi data dengan benar")
}

func TestUpdateEventLargeImageSize(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"
	mockImage := &multipart.FileHeader{
		Filename: "large_image.jpg",
		Size:     10 * 1024 * 1024,
	}

	mockEventInput := entity.CommunityEventCore{
		Title:       "Updated Event Title",
		Description: "Updated event description",
		Location:    "Updated Location",
		MapLink:     "https://updated-maps.example.com",
		FormLink:    "https://updated-form.example.com",
		Quota:       75,
		Date:        "2024/01/15",
		Status:      "selesai",
	}

	// Set up the mock behavior for GetCommunityById
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Call the method with a large image size
	err := communityService.UpdateEvent(mockCommunityID, mockEventID, mockEventInput, mockImage)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "ukuran file tidak boleh lebih dari 5 MB")
}

func TestUpdateEventInvalidDateFormat(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"
	mockImage := &multipart.FileHeader{
		Filename: "event_image.jpg",
		Size:     1024,
	}

	mockEventInput := entity.CommunityEventCore{
		Title:       "Updated Event Title",
		Description: "Updated event description",
		Location:    "Updated Location",
		MapLink:     "https://updated-maps.example.com",
		FormLink:    "https://updated-form.example.com",
		Quota:       75,
		Date:        "invalid_date_format",
		Status:      "selesai",
	}

	// Set up the mock behavior for GetCommunityById
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Call the method with an invalid date format
	err := communityService.UpdateEvent(mockCommunityID, mockEventID, mockEventInput, mockImage)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "error, tanggal harus dalam format 'yyyy/mm/dd'")
}

func TestUpdateEventRepositoryError(t *testing.T) {
	repoData := new(mocks.CommunityRepositoryInterface)
	communityService := NewCommunityService(repoData)

	mockCommunityID := "community123"
	mockEventID := "event123"
	mockImage := &multipart.FileHeader{
		Filename: "event_image.jpg",
		Size:     1024,
	}

	mockEventInput := entity.CommunityEventCore{
		Title:       "Updated Event Title",
		Description: "Updated event description",
		Location:    "Updated Location",
		MapLink:     "https://updated-maps.example.com",
		FormLink:    "https://updated-form.example.com",
		Quota:       75,
		Date:        "2024/01/15",
		Status:      "selesai",
	}

	// Set up the mock behavior for GetCommunityById
	repoData.On("GetCommunityById", mockCommunityID).Return(entity.CommunityCore{}, nil)

	// Set up the mock behavior for UpdateEvent with repository error
	repoData.On("UpdateEvent", mockCommunityID, mockEventID, mockEventInput, mockImage).Return(errors.New("repository error"))

	// Call the method
	err := communityService.UpdateEvent(mockCommunityID, mockEventID, mockEventInput, mockImage)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "repository error")
}
