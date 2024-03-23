package service

import (
	"errors"
	"mime/multipart"

	"recything/features/admin/entity"
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	"recything/mocks"

	"recything/utils/constanta"
	"recything/utils/pagination"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var dataUser = user.UsersCore{
	Id:       "1",
	Fullname: "recything",
	Email:    "recything@example.com",
	Point:    2000,
}

var dataUsers = []user.UsersCore{
	{Id: "1", Fullname: "recything", Email: "recything@example.com", Point: 2000},
	{Id: "2", Fullname: "recything2", Email: "recything2@example.com", Point: 3000},
	{Id: "3", Fullname: "recything3", Email: "recything3@example.com", Point: 4000},
}

var dataAdmin = entity.AdminCore{
	Fullname:        "John Doe",
	Email:           "john@example.com",
	Password:        "password123",
	ConfirmPassword: "password123",
	Status:          "aktif",
}

var dataAdmins = []entity.AdminCore{
	{Id: "1", Fullname: "recything", Email: "recything@example.com", Status: "aktif"},
	{Id: "2", Fullname: "recything2", Email: "recything2@example.com", Status: "aktif"},
	{Id: "3", Fullname: "recything3", Email: "recything3@example.com", Status: "aktif"},
}

var dataReport = report.ReportCore{
	ID:         "1",
	ReportType: "tumpukan sampah",
	UserId:     "user1",
	Status:     "perlu tinjauan",
}

var dataReports = []report.ReportCore{
	{ID: "1", ReportType: "tumpukan sampah", UserId: "user1", Status: "perlu tinjauan"},
	{ID: "2", ReportType: "pelanggaran sampah", UserId: "user2", Status: "perlu tinjauan"},
	{ID: "3", ReportType: "tumpukan sampah", UserId: "user3", Status: "perlu tinjauan"},
}

func TestCreateAdmin(t *testing.T) {

	t.Run("Succes Create", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "johnny@example.com",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}
		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(errors.New("not found"))
		mockRepo.On("Create", mock.AnythingOfType("*multipart.FileHeader"), requestBody).Return(requestBody, nil)

		_, err := adminService.Create(nil, requestBody)

		assert.NoError(t, err)
		assert.NotEqual(t, requestBody.Email, dataAdmin.Email)

		mockRepo.AssertExpectations(t)
	})
	t.Run("Data Empty", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "",
			Email:           "",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Status Input Invalid", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "berjalan",
		}

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		assert.NotEqualValues(t, []string{"aktif", "tidak aktif"}, requestBody.Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Invalid Fomat", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "johnexecom",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Invalid Length", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "pass",
			ConfirmPassword: "pass",
			Status:          "aktif",
		}

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Registered", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "password123456",
			ConfirmPassword: "password123456",
			Status:          "aktif",
		}

		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(errors.New("not found"))
		mockRepo.On("Create", mock.AnythingOfType("*multipart.FileHeader"), requestBody).Return(requestBody, errors.New("failed"))

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		assert.Equal(t, requestBody.Email, dataAdmin.Email)

		mockRepo.AssertExpectations(t)
	})

	// INFO THIS FUNCTION
	// - Password Not Match No testing

}

func TestGetAllAdmins(t *testing.T) {

	mockData := []entity.AdminCore{
		{Id: "1", Fullname: "John Doe", Email: "john@example.com", Status: "aktif"},
		{Id: "2", Fullname: "Jane Doe", Email: "jane@example.com", Status: "aktif"},
	}
	t.Run("Succes GetAll", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		mockRepo.On("SelectAll", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).
			Return(mockData, pagination.PageInfo{}, len(mockData), nil)

		admins, _, _, err := adminService.GetAll("1", "10", "")

		assert.NoError(t, err)
		assert.Len(t, admins, len(mockData))
		mockRepo.AssertExpectations(t)
	})
	t.Run("Wrong Limit Pagination", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)

		result, pageInfo, count, err := adminService.GetAll("", "20", "")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, pagination.PageInfo{}, pageInfo)
		assert.Equal(t, 0, count)
	})
	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		mockRepo.On("SelectAll", 1, 10, "").Return(dataAdmins, pagination.PageInfo{}, len(dataAdmins), errors.New("repository error"))

		result, pageInfo, count, err := adminService.GetAll("", "", "")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Empty(t, pageInfo)
		assert.Empty(t, count)

		mockRepo.AssertExpectations(t)
	})

}

