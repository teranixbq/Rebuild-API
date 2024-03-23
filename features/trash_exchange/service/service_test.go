package service

// Import paket-paket yang dibutuhkan untuk unit test
import (
	"errors"
	"testing"

	de "recything/features/drop-point/entity"
	tce "recything/features/trash_category/entity"
	tee "recything/features/trash_exchange/entity"
	ue "recything/features/user/entity"
	"recything/mocks"
	"recything/utils/constanta"
	"recything/utils/pagination"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTrashExchangeSuccess(t *testing.T) {
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	userRepo := new(mocks.UsersRepositoryInterface)
	dropPointRepo := new(mocks.DropPointRepositoryInterface)
	trashCategoryRepo := new(mocks.TrashCategoryRepositoryInterface)

	trashExchangeService := NewTrashExchangeService(repoData, dropPointRepo, userRepo, trashCategoryRepo)

	mockData := tee.TrashExchangeCore{
		Name:          "John Doe",
		EmailUser:     "john@example.com",
		DropPointName: "Recycling Center",
		TrashExchangeDetails: []tee.TrashExchangeDetailCore{
			{TrashType: "Plastic", Amount: 10.5},
			{TrashType: "Paper", Amount: 5.0},
		},
	}

	mockUser := ue.UsersCore{
		Id:    "user123",
		Email: "john@example.com",
		Point: 100,
	}

	mockDropPoint := de.DropPointsCore{
		Id:   "dropPoint123",
		Name: "Recycling Center",
	}

	mockTrashCategory1 := tce.TrashCategoryCore{
		ID:        "category1",
		TrashType: "Plastic",
		Point:     5,
		Unit:      "kg",
	}

	mockTrashCategory2 := tce.TrashCategoryCore{
		ID:        "category2",
		TrashType: "Paper",
		Point:     3,
		Unit:      "kg",
	}

	for _, detail := range mockData.TrashExchangeDetails {
		repoData.On("CreateTrashExchangeDetails", mock.AnythingOfType("entity.TrashExchangeDetailCore")).
			Return(detail, nil)
	}

	repoData.On("CreateTrashExchange", mock.AnythingOfType("entity.TrashExchangeCore")).Return(mockData, nil)
	userRepo.On("FindByEmail", "john@example.com").Return(mockUser, nil)
	dropPointRepo.On("GetDropPointByName", "Recycling Center").Return(mockDropPoint, nil)
	trashCategoryRepo.On("GetByType", "Plastic").Return(mockTrashCategory1, nil)
	trashCategoryRepo.On("GetByType", "Paper").Return(mockTrashCategory2, nil)
	userRepo.On("UpdateById", "user123", mock.AnythingOfType("entity.UsersCore")).Return(nil)

	result, err := trashExchangeService.CreateTrashExchange(mockData)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCreateTrashExchangeEmptyError(t *testing.T) {
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	userRepo := new(mocks.UsersRepositoryInterface)
	dropPointRepo := new(mocks.DropPointRepositoryInterface)
	trashCategoryRepo := new(mocks.TrashCategoryRepositoryInterface)

	trashExchangeService := NewTrashExchangeService(repoData, dropPointRepo, userRepo, trashCategoryRepo)

	mockData := tee.TrashExchangeCore{
		Name:          "",
		EmailUser:     "",
		DropPointName: "",
		TrashExchangeDetails: []tee.TrashExchangeDetailCore{
			{TrashType: "Plastic", Amount: 10.5},
			{TrashType: "Paper", Amount: 5.0},
		},
	}

	mockUser := ue.UsersCore{
		Id:    "user123",
		Email: "john@example.com",
		Point: 100,
	}

	mockDropPoint := de.DropPointsCore{
		Id:   "dropPoint123",
		Name: "Recycling Center",
	}

	mockTrashCategory1 := tce.TrashCategoryCore{
		ID:        "category1",
		TrashType: "Plastic",
		Point:     5,
		Unit:      "kg",
	}

	mockTrashCategory2 := tce.TrashCategoryCore{
		ID:        "category2",
		TrashType: "Paper",
		Point:     3,
		Unit:      "kg",
	}

	for _, detail := range mockData.TrashExchangeDetails {
		repoData.On("CreateTrashExchangeDetails", mock.AnythingOfType("entity.TrashExchangeDetailCore")).
			Return(detail, nil)
	}

	// Simulate an error in CreateTrashExchange
	repoData.On("CreateTrashExchange", mock.AnythingOfType("entity.TrashExchangeCore")).Return(tee.TrashExchangeCore{}, errors.New("harap lengkapi data dengan benar"))
	userRepo.On("FindByEmail", "john@example.com").Return(mockUser, nil)
	dropPointRepo.On("GetDropPointByName", "Recycling Center").Return(mockDropPoint, nil)
	trashCategoryRepo.On("GetByType", "Plastic").Return(mockTrashCategory1, nil)
	trashCategoryRepo.On("GetByType", "Paper").Return(mockTrashCategory2, nil)
	userRepo.On("UpdateById", "user123", mock.AnythingOfType("entity.UsersCore")).Return(nil)

	result, err := trashExchangeService.CreateTrashExchange(mockData)

	assert.Error(t, err)
	assert.Equal(t, "", result.Id)
	assert.EqualError(t, err, "error : harap lengkapi data dengan benar")
}

func TestCreateTrashExchangeInvalidEmailError(t *testing.T) {
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	userRepo := new(mocks.UsersRepositoryInterface)
	dropPointRepo := new(mocks.DropPointRepositoryInterface)
	trashCategoryRepo := new(mocks.TrashCategoryRepositoryInterface)

	trashExchangeService := NewTrashExchangeService(repoData, dropPointRepo, userRepo, trashCategoryRepo)

	mockData := tee.TrashExchangeCore{
		Name:          "juwan",
		EmailUser:     "juwan@gmail.com",
		DropPointName: "wasebutentuya",
		TrashExchangeDetails: []tee.TrashExchangeDetailCore{
			{TrashType: "Plastic", Amount: 10.5},
			{TrashType: "Paper", Amount: 5.0},
		},
	}

	mockDropPoint := de.DropPointsCore{
		Id:   "dropPoint123",
		Name: "Recycling Center",
	}

	mockTrashCategory1 := tce.TrashCategoryCore{
		ID:        "category1",
		TrashType: "Plastic",
		Point:     5,
		Unit:      "kg",
	}

	mockTrashCategory2 := tce.TrashCategoryCore{
		ID:        "category2",
		TrashType: "Paper",
		Point:     3,
		Unit:      "kg",
	}

	for _, detail := range mockData.TrashExchangeDetails {
		repoData.On("CreateTrashExchangeDetails", mock.AnythingOfType("entity.TrashExchangeDetailCore")).
			Return(detail, nil)
	}

	// Simulate an error in CreateTrashExchange
	repoData.On("CreateTrashExchange", mock.AnythingOfType("entity.TrashExchangeCore")).Return(tee.TrashExchangeCore{}, errors.New("harap lengkapi data dengan benar"))
	userRepo.On("FindByEmail", "juwan@gmail.com").Return(ue.UsersCore{}, errors.New("pengguna dengan email tersebut tidak ditemukan"))
	dropPointRepo.On("GetDropPointByName", "Recycling Center").Return(mockDropPoint, nil)
	trashCategoryRepo.On("GetByType", "Plastic").Return(mockTrashCategory1, nil)
	trashCategoryRepo.On("GetByType", "Paper").Return(mockTrashCategory2, nil)
	userRepo.On("UpdateById", "user123", mock.AnythingOfType("entity.UsersCore")).Return(nil)

	result, err := trashExchangeService.CreateTrashExchange(mockData)

	assert.Error(t, err)
	assert.Equal(t, "", result.Id)
	assert.EqualError(t, err, "pengguna dengan email tersebut tidak ditemukan")
}

func TestCreateTrashExchangeInvalidDropPointError(t *testing.T) {
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	userRepo := new(mocks.UsersRepositoryInterface)
	dropPointRepo := new(mocks.DropPointRepositoryInterface)
	trashCategoryRepo := new(mocks.TrashCategoryRepositoryInterface)

	trashExchangeService := NewTrashExchangeService(repoData, dropPointRepo, userRepo, trashCategoryRepo)

	mockData := tee.TrashExchangeCore{
		Name:          "juwan",
		EmailUser:     "juwan@gmail.com",
		DropPointName: "wasebutentuya",
		TrashExchangeDetails: []tee.TrashExchangeDetailCore{
			{TrashType: "Plastic", Amount: 10.5},
			{TrashType: "Paper", Amount: 5.0},
		},
	}

	mockUser := ue.UsersCore{
		Id:    "user123",
		Email: "juwan@gmail.com",
		Point: 100,
	}

	mockTrashCategory1 := tce.TrashCategoryCore{
		ID:        "category1",
		TrashType: "Plastic",
		Point:     5,
		Unit:      "kg",
	}

	mockTrashCategory2 := tce.TrashCategoryCore{
		ID:        "category2",
		TrashType: "Paper",
		Point:     3,
		Unit:      "kg",
	}

	for _, detail := range mockData.TrashExchangeDetails {
		repoData.On("CreateTrashExchangeDetails", mock.AnythingOfType("entity.TrashExchangeDetailCore")).
			Return(detail, nil)
	}

	// Simulate an error in CreateTrashExchange
	repoData.On("CreateTrashExchange", mock.AnythingOfType("entity.TrashExchangeCore")).Return(tee.TrashExchangeCore{}, errors.New("harap lengkapi data dengan benar"))
	userRepo.On("FindByEmail", "juwan@gmail.com").Return(mockUser, nil)
	dropPointRepo.On("GetDropPointByName", "wasebutentuya").Return(de.DropPointsCore{}, errors.New("nama drop point tidak ditemukan"))
	trashCategoryRepo.On("GetByType", "Plastic").Return(mockTrashCategory1, nil)
	trashCategoryRepo.On("GetByType", "Paper").Return(mockTrashCategory2, nil)
	userRepo.On("UpdateById", "user123", mock.AnythingOfType("entity.UsersCore")).Return(nil)

	result, err := trashExchangeService.CreateTrashExchange(mockData)

	assert.Error(t, err)
	assert.Equal(t, "", result.Id)
	assert.EqualError(t, err, "nama drop point tidak ditemukan")
}

func TestCreateTrashExchangeEmptyDetailError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	userRepo := new(mocks.UsersRepositoryInterface)
	dropPointRepo := new(mocks.DropPointRepositoryInterface)
	trashCategoryRepo := new(mocks.TrashCategoryRepositoryInterface)

	trashExchangeService := NewTrashExchangeService(repoData, dropPointRepo, userRepo, trashCategoryRepo)

	// Mock data with empty TrashType and Amount
	mockData := tee.TrashExchangeCore{
		Name:          "John Doe",
		EmailUser:     "john@example.com",
		DropPointName: "Recycling Center",
		TrashExchangeDetails: []tee.TrashExchangeDetailCore{
			{TrashType: "", Amount: 10.5}, // Empty TrashType
			{TrashType: "Paper", Amount: 5.0},
		},
	}

	mockUser := ue.UsersCore{
		Id:    "user123",
		Email: "john@example.com",
		Point: 100,
	}

	mockDropPoint := de.DropPointsCore{
		Id:   "dropPoint123",
		Name: "Recycling Center",
	}

	// Expectations
	userRepo.On("FindByEmail", "john@example.com").Return(mockUser, nil)
	dropPointRepo.On("GetDropPointByName", "Recycling Center").Return(mockDropPoint, nil)

	// Mock repository response for the first detail
	mockTrashCategory1 := tce.TrashCategoryCore{
		ID:        "category1",
		TrashType: "Plastic",
		Point:     5,
		Unit:      "kg",
	}
	trashCategoryRepo.On("GetByType", "Plastic").Return(mockTrashCategory1, nil)

	// Mock repository response for the second detail
	mockTrashCategory2 := tce.TrashCategoryCore{
		ID:        "category2",
		TrashType: "Paper",
		Point:     3,
		Unit:      "kg",
	}
	trashCategoryRepo.On("GetByType", "Paper").Return(mockTrashCategory2, nil)

	// Call the method to be tested
	result, err := trashExchangeService.CreateTrashExchange(mockData)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, errors.New("error : harap lengkapi data dengan benar"), err)
	assert.Equal(t, tee.TrashExchangeCore{}, result)
}

