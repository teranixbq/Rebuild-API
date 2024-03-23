package service

import (
	"errors"
	"log"
	"mime/multipart"
	admin "recything/features/admin/entity"
	user "recything/features/user/entity"
	"time"

	"recything/features/mission/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"recything/utils/storage"
	"recything/utils/validation"
)

type missionService struct {
	MissionRepo entity.MissionRepositoryInterface
	AdminRepo   admin.AdminRepositoryInterface
	UserRepo    user.UsersRepositoryInterface
}

func NewMissionService(missionRepo entity.MissionRepositoryInterface, adminRepo admin.AdminRepositoryInterface, userRepo user.UsersRepositoryInterface) entity.MissionServiceInterface {
	return &missionService{
		MissionRepo: missionRepo,
		AdminRepo:   adminRepo,
		UserRepo:    userRepo,
	}
}

func (ms *missionService) CreateMission(image *multipart.FileHeader, data entity.Mission) error {


	errEmpty := validation.CheckDataEmpty(data.Title, data.Description, data.StartDate, data.EndDate, data.Point, data.DescriptionStage, data.TitleStage)
	if errEmpty != nil {
		return errEmpty
	}

	err := validation.ValidateDate(data.StartDate, data.EndDate)
	if err != nil {
		return err
	}

	imageURL, errUpload := storage.UploadThumbnail(image)
	if errUpload != nil {
		return err
	}

	data.MissionImage = imageURL
	err = ms.MissionRepo.CreateMission(data)
	if err != nil {
		return err
	}

	return nil
}

func (ms *missionService) FindAllMission(page, limit, search, filter string) ([]entity.Mission, pagination.PageInfo, helper.CountMission, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, helper.CountMission{}, err
	}
	validFilter := validation.ValidateMissionStatus(filter)

	data, pagnationInfo, count, err := ms.MissionRepo.FindAllMission(pageInt, limitInt, search, validFilter)
	if err != nil {
		return nil, pagination.PageInfo{}, count, err
	}

	for i := range data {
		admin, err := ms.AdminRepo.SelectById(data[i].AdminID)
		if err != nil {
			return nil, pagination.PageInfo{}, count, err
		}

		data[i].Creator = admin.Fullname
	}

	return data, pagnationInfo, count, nil
}

func (ms *missionService) FindAllMissionUser(userID string, filter string) ([]entity.MissionHistories, error) {
	var data string
	var err error

	if filter != "" {
		data, err = validation.CheckEqualData(filter, constanta.STATUS_MISSION_USER)
		if err != nil {
			return nil, errors.New("error: filter tidak sesuai")
		}

	}

	missions, err := ms.MissionRepo.FindAllMissionUser(userID, data)
	if err != nil {
		return nil, err
	}
	return missions, nil
}

func (ms *missionService) UpdateMission(image *multipart.FileHeader, missionID string, data entity.Mission) error {

	err := validation.ValidateDateForUpdate(data.StartDate, data.EndDate)
	if err != nil {
		return err
	}
	err = validation.CheckDataEmpty(data.Title, data.Description, data.Point, data.EndDate, data.StartDate)
	if err != nil {
		return err
	}

	imageURL, err := ms.MissionRepo.GetImageURL(missionID)
	if err != nil {
		return err
	}

	if image != nil {
		newImageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			return err
		}
		data.MissionImage = newImageURL
	} else {
		data.MissionImage = imageURL
	}

	err = ms.MissionRepo.UpdateMission(missionID, data)
	if err != nil {
		return err
	}

	return nil
}

// Claimed Mission
func (ms *missionService) ClaimMission(userID string, data entity.ClaimedMission) error {
	if data.MissionID == "" {
		return errors.New(constanta.ERROR_EMPTY)
	}

	_, err := ms.MissionRepo.FindById(data.MissionID)
	if err != nil {
		return err
	}

	err = ms.MissionRepo.ClaimMission(userID, data)
	if err != nil {
		return err
	}

	return nil
}

func (ms *missionService) FindById(missionID string) (entity.Mission, error) {

	dataMission, err := ms.MissionRepo.FindById(missionID)
	if err != nil {
		return entity.Mission{}, err
	}

	admin, err := ms.AdminRepo.SelectById(dataMission.AdminID)
	if err != nil {
		return entity.Mission{}, err
	}

	dataMission.Creator = admin.Fullname
	return dataMission, nil
}

func (ms *missionService) DeleteMission(missionID string) error {

	err := ms.MissionRepo.DeleteMission(missionID)
	if err != nil {
		return err
	}

	return nil
}

