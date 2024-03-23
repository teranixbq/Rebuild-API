package service

import (
	"errors"
	"mime/multipart"
	admin "recything/features/admin/entity"
	"recything/features/mission/entity"
	user "recything/features/user/entity"
	"recything/utils/constanta"

	"recything/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// var dataUploads = []entity.UploadMissionTaskCore{
// 	{ID: "123", UserID: "091", MissionID: "092", Description: "buang sampah sembarangan"},
// 	{ID: "124", UserID: "099", MissionID: "093", Description: "tumpukan sampah"},
// }

var dataUpload = entity.UploadMissionTaskCore{
	ID: "333", UserID: "123", MissionID: "092", Description: "buang sampah sembarangan", Status: "ditolak",
}

// var dataMissions = []entity.Mission{
// 	{ID: "092",
// 		Title:  "Mari Buang Sampah",
// 		Status: "perlu ditinjau"},
// 	{ID: "093",
// 		Title:  "Mari Buang Sampah",
// 		Status: "perlu ditinjau"},
// }

var dataMissi = entity.Mission{
	ID:     "092",
	Title:  "Mari Buang Sampah",
	Status: "perlu ditinjau",
}

var dataUser = user.UsersCore{
	Id:       "123",
	Fullname: "yeye",
	Email:    "yeye@gmail.com",
}

var dataClaim = entity.ClaimedMission{
	ID:        "321",
	MissionID: "092",
	UserID:    "123",
	Claimed:   true,
}

func TestCreateMission(t *testing.T) {
	t.Run("Data Empty", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		//data.Title, data.Description, data.StartDate, data.EndDate, data.Point, data.DescriptionStage, data.TitleStage)

		image := &multipart.FileHeader{
			Filename: "/home/teranix/Downloads/recythin.png",
		}
		requestBody := entity.Mission{
			Creator:          "Admin",
			Point:            2000,
			Title:            "Sampah B3",
			Description:      "Buang sampah ditempatnya",
			StartDate:        "2023-12-20",
			EndDate:          "2023-12-30",
			DescriptionStage: "",
			TitleStage:       "",
		}

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		err := missionService.CreateMission(image, requestBody)

		assert.Error(t, err)
		missionRepo.AssertExpectations(t)
	})
	t.Run("Invalid Date", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		//data.Title, data.Description, data.StartDate, data.EndDate, data.Point, data.DescriptionStage, data.TitleStage)

		image := &multipart.FileHeader{
			Filename: "/home/teranix/Downloads/recythin.png",
		}
		requestBody := entity.Mission{
			Creator:          "Admin",
			Point:            2000,
			Title:            "Sampah B3",
			Description:      "Buang sampah ditempatnya",
			StartDate:        "2023-11-20",
			EndDate:          "2023-12-30",
			DescriptionStage: "Desc Stage",
			TitleStage:       "Title Stage",
		}

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		err := missionService.CreateMission(image, requestBody)

		assert.Error(t, err)
		missionRepo.AssertExpectations(t)
	})

	t.Run("Error Storage", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		//data.Title, data.Description, data.StartDate, data.EndDate, data.Point, data.DescriptionStage, data.TitleStage)

		image := &multipart.FileHeader{
			Filename: "/home/teranix/Downloads/recythin.png",
		}
		requestBody := entity.Mission{
			Creator:          "Admin",
			Point:            2000,
			Title:            "Sampah B3",
			Description:      "Buang sampah ditempatnya",
			StartDate:        "2024-12-20",
			EndDate:          "2024-12-30",
			DescriptionStage: "Desc Stage",
			TitleStage:       "Title Stage",
		}

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		err := missionService.CreateMission(image, requestBody)

		assert.Nil(t, err)
		missionRepo.AssertExpectations(t)
	})
}

func TestFindAllMission(t *testing.T) {
	t.Run("Wrong limit", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		result, pagnation, count, err := missionService.FindAllMission("", "20", "", "")

		assert.Error(t, err)
		assert.Empty(t, pagnation)
		assert.Empty(t, count)
		assert.Empty(t, result)

		missionRepo.AssertExpectations(t)

	})
}