func TestCreateTrashExchangeTrashCategoryNotFoundError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	userRepo := new(mocks.UsersRepositoryInterface)
	dropPointRepo := new(mocks.DropPointRepositoryInterface)
	trashCategoryRepo := new(mocks.TrashCategoryRepositoryInterface)

	trashExchangeService := NewTrashExchangeService(repoData, dropPointRepo, userRepo, trashCategoryRepo)

	// Mock data with empty TrashType and Amount
	mockData := tee.TrashExchangeCore{
		Name:          "John Doe",
		EmailUser:     "john@example.com",
		DropPointName: "Recycling Center",
		TrashExchangeDetails: []tee.TrashExchangeDetailCore{
			{TrashType: "Paper", Amount: 5.0},
		},
	}

	mockUser := ue.UsersCore{
		Id:    "user123",
		Email: "john@example.com",
		Point: 100,
	}

	mockDropPoint := de.DropPointsCore{
		Id:   "dropPoint123",
		Name: "Recycling Center",
	}

	// Expectations
	userRepo.On("FindByEmail", "john@example.com").Return(mockUser, nil)
	dropPointRepo.On("GetDropPointByName", "Recycling Center").Return(mockDropPoint, nil)

	// Mock repository response for the first detail
	mockTrashCategory1 := tce.TrashCategoryCore{
		ID:        "category1",
		TrashType: "Plastic",
		Point:     5,
		Unit:      "kg",
	}
	trashCategoryRepo.On("GetByType", "Paper").Return(mockTrashCategory1, errors.New("kategori sampah tidak ditemukan"))

	// Call the method to be tested
	result, err := trashExchangeService.CreateTrashExchange(mockData)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, errors.New("kategori sampah tidak ditemukan"), err)
	assert.Equal(t, tee.TrashExchangeCore{}, result)
}