func TestGetAdminById(t *testing.T) {
	// Mock data
	mockData := entity.AdminCore{
		Id:       "1",
		Fullname: "John Doe",
		Email:    "john@example.com",
		Status:   "aktif",
	}
	t.Run("Succes GetBYID", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		mockRepo.On("SelectById", mock.AnythingOfType("string")).Return(mockData, nil)

		admin, err := adminService.GetById("1")

		assert.NoError(t, err)
		assert.Equal(t, mockData.Id, admin.Id)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		adminID := "2"
		mockRepo.On("SelectById", mock.AnythingOfType("string")).Return(mockData, errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		admin, err := adminService.GetById(adminID)

		assert.Error(t, err)
		assert.NotEqual(t, adminID, mockData.Id)
		assert.Empty(t, admin)

		mockRepo.AssertExpectations(t)
	})

}

func TestUpdateAdminById(t *testing.T) {

	// Mock data
	updateAdmin := entity.AdminCore{
		Id:              "2",
		Fullname:        "Updated Admin",
		Email:           "updatedadmin@example.com",
		Password:        "123456789",
		ConfirmPassword: "123456789",
		Status:          "aktif",
	}

	t.Run("Succes Update", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)
		mockFileHeader := &multipart.FileHeader{
			Filename: "testfile.png",
		}

		mockRepo.On("Update", mock.AnythingOfType("*multipart.FileHeader"), mock.AnythingOfType("string"), mock.AnythingOfType("entity.AdminCore")).
			Return(nil)
		mockRepo.On("SelectById", mock.AnythingOfType("string")).Return(updateAdmin, nil)

		// Test case
		admin, _ := adminService.GetById("1")
		err := adminService.UpdateById(mockFileHeader, admin.Id, updateAdmin)

		// Assertions
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Data Empty", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "",
			Email:           "",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}

		err := adminService.UpdateById(nil, "1", requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Status Input Invalid", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "berjalan",
		}

		err := adminService.UpdateById(nil, "1", requestBody)

		assert.Error(t, err)
		assert.NotEqualValues(t, []string{"aktif", "tidak aktif"}, requestBody.Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Invalid Fomat", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "johnexecom",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}

		err := adminService.UpdateById(nil, "1", requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Invalid Length", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "pass",
			ConfirmPassword: "pass",
			Status:          "aktif",
		}

		err := adminService.UpdateById(nil, "1", requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Repository", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)
		mockFileHeader := &multipart.FileHeader{
			Filename: "testfile.txt",
		}

		mockRepo.On("Update", mock.AnythingOfType("*multipart.FileHeader"), mock.AnythingOfType("string"), mock.AnythingOfType("entity.AdminCore")).
			Return(errors.New(constanta.ERROR_INVALID_UPDATE))
		mockRepo.On("SelectById", mock.AnythingOfType("string")).Return(updateAdmin, nil)

		// Test case
		admin, _ := adminService.GetById("1")
		err := adminService.UpdateById(mockFileHeader, admin.Id, updateAdmin)

		// Assertions
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteAdmin(t *testing.T) {

	dataAdmin := entity.AdminCore{
		Id:       "1",
		Email:    "admin@example.com",
		Password: "hashedpassword",
	}

	t.Run("Succes Delete", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		adminID := "1"
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(nil)

		// Test case
		err := adminService.DeleteById(adminID)

		assert.NoError(t, err)
		assert.Equal(t, adminID, dataAdmin.Id)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)
		adminID := "2"

		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(errors.New("failed"))

		// Test case
		err := adminService.DeleteById(adminID)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

}

func TestFindByEmailANDPassword(t *testing.T) {
	dataAdmin := entity.AdminCore{
		Email:    "admin@example.com",
		Password: "hashedpassword",
	}

	t.Run("Succes Login", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)
		mockAdmin := entity.AdminCore{
			Email:    "admin@example.com",
			Password: "hashedpassword",
		}

		mockRepo.On("FindByEmailANDPassword", mockAdmin).Return(mockAdmin, nil)

		// Function Test
		admin, token, err := adminService.FindByEmailANDPassword(mockAdmin)

		assert.NoError(t, err)
		assert.Equal(t, dataAdmin, admin)
		assert.NotEmpty(t, token)
		mockRepo.AssertExpectations(t)

	})
	t.Run("Wrong Email or Password", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)
		mockAdmin := entity.AdminCore{
			Email:    "admin@example.com",
			Password: "hashedpasswordnewwrong",
		}

		mockRepo.On("FindByEmailANDPassword", mockAdmin).Return(mockAdmin, errors.New("failed"))

		// Function Test
		_, token, err := adminService.FindByEmailANDPassword(mockAdmin)

		assert.NotEqual(t, mockAdmin, dataAdmin)
		assert.EqualError(t, err, "error : email atau password salah")
		assert.Empty(t, token)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Data Empty", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Email:    "",
			Password: "",
		}

		_, _, err := adminService.FindByEmailANDPassword(requestBody)

		assert.Error(t, err)
		assert.Empty(t, requestBody.Email)
		assert.Empty(t, requestBody.Password)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Invalid Fomat", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Email:    "johnexecom",
			Password: "password123",
		}

		_, _, err := adminService.FindByEmailANDPassword(requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

}