func TestUpdateMission(t *testing.T) {

	t.Run("Data Empty", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		requestBody := entity.Mission{
			Creator:          "Admin",
			Point:            2000,
			Title:            "Sampah B3",
			Description:      "",
			StartDate:        "2024-12-20",
			EndDate:          "2024-12-30",
			DescriptionStage: "Desc Stage",
			TitleStage:       "Title Stage",
		}
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		err := missionService.UpdateMission(nil, dataMissi.ID, requestBody)

		assert.Error(t, err)
		missionRepo.AssertExpectations(t)

	})

	t.Run("Invalid Date", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		requestBody := entity.Mission{

			StartDate: "2023-12-16",
			EndDate:   "2021-12-17",
		}
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		err := missionService.UpdateMission(nil, dataMissi.ID, requestBody)

		assert.Error(t, err)
		missionRepo.AssertExpectations(t)

	})

}

func TestDeleteMission(t *testing.T) {
	t.Run("not founc", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		missionID := "1"
		expectedError := errors.New("mission not found") // Define the expected error

		missionRepo.On("DeleteMission", missionID).Return(expectedError).Once()

		err := missionService.DeleteMission(missionID)

		assert.EqualError(t, err, expectedError.Error())

		missionRepo.AssertExpectations(t)
	})
	t.Run("success", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		missionID := "1"

		// Set the expectations for a successful deletion by ID
		missionRepo.On("DeleteMission", missionID).Return(nil).Once()

		err := missionService.DeleteMission(missionID)

		assert.NoError(t, err)
		missionRepo.AssertExpectations(t)

	})
}

func TestGetMissionByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		data := entity.Mission{
			ID:               "1",
			Title:            "title",
			Creator:          "creator",
			Status:           "status",
			AdminID:          "1",
			MissionImage:     "image",
			Point:            10,
			Description:      "desc",
			StartDate:        "2023-12-12",
			EndDate:          "2023-12-13",
			TitleStage:       "stage title",
			DescriptionStage: "description title",
		}
		missionID := "1"

		missionRepo.On("FindById", missionID).Return(data, nil).Once()

		dataadmin := admin.AdminCore{
			Id:              "1",
			Fullname:        "admin",
			Image:           "image",
			Role:            "admin",
			Email:           "admin@gmail.com",
			Password:        "12345678",
			ConfirmPassword: "12345678",
			Status:          "aktif",
		}
		adminRepo.On("SelectById", data.AdminID).Return(dataadmin, nil).Once()

		result, err := missionService.FindById(missionID)

		assert.NoError(t, err)
		assert.Equal(t, data.ID, result.ID)
		assert.Equal(t, dataadmin.Fullname, result.Creator)
		missionRepo.AssertExpectations(t)
		adminRepo.AssertExpectations(t)
	})

	t.Run("admin not found", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		data := entity.Mission{
			ID:               "1",
			Title:            "title",
			Creator:          "creator",
			Status:           "status",
			AdminID:          "1",
			MissionImage:     "image",
			Point:            10,
			Description:      "desc",
			StartDate:        "2023-12-12",
			EndDate:          "2023-12-13",
			TitleStage:       "stage title",
			DescriptionStage: "description title",
		}
		missionID := "1"
		missionRepo.On("FindById", missionID).Return(data, errors.New("data tidak ditemukan")).Once()
		result, err := missionService.FindById(missionID)

		adminRepo.On("SelectById", data.AdminID).Return(admin.AdminCore{}, err).Once()
		adminRepo.SelectById(data.AdminID)
		assert.Error(t, err)
		assert.NotEqual(t, data.ID, result.ID)

		missionRepo.AssertExpectations(t)
		adminRepo.AssertExpectations(t)
	})
	t.Run("mission not found", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		data := entity.Mission{
			ID:               "1",
			Title:            "title",
			Creator:          "creator",
			Status:           "status",
			AdminID:          "1",
			MissionImage:     "image",
			Point:            10,
			Description:      "desc",
			StartDate:        "2023-12-12",
			EndDate:          "2023-12-13",
			TitleStage:       "stage title",
			DescriptionStage: "description title",
		}
		missionID := "2"

		dataadmin := admin.AdminCore{
			Id:              "1",
			Fullname:        "admin",
			Image:           "image",
			Role:            "admin",
			Email:           "admin@gmail.com",
			Password:        "12345678",
			ConfirmPassword: "12345678",
			Status:          "aktif",
		}
		missionRepo.On("FindById", missionID).Return(data, errors.New("data tidak ditemukan")).Once()
		result, err := missionService.FindById(missionID)

		adminRepo.On("SelectById", data.AdminID).Return(dataadmin, err).Once()
		adminRepo.SelectById(data.AdminID)
		assert.Error(t, err)
		assert.NotEqual(t, data.ID, result.ID)

		missionRepo.AssertExpectations(t)
		adminRepo.AssertExpectations(t)
	})
}