func TestCreateTrashExchangeUpdateUserError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	userRepo := new(mocks.UsersRepositoryInterface)
	dropPointRepo := new(mocks.DropPointRepositoryInterface)
	trashCategoryRepo := new(mocks.TrashCategoryRepositoryInterface)

	trashExchangeService := NewTrashExchangeService(repoData, dropPointRepo, userRepo, trashCategoryRepo)

	// Mock data
	mockData := tee.TrashExchangeCore{
		Name:          "John Doe",
		EmailUser:     "john@example.com",
		DropPointName: "Recycling Center",
		TrashExchangeDetails: []tee.TrashExchangeDetailCore{
			{TrashType: "Plastic", Amount: 10.5},
			{TrashType: "Paper", Amount: 5.0},
		},
	}

	mockUser := ue.UsersCore{
		Id:    "user123",
		Email: "john@example.com",
		Point: 100,
	}

	mockDropPoint := de.DropPointsCore{
		Id:   "dropPoint123",
		Name: "Recycling Center",
	}

	mockTrashCategory1 := tce.TrashCategoryCore{
		ID:        "category1",
		TrashType: "Plastic",
		Point:     5,
		Unit:      "kg",
	}

	mockTrashCategory2 := tce.TrashCategoryCore{
		ID:        "category2",
		TrashType: "Paper",
		Point:     3,
		Unit:      "kg",
	}

	for _, detail := range mockData.TrashExchangeDetails {
		repoData.On("CreateTrashExchangeDetails", mock.AnythingOfType("entity.TrashExchangeDetailCore")).
			Return(detail, nil)
	}

	repoData.On("CreateTrashExchange", mock.AnythingOfType("entity.TrashExchangeCore")).Return(mockData, nil)
	userRepo.On("FindByEmail", "john@example.com").Return(mockUser, nil)
	dropPointRepo.On("GetDropPointByName", "Recycling Center").Return(mockDropPoint, nil)
	trashCategoryRepo.On("GetByType", "Plastic").Return(mockTrashCategory1, nil)
	trashCategoryRepo.On("GetByType", "Paper").Return(mockTrashCategory2, nil)

	// Simulate an error while updating user points
	userRepo.On("UpdateById", "user123", mock.AnythingOfType("entity.UsersCore")).Return(errors.New("gagal memperbarui nilai point pengguna"))

	// Call the method to be tested
	result, err := trashExchangeService.CreateTrashExchange(mockData)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "", result.Id)
	assert.Equal(t, errors.New("gagal memperbarui nilai point pengguna"), err)
}

