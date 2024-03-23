package service

import (
	"errors"
	"mime/multipart"
	"recything/features/report/entity"
	"recything/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReadAllReportSuccess(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReports := []entity.ReportCore{
		{ID: "report1", UserId: mockUserID, Description: "Report 1"},
		{ID: "report2", UserId: mockUserID, Description: "Report 2"},
	
	}

	repoData.On("ReadAllReport", mockUserID).Return(mockReports, nil)
	reports, err := reportService.ReadAllReport(mockUserID)

	assert.NoError(t, err)
	assert.Equal(t, mockReports, reports)
}

func TestReadAllReportEmptyUserID(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	reports, err := reportService.ReadAllReport("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "pengguna tidak ditemukan")
	assert.Empty(t, reports)
}

func TestReadAllReportRepositoryError(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"

	repoData.On("ReadAllReport", mockUserID).Return([]entity.ReportCore{}, errors.New("repository error"))
	reports, err := reportService.ReadAllReport(mockUserID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "gagal mendapatkan data")
	assert.Empty(t, reports)
}

func TestSelectByIdSuccess(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReports := entity.ReportCore{
		ID: "report1", UserId: mockUserID, Description: "Report 1",
	}

	repoData.On("SelectById", mockUserID).Return(mockReports, nil)
	reports, err := reportService.SelectById(mockUserID)

	assert.NoError(t, err)
	assert.Equal(t, mockReports, reports)
}

func TestSelectByIdEmptyUserID(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	reports, err := reportService.SelectById("")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "id tidak cocok")
	assert.Empty(t, reports)
}

func TestSelectByIdRepositoryError(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"

	repoData.On("SelectById", mockUserID).Return(entity.ReportCore{}, errors.New("repository error"))
	reports, err := reportService.SelectById(mockUserID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "repository error")
	assert.Empty(t, reports)
}

func TestCreateReportSuccess(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReportInput := entity.ReportCore{
		ReportType:    "pelanggaran sampah",
		ScaleType:     "skala besar",
		InsidentDate:  "2023-12-25",
		InsidentTime:  "12:45",
		Location:      "jl nias",
		AddressPoint:  "didepan rumah",
		TrashType:     "sampah kering",
		Description:   "didepan rumah pak rt ada yang buang sampah sembarangan",
		Status:        "perlu ditinjau",
	
	}

	mockImages := []*multipart.FileHeader{
		{
			Filename: "image1.jpg",
			Size:     1024,
		
		},
	
	}

	expectedCreatedReport := entity.ReportCore{
		ID:           "report123",
		UserId:       mockUserID,
		ReportType:   mockReportInput.ReportType,
		ScaleType:    mockReportInput.ScaleType,
		InsidentDate: mockReportInput.InsidentDate,
		InsidentTime: mockReportInput.InsidentTime,
		Location:     mockReportInput.Location,
		AddressPoint: mockReportInput.AddressPoint,
		TrashType:    mockReportInput.TrashType,
		Description:  mockReportInput.Description,
		Status:       mockReportInput.Status,
	
	}

	repoData.On("Insert", mock.AnythingOfType("entity.ReportCore"), mock.AnythingOfType("[]*multipart.FileHeader")).Return(expectedCreatedReport, nil)

	createdReport, err := reportService.Create(mockReportInput, mockUserID, mockImages)

	assert.NoError(t, err)
	assert.Equal(t, expectedCreatedReport, createdReport)
}

func TestCreateReportInvalidImageSize(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReportInput := entity.ReportCore{
		ReportType:    "pelanggaran sampah",
		ScaleType:     "skala besar",
		InsidentDate:  "2023-12-25",
		InsidentTime:  "12:45",
		Location:      "jl nias",
		AddressPoint:  "didepan rumah",
		TrashType:     "sampah kering",
		Description:   "didepan rumah pak rt ada yang buang sampah sembarangan",
		Status:        "perlu ditinjau",
	
	}

	mockImages := []*multipart.FileHeader{
		{
			Filename: "large_image.jpg",
			Size:     30 * 1024 * 1024,
		
		},
	
	}

	expectedError := "ukuran file tidak boleh lebih dari 20 MB"
	_, err := reportService.Create(mockReportInput, mockUserID, mockImages)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err.Error())
}