func TestFindMissionApprovalById(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		uploadMissionID := "1"
		missionUpproval := entity.UploadMissionTaskCore{
			ID:          "1",
			UserID:      "1",
			User:        "YAYA",
			MissionID:   "1",
			MissionName: "mission name",
			Description: "description",
			Reason:      "ini reason",
			Status:      "perlu tinjauaan",
		}

		missionRepo.On("FindMissionApprovalById", uploadMissionID).Return(missionUpproval, nil).Once()
		missionRepo.On("FindById", missionUpproval.MissionID).Return(entity.Mission{
			ID:    missionUpproval.MissionID,
			Title: "Mission Title",
		}, nil).Once()

		userRepo.On("GetById", missionUpproval.UserID).Return(user.UsersCore{
			Id:       "1",
			Fullname: "User Fullname",
		}, nil).Once()

		result, err := missionService.FindMissionApprovalById(uploadMissionID)

		assert.NoError(t, err)
		assert.Equal(t, missionUpproval.ID, result.ID)
		assert.Equal(t, "Mission Title", result.MissionName)
		assert.Equal(t, "User Fullname", result.User)
		missionRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})
	t.Run("not found", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		uploadMissionID := "2"
		missionUpproval := entity.UploadMissionTaskCore{
			ID:          "1",
			UserID:      "1",
			User:        "YAYA",
			MissionID:   "1",
			MissionName: "mission name",
			Description: "description",
			Reason:      "ini reason",
			Status:      "perlu tinjauaan",
		}

		missionRepo.On("FindMissionApprovalById", uploadMissionID).Return(missionUpproval, nil)
		missionRepo.On("FindById", missionUpproval.MissionID).Return(entity.Mission{
			ID:    missionUpproval.MissionID,
			Title: "Mission Title",
		}, nil)

		userRepo.On("GetById", missionUpproval.UserID).Return(user.UsersCore{
			Id:       "2",
			Fullname: "User Fullname",
		}, nil)

		result, err := missionService.FindMissionApprovalById(uploadMissionID)

		assert.NoError(t, err)
		assert.NotEqual(t, uploadMissionID, result.ID)
		missionRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)

	})
}

func TestFindHistoryById(t *testing.T) {
	t.Run("Succes Find History By ID", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		userID := "123"
		transactionID := "333"

		missionRepo.On("FindHistoryById", userID, transactionID).Return(dataUpload, nil)
		missionRepo.On("FindById", mock.AnythingOfType("string")).Return(dataMissi, nil)
		userRepo.On("GetById", mock.AnythingOfType("string")).Return(dataUser, nil)

		resultUpload, err := missionService.FindHistoryById(userID, transactionID)

		assert.NoError(t, err)
		assert.Equal(t, userID, dataUser.Id)
		assert.Equal(t, transactionID, dataUpload.ID)
		assert.Equal(t, resultUpload.MissionID, dataMissi.ID)

		missionRepo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		userID := "5"
		transactionID := "09"

		missionRepo.On("FindHistoryById", userID, transactionID).Return(entity.UploadMissionTaskCore{}, errors.New(constanta.ERROR))

		result, err := missionService.FindHistoryById(userID, transactionID)

		assert.Error(t, err)
		assert.NotEqual(t, userID, dataUser.Id)
		assert.NotEqual(t, transactionID, dataUpload.ID)
		assert.Empty(t, result)
		missionRepo.AssertExpectations(t)
	})
}

func TestUpdateUploadMissionTask(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		var image []*multipart.FileHeader

		userID := "123"
		transactionID := "333"

		missionRepo.On("FindUploadById", transactionID).Return(nil)
		missionRepo.On("FindUploadMissionStatus", transactionID, mock.AnythingOfType("string"), userID, mock.AnythingOfType("string")).Return(nil)
		missionRepo.On("UpdateUploadMissionTask", transactionID, image, dataUpload).Return(nil)

		err := missionService.UpdateUploadMissionTask(userID, transactionID, nil, dataUpload)
		assert.NoError(t, err)
		missionRepo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		var image []*multipart.FileHeader

		userID := "1231"
		transactionID := "3332"

		missionRepo.On("FindUploadById", transactionID).Return(nil)
		missionRepo.On("FindUploadMissionStatus", transactionID, mock.AnythingOfType("string"), userID, mock.AnythingOfType("string")).Return(nil)
		missionRepo.On("UpdateUploadMissionTask", transactionID, image, dataUpload).Return(errors.New(constanta.ERROR))

		err := missionService.UpdateUploadMissionTask(userID, transactionID, nil, dataUpload)
		assert.Error(t, err)
		assert.NotEqual(t, userID, dataUser.Id)
		assert.NotEqual(t, transactionID, dataUpload.ID)
		missionRepo.AssertExpectations(t)
	})

	t.Run("Error Repo", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		userID := "123"
		transactionID := "333"

		missionRepo.On("FindUploadById", transactionID).Return(errors.New(constanta.ERROR))
		err := missionService.UpdateUploadMissionTask(userID, transactionID, nil, dataUpload)
		assert.Error(t, err)
		missionRepo.AssertExpectations(t)
	})
}