func TestCreateTrashExchangeError(t *testing.T) {
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	userRepo := new(mocks.UsersRepositoryInterface)
	dropPointRepo := new(mocks.DropPointRepositoryInterface)
	trashCategoryRepo := new(mocks.TrashCategoryRepositoryInterface)

	trashExchangeService := NewTrashExchangeService(repoData, dropPointRepo, userRepo, trashCategoryRepo)

	mockData := tee.TrashExchangeCore{
		Name:          "John Doe",
		EmailUser:     "john@example.com",
		DropPointName: "Recycling Center",
		TrashExchangeDetails: []tee.TrashExchangeDetailCore{
			{TrashType: "Plastic", Amount: 10.5},
			{TrashType: "Paper", Amount: 5.0},
		},
	}

	mockUser := ue.UsersCore{
		Id:    "user123",
		Email: "john@example.com",
		Point: 100,
	}

	mockDropPoint := de.DropPointsCore{
		Id:   "dropPoint123",
		Name: "Recycling Center",
	}

	mockTrashCategory1 := tce.TrashCategoryCore{
		ID:        "category1",
		TrashType: "Plastic",
		Point:     5,
		Unit:      "kg",
	}

	mockTrashCategory2 := tce.TrashCategoryCore{
		ID:        "category2",
		TrashType: "Paper",
		Point:     3,
		Unit:      "kg",
	}

	for _, detail := range mockData.TrashExchangeDetails {
		repoData.On("CreateTrashExchangeDetails", mock.AnythingOfType("entity.TrashExchangeDetailCore")).
			Return(detail, nil)
	}

	// Simulate an error in CreateTrashExchange
	repoData.On("CreateTrashExchange", mock.AnythingOfType("entity.TrashExchangeCore")).Return(tee.TrashExchangeCore{}, errors.New("harap lengkapi data dengan benar"))
	userRepo.On("FindByEmail", "john@example.com").Return(mockUser, nil)
	dropPointRepo.On("GetDropPointByName", "Recycling Center").Return(mockDropPoint, nil)
	trashCategoryRepo.On("GetByType", "Plastic").Return(mockTrashCategory1, nil)
	trashCategoryRepo.On("GetByType", "Paper").Return(mockTrashCategory2, nil)
	userRepo.On("UpdateById", "user123", mock.AnythingOfType("entity.UsersCore")).Return(nil)

	result, err := trashExchangeService.CreateTrashExchange(mockData)

	assert.Error(t, err)
	assert.Equal(t, "", result.Id)
	assert.EqualError(t, err, "gagal menyimpan data trash exchange")
}