// Manage User
func TestGetAllUsers(t *testing.T) {
	expectedUsers := []user.UsersCore{
		{Id: "1", Fullname: "User1", Email: "user1@example.com", Point: 20000},
		{Id: "2", Fullname: "User2", Email: "user2@example.com", Point: 3000},
	}

	t.Run("Succes GetAll", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)
		mockRepo.On("GetAllUsers", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUsers, pagination.PageInfo{}, len(expectedUsers), nil)

		// Panggil fungsi GetAllUsers dari AdminService
		users, _, count, err := adminService.GetAllUsers("", "", "")

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.NotEmpty(t, count)
		assert.Equal(t, len(expectedUsers), count)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Limit > 10", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)
		users, pageInfo, count, err := adminService.GetAllUsers("", "", "20")

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, pagination.PageInfo{}, pageInfo)
		assert.Equal(t, 0, count)
	})
	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		mockRepo.On("GetAllUsers", "", 1, 10).Return(dataUsers, pagination.PageInfo{}, len(dataUsers), errors.New("repository error"))

		result, pageInfo, count, err := adminService.GetAllUsers("", "", "")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Empty(t, pageInfo)
		assert.Empty(t, count)

		mockRepo.AssertExpectations(t)
	})

}

func TestGetByIdUsers(t *testing.T) {
	dataUsers := user.UsersCore{
		Id:       "1",
		Fullname: "recything",
		Email:    "recything@example.com",
		Point:    2000,
	}

	t.Run("Succes GetUserById", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)

		usersID := "1"
		mockRepo.On("GetByIdUser", usersID).Return(dataUsers, nil)

		user, err := adminService.GetByIdUsers(usersID)

		assert.NoError(t, err)
		assert.Equal(t, usersID, dataUsers.Id)
		assert.NotEmpty(t, user)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)

		usersID := "2"
		mockRepo.On("GetByIdUser", usersID).Return(dataUsers, errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		user, err := adminService.GetByIdUsers(usersID)

		assert.Error(t, err)
		assert.NotEqual(t, usersID, dataUsers.Id)
		assert.Empty(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteUsers(t *testing.T) {

	t.Run("Succes DeleteUsers", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)

		usersID := "1"
		mockRepo.On("DeleteUsers", usersID).Return(nil)

		err := adminService.DeleteUsers(usersID)

		assert.NoError(t, err)
		assert.Equal(t, usersID, dataUser.Id)

		mockRepo.AssertExpectations(t)

	})

	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)

		usersID := "2"
		mockRepo.On("DeleteUsers", usersID).Return(errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		err := adminService.DeleteUsers(usersID)

		assert.Error(t, err)
		assert.NotEqual(t, usersID, dataUser.Id)

		mockRepo.AssertExpectations(t)
	})
}

func TestGetReportById(t *testing.T) {
	t.Run("Succes GetReportId", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		reportID := "1"
		mockRepo.On("GetReportById", mock.Anything).Return(dataReport, nil)

		report, err := adminService.GetReportById(reportID)

		assert.NoError(t, err)
		assert.Equal(t, dataReport.ID, reportID)
		assert.NotNil(t, report)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		reportID := "2"
		mockRepo.On("GetReportById", mock.Anything).Return(dataReport, errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		report, err := adminService.GetReportById(reportID)

		assert.Error(t, err)
		assert.NotEqual(t, dataReport.ID, reportID)
		assert.Empty(t, report)

		mockRepo.AssertExpectations(t)
	})

}