func TestClaimMission(t *testing.T) {

	t.Run("succes", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		missionID := "092"
		userID := "123"
		missionRepo.On("FindById", missionID).Return(dataMissi, nil)
		missionRepo.On("ClaimMission", userID, dataClaim).Return(nil)

		err := missionService.ClaimMission(userID, dataClaim)
		assert.NoError(t, err)
		missionRepo.AssertExpectations(t)
	})

	t.Run("mission id empty", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		userID := "123"

		request := entity.ClaimedMission{
			MissionID: "",
			UserID:    "",
		}

		err := missionService.ClaimMission(userID, request)

		assert.Error(t, err)
		assert.NotEqual(t, dataClaim.MissionID, request.MissionID)

		missionRepo.AssertExpectations(t)

	})

	t.Run("Data Not Found", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		missionID := "092ss"
		missionRepo.On("FindById", missionID).Return(entity.Mission{}, errors.New(constanta.ERROR))

		_, err := missionService.FindById(missionID)

		assert.Error(t, err)
		assert.NotEqual(t, missionID, dataMissi.ID)

		missionRepo.AssertExpectations(t)
	})

}

func TestFindAllMissionUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		userID := "123"
		filter := ""
		expectedMissions := []entity.MissionHistories{
			{
				MissionID: "092",
			},
			{
				MissionID: "093",
			},
		}

		missionRepo.On("FindAllMissionUser", userID, "").Return(expectedMissions, nil)
		missions, err := missionService.FindAllMissionUser(userID, filter)
		assert.NoError(t, err)
		assert.Equal(t, expectedMissions, missions)
		missionRepo.AssertExpectations(t)

	})
	t.Run("filter not empty", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		filter := "berjalan"
		userID := "123"

		expectedFilteredMissions := []entity.MissionHistories{
			{
				MissionID:     "01",
				StatusMission: "aktive",
			},
			{
				MissionID:     "02",
				StatusMission: "aktive",
			},
		}

		filteredMissions, err := missionService.FindAllMissionUser(userID, "invalid filter")

		assert.Error(t, err)
		assert.NotEqual(t, expectedFilteredMissions, filteredMissions)
		assert.NotEqual(t, filter, "invalid filter")
		missionRepo.AssertExpectations(t)

	})

	t.Run("Error Repository", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		userID := "123"
		filter := ""
		expectedMissions := []entity.MissionHistories{
			{
				MissionID: "092",
			},
			{
				MissionID: "093",
			},
		}

		missionRepo.On("FindAllMissionUser", userID, "").Return(expectedMissions, errors.New(constanta.ERROR))
		missions, err := missionService.FindAllMissionUser(userID, filter)
		assert.Error(t, err)
		assert.NotEqual(t, expectedMissions, missions)
		missionRepo.AssertExpectations(t)

	})

}

func TestUpdateStatusMissionApproval(t *testing.T) {

	t.Run("StatusEmpty", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		errStatusEmpty := errors.New("error : harap lengkapi data dengan benar")

		err := missionService.UpdateStatusMissionApproval("UploadMissionTaskID", "", "reason")
		assert.EqualError(t, err, errStatusEmpty.Error())
	})

	t.Run("FindMissionApprovalError", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		missionRepo.On("FindMissionApprovalById", "UploadMissionTaskID").Return(entity.UploadMissionTaskCore{}, errors.New("find error"))

		err := missionService.UpdateStatusMissionApproval("UploadMissionTaskID", "status", "reason")
		assert.Error(t, err)
	})
}

func TestFindAllMissionApproval(t *testing.T) {

	t.Run("Wrong Limit", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		result, pagination, count, err := missionService.FindAllMissionApproval("", "50000", "", "")

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Empty(t, pagination)
		assert.Empty(t, count)
		missionRepo.AssertExpectations(t)
	})

}