func TestDeleteTrashExchangeByIdSuccess(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	trashExchangeId := "trashExchange123"

	// Expectations
	repoData.On("DeleteTrashExchangeById", trashExchangeId).Return(nil)

	// Call the method to be tested
	err := trashExchangeService.DeleteTrashExchangeById(trashExchangeId)

	// Assertions
	assert.NoError(t, err)
}

func TestDeleteTrashExchangeByIdEmptyIdError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	trashExchangeId := ""

	// Expectations
	repoData.On("DeleteTrashExchangeById", trashExchangeId).Return(errors.New("error"))

	// Call the method to be tested
	err := trashExchangeService.DeleteTrashExchangeById(trashExchangeId)

	// Assertions
	assert.Error(t, err)
}

func TestDeleteTrashExchangeByIdError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	trashExchangeId := "idbaru"

	// Expectations
	repoData.On("DeleteTrashExchangeById", trashExchangeId).Return(errors.New("error"))

	// Call the method to be tested
	err := trashExchangeService.DeleteTrashExchangeById(trashExchangeId)

	// Assertions
	assert.Error(t, err)
}

func TestGetAllTrashExchangeSuccess(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	page := "1"
	limit := "10"
	search := "example"

	// Mock repository response
	mockTrashExchanges := []tee.TrashExchangeCore{
		{
			Id:            "1",
			Name:          "John Doe",
			EmailUser:     "john@example.com",
			DropPointName: "Recycling Center",
			// ... other fields
		},
		// Add more mock data as needed
	}

	mockPageInfo := pagination.PageInfo{
		CurrentPage: 1,
		Limit:       10,
		LastPage:    5,
	}

	mockCount := 20

	// Expectations
	repoData.On("GetAllTrashExchange", 1, 10, search).Return(mockTrashExchanges, mockPageInfo, mockCount, nil)

	// Call the method to be tested
	trashExchanges, pageInfo, count, err := trashExchangeService.GetAllTrashExchange(page, limit, search)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, trashExchanges)
	assert.Equal(t, mockTrashExchanges, trashExchanges)
	assert.Equal(t, mockPageInfo, pageInfo)
	assert.Equal(t, mockCount, count)

	// Assert that the expected method was called with the correct arguments
	repoData.AssertExpectations(t)
}
func TestGetAllTrashExchangeError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	page := "1"
	limit := "10"
	search := "example"

	// Mock repository error response
	mockError := errors.New("repository error")

	// Expectations for the error case
	repoData.On("GetAllTrashExchange", 1, 10, search).Return(nil, pagination.PageInfo{}, 0, mockError)

	// Call the method to be tested
	trashExchanges, _, _, err := trashExchangeService.GetAllTrashExchange(page, limit, search)

	// Assertions for the error case
	assert.Error(t, err)
	assert.Nil(t, trashExchanges)
	assert.EqualError(t, err, mockError.Error())
}