func TestUpdateStatusReport(t *testing.T) {
	dataReportDB := report.ReportCore{
		ID:         "1",
		ReportType: "tumpukan sampah",
		UserId:     "user1",
		Status:     "ditolak",
	}

	t.Run("Succes Update", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		reportID := "1"
		status := "diterima"
		reason := ""

		mockRepo.On("UpdateStatusReport", reportID, status, reason).Return(report.ReportCore{}, nil)
		mockRepo.On("GetReportById", reportID).Return(report.ReportCore{Status: "perlu tinjauan"}, nil)

		result, _ := adminService.UpdateStatusReport(reportID, status, reason)

		// assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, reportID, dataReport.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Data Empty", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		status := ""
		reason := ""
		reportID := "1"

		_, err := adminService.UpdateStatusReport(reportID, status, reason)

		assert.Error(t, err)
		assert.Equal(t, reportID, dataReport.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Status Accept With Reason", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		status := "diterima"
		reason := "diterima"
		reportID := "1"

		_, err := adminService.UpdateStatusReport(reportID, status, reason)

		assert.Error(t, err)
		assert.Equal(t, reportID, dataReport.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Status Reject With Empty Reason", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		status := "ditolak"
		reason := ""
		reportID := "1"

		_, err := adminService.UpdateStatusReport(reportID, status, reason)

		assert.Error(t, err)
		assert.Equal(t, reportID, dataReport.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Report ID Not Found", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		reportID := "9"
		mockRepo.On("GetReportById", reportID).Return(dataReportDB, errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		_, err := adminService.GetReportById(reportID)

		assert.Error(t, err)
		assert.NotEqual(t, reportID, dataReport.ID)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		id := "1"
		status := "diterima"
		reason := ""

		mockRepo.On("GetReportById", id).Return(report.ReportCore{}, nil) 
		mockRepo.On("UpdateStatusReport", id, status, reason).Return(report.ReportCore{}, errors.New("repository error"))

		result, err := adminService.UpdateStatusReport(id, status, reason)

		assert.Error(t, err)
		assert.Equal(t,id,dataReport.ID)
		assert.Empty(t,result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllReport(t *testing.T) {

	t.Run("Succes Get all Report", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		mockRepo.On("GetAllReport", "", "", 1, 10).Return(dataReports, pagination.PageInfo{}, pagination.CountDataInfo{}, nil)

		reports, pageInfo, count, err := adminService.GetAllReport("", "", "", "")

		assert.NoError(t, err)
		assert.NotNil(t, reports)
		assert.NotNil(t, pageInfo)
		assert.NotNil(t, count)

		mockRepo.AssertExpectations(t)
	})
	t.Run("Validation Error", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		status := "perlu ditinjau"
		search := "example search"
		page := "1"
		limit := "10"

		mockRepo.On("GetAllReport", status, search, 1, 10).Return(dataReports, pagination.PageInfo{}, pagination.CountDataInfo{}, errors.New("failed"))

		reports, pageInfo, count, err := adminService.GetAllReport(status, search, page, limit)

		assert.Error(t, err)
		assert.Nil(t, reports)
		assert.NotNil(t, pageInfo)
		assert.NotNil(t, count)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Wrong Limit Pagination", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)

		result, pageInfo, count, err := adminService.GetAllReport("", "", "", "20")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Empty(t, pageInfo)
		assert.Empty(t, count)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Status Invalid Input", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		status := "perlu"
		search := "example search"
		page := "1"
		limit := "10"

		reports, pageInfo, count, err := adminService.GetAllReport(status, search, page, limit)

		assert.Error(t, err)
		assert.Nil(t, reports)
		assert.NotNil(t, pageInfo)
		assert.NotNil(t, count)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		status := "perlu ditinjau"
		search := "example search"
		page := "1"
		limit := "10"

		mockRepo.On("GetAllReport", status, search, 1, 10).Return(dataReports, pagination.PageInfo{}, pagination.CountDataInfo{}, errors.New("repository error"))

		reports, pageInfo, count, err := adminService.GetAllReport(status, search, page, limit)

		assert.Error(t, err)
		assert.Nil(t, reports)
		assert.Empty(t, pageInfo)
		assert.Empty(t, count)

		mockRepo.AssertExpectations(t)
	})
}