func TestCreateReportInvalidReportType(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReportInput := entity.ReportCore{
		ReportType: "invalid_report_type",
	
	}

	mockImages := []*multipart.FileHeader{
	
	}

	createdReport, err := reportService.Create(mockReportInput, mockUserID, mockImages)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error : report type tidak sesuai")
	assert.Equal(t, entity.ReportCore{}, createdReport)
}

func TestCreateReportInvalidScaleType(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReportInput := entity.ReportCore{
		ReportType: "pelanggaran sampah",
		ScaleType:  "invalid_scale_type",
	
	}

	mockImages := []*multipart.FileHeader{
	
	}

	createdReport, err := reportService.Create(mockReportInput, mockUserID, mockImages)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error : scale type tidak sesuai")
	assert.Equal(t, entity.ReportCore{}, createdReport)
}

func TestCreateReportInvalidIncidentDate(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReportInput := entity.ReportCore{
		ReportType:   "pelanggaran sampah",
		ScaleType:    "skala besar",
		InsidentDate: "invalid_date",
	
	}

	mockImages := []*multipart.FileHeader{
	
	}

	createdReport, err := reportService.Create(mockReportInput, mockUserID, mockImages)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error, tanggal harus dalam format 'yyyy-mm-dd'")
	assert.Equal(t, entity.ReportCore{}, createdReport)
}

func TestCreateReportInvalidIncidentTime(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReportInput := entity.ReportCore{
		ReportType:   "pelanggaran sampah",
		ScaleType:    "skala besar",
		InsidentDate: "2023-01-01",
		InsidentTime: "invalid_time",
	
	}

	mockImages := []*multipart.FileHeader{
	
	}

	createdReport, err := reportService.Create(mockReportInput, mockUserID, mockImages)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error, jam harus dalam format 'hh:mm'")
	assert.Equal(t, entity.ReportCore{}, createdReport)
}

func TestCreateReportImageSizeExceedsLimit(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReportInput := entity.ReportCore{
		ReportType:   "pelanggaran sampah",
		ScaleType:    "skala besar",
		InsidentDate: "2023-01-01",
		InsidentTime: "12:00",
	
	}

	mockLargeImage := &multipart.FileHeader{
		Size: 21 * 1024 * 1024,
	
	}

	mockImages := []*multipart.FileHeader{mockLargeImage}

	createdReport, err := reportService.Create(mockReportInput, mockUserID, mockImages)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ukuran file tidak boleh lebih dari 20 MB")
	assert.Equal(t, entity.ReportCore{}, createdReport)
}

func TestCreateReportRepositoryError(t *testing.T) {
	repoData := new(mocks.ReportRepositoryInterface)
	reportService := NewReportService(repoData)

	mockUserID := "user123"
	mockReportInput := entity.ReportCore{
		ReportType:    "pelanggaran sampah",
		ScaleType:     "skala besar",
		InsidentDate:  "2023-12-25",
		InsidentTime:  "12:45",
		Location:      "jl nias",
		AddressPoint:  "didepan rumah",
		TrashType:     "sampah kering",
		Description:   "didepan rumah pak rt ada yang buang sampah sembarangan",
		Status:        "perlu ditinjau",
	
	}

	mockImages := []*multipart.FileHeader{
		{
			Filename: "image1.jpg",
			Size:     1024,
		
		},
	
	}

	expectedError := errors.New("repository error message")
	repoData.On("Insert", mock.AnythingOfType("entity.ReportCore"), mock.AnythingOfType("[]*multipart.FileHeader")).Return(entity.ReportCore{}, expectedError)

	createdReport, err := reportService.Create(mockReportInput, mockUserID, mockImages)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, entity.ReportCore{}, createdReport)
}