func TestGetAllTrashExchangeInvalidPageParamsError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	page := "1"
	limit := "unlimited"
	search := "example"

	// Mock repository error response
	mockError := errors.New("limit harus berupa angka")

	// Expectations for the error case
	repoData.On("GetAllTrashExchange", 1, 10, search).Return(nil, pagination.PageInfo{}, 0, mockError)

	// Call the method to be tested
	trashExchanges, _, _, err := trashExchangeService.GetAllTrashExchange(page, limit, search)

	// Assertions for the error case
	assert.Error(t, err)
	assert.Nil(t, trashExchanges)
	assert.EqualError(t, err, mockError.Error())
}

func TestGetTrashExchangeByIdSuccess(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	trashExchangeID := "123"

	// Mock repository response
	mockTrashExchange := tee.TrashExchangeCore{
		Id:            "123",
		Name:          "John Doe",
		EmailUser:     "john@example.com",
		DropPointName: "Recycling Center",
		// ... other fields
	}

	// Expectations
	repoData.On("GetTrashExchangeById", trashExchangeID).Return(mockTrashExchange, nil)

	// Call the method to be tested
	result, err := trashExchangeService.GetTrashExchangeById(trashExchangeID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockTrashExchange, result)

	// Assert that the expected method was called with the correct arguments
	repoData.AssertExpectations(t)
}

func TestGetTrashExchangeByIdInvalidIDError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	invalidTrashExchangeID := ""

	// Call the method to be tested
	result, err := trashExchangeService.GetTrashExchangeById(invalidTrashExchangeID)

	// Assertions for invalid ID case
	assert.Error(t, err)
	assert.Equal(t, constanta.ERROR_ID_INVALID, err.Error())
	assert.Equal(t, tee.TrashExchangeCore{}, result)

	// Assert that the repository method was not called
	repoData.AssertNotCalled(t, "GetTrashExchangeById", mock.Anything)
}

func TestGetTrashExchangeByIdRepositoryError(t *testing.T) {
	// Create mock repository
	repoData := new(mocks.TrashExchangeRepositoryInterface)
	trashExchangeService := NewTrashExchangeService(repoData, nil, nil, nil)

	// Mock data
	trashExchangeID := "123"

	// Expectations for repository error case
	repoData.On("GetTrashExchangeById", trashExchangeID).Return(tee.TrashExchangeCore{}, errors.New("repository error"))

	// Call the method to be tested
	result, err := trashExchangeService.GetTrashExchangeById(trashExchangeID)

	// Assertions for repository error case
	assert.Error(t, err)
	assert.Equal(t, "repository error", err.Error())
	assert.Equal(t, tee.TrashExchangeCore{}, result)

	// Assert that the repository method was called with the correct arguments
	repoData.AssertExpectations(t)
}