// Upload Mission User
func (ms *missionService) CreateUploadMissionTask(userID string, data entity.UploadMissionTaskCore, images []*multipart.FileHeader) (entity.UploadMissionTaskCore, error) {

	err := ms.MissionRepo.FindUploadMissionStatus("", data.MissionID, userID, "")
	if err == nil {
		return entity.UploadMissionTaskCore{}, errors.New("error : sudah mengupload data")
	}

	_, err = ms.MissionRepo.FindById(data.MissionID)
	if err != nil {
		return entity.UploadMissionTaskCore{}, err
	}

	err = ms.MissionRepo.FindClaimed(userID, data.MissionID)
	if err != nil {
		return entity.UploadMissionTaskCore{}, errors.New("error : belum melakukan klaim mission")
	}

	loc, err := time.LoadLocation(constanta.ASIABANGKOK)
	if err != nil {
		return entity.UploadMissionTaskCore{}, err
	}

	data.CreatedAt = time.Now().In(loc)
	dataUpload, err := ms.MissionRepo.CreateUploadMissionTask(userID, data, images)
	if err != nil {
		return entity.UploadMissionTaskCore{}, err
	}

	return dataUpload, nil
}

func (ms *missionService) UpdateUploadMissionTask(userID, id string, images []*multipart.FileHeader, data entity.UploadMissionTaskCore) error {
	err := ms.MissionRepo.FindUploadById(id)
	if err != nil {
		return err
	}

	err = ms.MissionRepo.FindUploadMissionStatus(id, "", userID, constanta.DITOLAK)
	if err != nil {
		return errors.New("error : tidak ada data mission yang ditolak")
	}

	errUpdate := ms.MissionRepo.UpdateUploadMissionTask(id, images, data)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

// Mission Approval
func (ms *missionService) FindAllMissionApproval(page, limit, search, filter string) ([]entity.UploadMissionTaskCore, pagination.PageInfo, helper.CountMissionApproval, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, helper.CountMissionApproval{}, err
	}
	data, pagination, count, err := ms.MissionRepo.FindAllMissionApproval(pageInt, limitInt, search, filter)
	if err != nil {
		return nil, pagination, count, err

	}
	newData := []entity.UploadMissionTaskCore{}

	for _, missionTask := range data {

		missionData, _ := ms.MissionRepo.FindById(missionTask.MissionID)
		missionTask.MissionName = missionData.Title

		userData, _ := ms.UserRepo.GetById(missionTask.UserID)
		missionTask.User = userData.Fullname
		newData = append(newData, missionTask)
	}

	return newData, pagination, count, nil
}

func (ms *missionService) FindMissionApprovalById(UploadMissionTaskID string) (entity.UploadMissionTaskCore, error) {

	data, err := ms.MissionRepo.FindMissionApprovalById(UploadMissionTaskID)
	if err != nil {
		return entity.UploadMissionTaskCore{}, err
	}
	missionData, _ := ms.MissionRepo.FindById(data.MissionID)
	data.MissionName = missionData.Title

	userData, _ := ms.UserRepo.GetById(data.UserID)
	data.User = userData.Fullname

	return data, nil
}

func (ms *missionService) UpdateStatusMissionApproval(UploadMissionTaskID, status, reason string) error {

	err := validation.CheckDataEmpty(status)
	if err != nil {
		return err
	}

	approv, err := ms.FindMissionApprovalById(UploadMissionTaskID)
	if err != nil {
		return err
	}

	mission, err := ms.FindById(approv.MissionID)
	if err != nil {
		return err
	}

	user, err := ms.UserRepo.GetById(approv.UserID)
	if err != nil {
		return err	
	}

	if status == constanta.DISETUJUI {
		approv.Status = status
		approv.Reason = ""

		bonus := helper.CalculateBonus(user.Badge, mission.Point)
		totalPoint := user.Point + int(bonus)
		err := ms.UserRepo.UpdateUserPoint(approv.UserID, totalPoint)
		if err != nil {
			return err
		}
	}

	if status == constanta.DITOLAK {
		err := validation.CheckDataEmpty(reason)
		if err != nil {
			return err
		}
		approv.Status = status
		approv.Reason = reason
	}

	err = ms.MissionRepo.UpdateStatusMissionApproval(UploadMissionTaskID, approv.Status, approv.Reason)
	if err != nil {
		return err
	}

	return nil

}

func (ms *missionService) FindHistoryById(userID, transactionID string) (entity.UploadMissionTaskCore, error) {
	log.Println("")
	data, err := ms.MissionRepo.FindHistoryById(userID, transactionID)
	if err != nil {
		return data, err
	}

	missionData, _ := ms.MissionRepo.FindById(data.MissionID)
	data.MissionName = missionData.Title

	userData, _ := ms.UserRepo.GetById(data.UserID)
	data.User = userData.Fullname

	return data, nil

}
