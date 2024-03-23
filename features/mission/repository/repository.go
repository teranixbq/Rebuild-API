package repository

import (
	"errors"
	"fmt"

	"mime/multipart"
	"recything/features/mission/entity"
	"recything/features/mission/model"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"recything/utils/storage"
	"recything/utils/validation"
	"strings"
	"time"

	"gorm.io/gorm"
)

type MissionRepository struct {
	db *gorm.DB
}

func NewMissionRepository(db *gorm.DB) entity.MissionRepositoryInterface {
	return &MissionRepository{
		db: db,
	}
}

func (mr *MissionRepository) CreateMission(input entity.Mission) error {
	data := entity.MissionCoreToMissionModel(input)
	tx := mr.db.Create(&data)
	if tx.Error != nil {
		if validation.IsDuplicateError(tx.Error) {
			return errors.New(constanta.ERROR_DATA_EXIST)
		}
		return tx.Error
	}
	return nil
}

func (mr *MissionRepository) FindAllMission(page, limit int, search, filter string) ([]entity.Mission, pagination.PageInfo, helper.CountMission, error) {

	var validMission []model.Mission
	mr.db.Find(&validMission)

	for _, v := range validMission {
		err := mr.UpdateMissionStatus(v)
		if err!=nil{
			return nil, pagination.PageInfo{}, helper.CountMission{}, err
		}
	}

	counts, _ := mr.GetCountDataMission(filter, search)
	totalCount := int(counts.TotalCount)
	data := []model.Mission{}
	offsetInt := (page - 1) * limit
	paginationQuery := mr.db.Limit(limit).Offset(offsetInt)

	if filter != "" {
		if search != "" {
			tx := paginationQuery.Where("status LIKE ? AND title LIKE ?", "%"+filter+"%", "%"+search+"%").Find(&data)
			if tx.Error != nil {
				return nil, pagination.PageInfo{}, counts, tx.Error
			}
		}

		if search == "" {
			tx := paginationQuery.Where("status LIKE ?", "%"+filter+"%").Find(&data)
			if tx.Error != nil {
				return nil, pagination.PageInfo{}, counts, tx.Error
			}
		}

	}

	if filter == "" {
		if search != "" {
			tx := paginationQuery.Where("title LIKE ?", "%"+search+"%").Find(&data)
			if tx.Error != nil {
				return nil, pagination.PageInfo{}, counts, tx.Error
			}
		}
	}

	if search == "" && filter == "" {
		tx := paginationQuery.Find(&data)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, tx.Error
		}
	}

	dataMission := entity.ListMissionModelToMissionCore(data)
	paginationInfo := pagination.CalculateData(totalCount, limit, page)
	return dataMission, paginationInfo, counts, nil
}

func (mr *MissionRepository) UpdateMissionStatus(data model.Mission) error {
	newStatus, err := helper.ChangeStatusMission(data.EndDate)
	if err != nil {
		return err
	}

	data.Status = newStatus
	err = mr.db.Model(&data).Updates(map[string]interface{}{"status": data.Status}).Error
	if err != nil {
		return err
	}

	return nil
}

func (mr *MissionRepository) FindAllMissionUser(userID string, filter string) ([]entity.MissionHistories, error) {

	dataMission := []model.Mission{}
	mr.db.Find(&dataMission)

	for _, v := range dataMission {
		err := mr.UpdateMissionStatus(v)
		if err != nil {
			return nil, err
		}
	}

	if filter != "" {
		if filter == constanta.BERJALAN {
			var missionsWithoutTasks []model.Mission

			var claimedMissionIDs []string
			mr.db.Model(&model.ClaimedMission{}).Where("user_id = ?", userID).Pluck("mission_id", &claimedMissionIDs)

			var rejectedOrPendingMissionIDs []string
			mr.db.Model(&model.UploadMissionTask{}).Where("user_id = ? AND status IN (?)", userID, []string{constanta.PERLU_TINJAUAN, constanta.DITOLAK}).Pluck("mission_id", &rejectedOrPendingMissionIDs)

			uniqueMissionIDs := append(claimedMissionIDs, rejectedOrPendingMissionIDs...)
			subQuery := mr.db.Model(&model.UploadMissionTask{}).Select("mission_id").Where("user_id = ? AND status = ?", userID, constanta.DISETUJUI)

			mr.db.Where("id IN (?) AND id NOT IN (?)", uniqueMissionIDs, subQuery).Order("created_at DESC").Find(&missionsWithoutTasks)

			histories := []entity.MissionHistories{}

			for _, v := range missionsWithoutTasks {
				upmistask := model.UploadMissionTask{}
				claimed := model.ClaimedMission{}
				mr.db.Where("mission_id = ? AND user_id = ?", v.ID, userID).First(&upmistask)
				mr.db.Where("mission_id = ? AND user_id = ?", v.ID, userID).First(&claimed)

				if upmistask.ID == "" {
					newHis := entity.MissionToMissionHistoriesCore(v, claimed, upmistask)
					newHis.TransactionID = ""
					newHis.StatusApproval = constanta.NeedProof

					//mission overdue
					endDate, _ := time.Parse("2006-01-02", newHis.EndDate)
					today := time.Now()
					if today.After(endDate) {
						mr.db.Where("id = ?", newHis.ClaimedID).Unscoped().Delete(&claimed)
					}

					histories = append(histories, newHis)
				}

				if upmistask.ID != "" {
					newHis := entity.MissionToMissionHistoriesCore(v, claimed, upmistask)

					//status approv disetujui
					if upmistask.Status == constanta.PERLU_TINJAUAN {
						newHis.StatusApproval = constanta.NeedReview
					}

					//status approv ditolak
					if upmistask.Status == constanta.DITOLAK {
						newHis.StatusApproval = upmistask.Reason
					}

					histories = append(histories, newHis)
				}
			}

			return histories, nil
		}

		if filter == constanta.SELESAI {
			var missions []model.Mission

			tx := mr.db.Joins("JOIN upload_mission_tasks ON missions.id = upload_mission_tasks.mission_id").
				Where("upload_mission_tasks.user_id = ? AND upload_mission_tasks.status = ?", userID, constanta.DISETUJUI).Order("created_at DESC").
				Find(&missions)
			if tx.Error != nil {
				return nil, tx.Error
			}

			histories := []entity.MissionHistories{}

			for _, v := range missions {
				claimed := model.ClaimedMission{}
				upmistask := model.UploadMissionTask{}
				mr.db.Where("mission_id = ? AND user_id = ?", v.ID, userID).First(&upmistask)
				mr.db.Where("mission_id = ? AND user_id = ?", v.ID, userID).First(&claimed)
				newHis := entity.MissionToMissionHistoriesCore(v, claimed, upmistask)
				histories = append(histories, newHis)
			}
			return histories, nil
		}
	}

	missions := []model.Mission{}
	tx := mr.db.Where("status = ?", constanta.ACTIVE).Order("created_at DESC").Find(&missions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	histories := entity.ListMissionModelTomissionHistoriesCore(missions)
	return histories, nil

}

func (mr *MissionRepository) GetCountDataMission(filter, search string) (helper.CountMission, error) {
	counts := helper.CountMission{}

	if filter != "" && search != "" {
		tx := mr.db.Model(&model.Mission{}).Where("status LIKE ? AND title LIKE ?", "%"+filter+"%", "%"+search+"%").Count(&counts.TotalCount)
		if tx.Error != nil {
			return counts, tx.Error
		}

		tx = mr.db.Model(&model.Mission{}).Where("status = ? AND title LIKE ?", constanta.ACTIVE, "%"+search+"%").Count(&counts.CountActive)
		if tx.Error != nil {
			return counts, tx.Error
		}

		tx = mr.db.Model(&model.Mission{}).Where("status = ? AND title LIKE ?", constanta.OVERDUE, "%"+search+"%").Count(&counts.CountExpired)
		if tx.Error != nil {
			return counts, tx.Error
		}

		if filter != constanta.ACTIVE {
			counts.CountActive = 0
		}
		if filter != constanta.OVERDUE {
			counts.CountExpired = 0
		}

		return counts, tx.Error
	}

	if search != "" {
		tx := mr.db.Model(&model.Mission{}).Where("status = ? AND title LIKE ?", constanta.OVERDUE, "%"+search+"%").Count(&counts.CountExpired)
		if tx.Error != nil {
			return counts, tx.Error
		}

		tx = mr.db.Model(&model.Mission{}).Where("status = ? AND title LIKE ?", constanta.ACTIVE, "%"+search+"%").Count(&counts.CountActive)
		if tx.Error != nil {
			return counts, tx.Error
		}
		tx = mr.db.Model(&model.Mission{}).Where("title LIKE ?", "%"+search+"%").Count(&counts.TotalCount)
		if tx.Error != nil {
			return counts, tx.Error
		}

		return counts, tx.Error
	}

	tx := mr.db.Model(&model.Mission{}).Count(&counts.TotalCount)
	if tx.Error != nil {
		return counts, tx.Error
	}

	tx = mr.db.Model(&model.Mission{}).Where("status LIKE ?", "%"+constanta.OVERDUE+"%").Count(&counts.CountExpired)
	if tx.Error != nil {
		return counts, tx.Error
	}

	tx = mr.db.Model(&model.Mission{}).Where("status LIKE ?", "%"+constanta.ACTIVE+"%").Count(&counts.CountActive)
	if tx.Error != nil {
		return counts, tx.Error
	}

	return counts, nil
}

func (mr *MissionRepository) GetCountDataMissionApproval(search string) (helper.CountMissionApproval, error) {

	counts := helper.CountMissionApproval{}
	if search != "" {
		newCounts := helper.CountMissionApproval{}
		join := fmt.Sprint("JOIN users ON upload_mission_tasks.user_id = users.id")
		query := fmt.Sprint("users.fullname LIKE ?")

		tx := mr.db.Model(&model.UploadMissionTask{}).
			Joins(join).
			Where(query, "%"+search+"%").Count(&newCounts.TotalCount)

		tx = mr.db.Model(&model.UploadMissionTask{}).
			Joins("JOIN users ON upload_mission_tasks.user_id = users.id").
			Where("users.fullname LIKE ? AND status LIKE ?", "%"+search+"%", "%"+constanta.DISETUJUI+"%").Count(&newCounts.CountApproved)

		tx = mr.db.Model(&model.UploadMissionTask{}).
			Joins("JOIN users ON upload_mission_tasks.user_id = users.id").
			Where("users.fullname LIKE ? AND status LIKE ?", "%"+search+"%", "%"+constanta.DITOLAK+"%").Count(&newCounts.CountRejected)

		tx = mr.db.Model(&model.UploadMissionTask{}).
			Joins("JOIN users ON upload_mission_tasks.user_id = users.id").
			Where("users.fullname LIKE ? AND status LIKE ?", "%"+search+"%", "%"+constanta.PERLU_TINJAUAN+"%").Count(&newCounts.CountPending)

		if tx.Error != nil {
			return counts, tx.Error
		}
		return newCounts, nil
	}

	tx := mr.db.Model(&model.UploadMissionTask{}).Count(&counts.TotalCount)
	if tx.Error != nil {
		return counts, tx.Error
	}

	err := mr.db.Model(&model.UploadMissionTask{}).Where("status LIKE ?", "%"+constanta.DISETUJUI+"%").Count(&counts.CountApproved).Error
	if err != nil {
		return counts, err
	}
	err = mr.db.Model(&model.UploadMissionTask{}).Where("status LIKE ?", "%"+constanta.DITOLAK+"%").Count(&counts.CountRejected).Error
	if err != nil {
		return counts, err
	}
	err = mr.db.Model(&model.UploadMissionTask{}).Where("status LIKE ?", "%"+constanta.PERLU_TINJAUAN+"%").Count(&counts.CountPending).Error
	if err != nil {
		return counts, err
	}

	return counts, nil
}

func (mr *MissionRepository) FindById(missionID string) (entity.Mission, error) {
	dataMission := model.Mission{}

	tx := mr.db.Where("id = ? ", missionID).First(&dataMission)
	if tx.Error != nil {
		return entity.Mission{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Mission{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := entity.MissionModelToMissionCore(dataMission)
	return dataResponse, nil
}

func (mr *MissionRepository) UpdateMission(missionID string, data entity.Mission) error {

	dataMission := entity.MissionCoreToMissionModel(data)
	getMission := model.Mission{}
	tx := mr.db.Where("id = ?", missionID).First(&getMission)
	if tx.Error != nil {
		return tx.Error
	}

	// ok := helper.FieldsEqual(getMission, data, "Title", "Description", "Point", "StartDate", "EndDate", "DescriptionStage, TitleStage")
	// if ok {
	// 	return errors.New(constanta.ERROR_INVALID_UPDATE)
	// }

	endDateValid, err := time.Parse("2006-01-02", data.EndDate)
	if err != nil {
		return err
	}
	currentTime := time.Now().Truncate(24 * time.Hour)
	if endDateValid.Before(currentTime) {
		data.Status = constanta.OVERDUE
	} else {
		data.Status = constanta.ACTIVE
	}

	tx = mr.db.Where("id = ?", missionID).Updates(&dataMission)
	if tx.Error != nil {
		if tx.Error != nil {
			if validation.IsDuplicateError(tx.Error) {
				return errors.New(constanta.ERROR_DATA_EXIST)
			}
			return tx.Error
		}
		return tx.Error
	}
	return nil
}

func (mr *MissionRepository) DeleteMission(missionID string) error {
	dataMission := model.Mission{}

	tx := mr.db.Where("id = ? ", missionID).Delete(&dataMission)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (mr *MissionRepository) GetImageURL(missionID string) (string, error) {
	mission := model.Mission{}
	err := mr.db.Where("id = ?", missionID).Take(&mission).Error
	if err != nil {
		return "", err
	}

	return mission.MissionImage, nil
}

// Claimed Mission
func (mr *MissionRepository) ClaimMission(userID string, data entity.ClaimedMission) error {
	input := entity.ClaimedCoreToClaimedMissionModel(data)

	errFind := mr.FindClaimed(userID, data.MissionID)
	if errFind == nil {
		return errors.New("error : mission sudah di klaim")
	}

	input.UserID = userID
	tx := mr.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (mr *MissionRepository) FindClaimed(userID, missionID string) error {
	dataClaimed := model.ClaimedMission{}
	tx := mr.db.Where("user_id = ? AND mission_id = ? AND claimed = 1", userID, missionID).First(&dataClaimed)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}
	return nil
}

// Upload Mission User
func (mr *MissionRepository) CreateUploadMissionTask(userID string, data entity.UploadMissionTaskCore, images []*multipart.FileHeader) (entity.UploadMissionTaskCore, error) {
	request := entity.UploadMissionTaskCoreToUploadMissionTaskModel(data)
	request.UserID = userID
	tx := mr.db.Create(&request)
	if tx.Error != nil {
		return entity.UploadMissionTaskCore{}, tx.Error
	}

	for _, image := range images {
		imageURL, uploadErr := storage.UploadProof(image)
		if uploadErr != nil {
			return entity.UploadMissionTaskCore{}, uploadErr
		}

		ImageList := entity.ImageUploadMissionCore{}
		ImageList.UploadMissionTaskID = request.ID
		ImageList.Image = imageURL

		ImageSave := entity.ImageUploadMissionCoreToImageUploadMissionModel(ImageList)

		if err := mr.db.Create(&ImageSave).Error; err != nil {
			return entity.UploadMissionTaskCore{}, err
		}

		data.Images = append(data.Images, ImageList)
	}

	dataResponse := entity.UploadMissionTaskModelToUploadMissionTaskCore(request)

	return dataResponse, nil
}

func (mr *MissionRepository) FindUploadMissionStatus(id, missionID, userID, status string) error {
	dataUpload := model.UploadMissionTask{}

	if id == "" {
		tx := mr.db.Where("user_id = ? AND mission_id = ?", userID, missionID).First(&dataUpload)
		if tx.Error != nil {
			return tx.Error
		}

		if tx.RowsAffected == 0 {
			return tx.Error
		}
	}

	if missionID == "" {
		tx := mr.db.Where("id = ? AND user_id = ? AND status = ?", id, userID, status).First(&dataUpload)
		if tx.Error != nil {
			return tx.Error
		}

		if tx.RowsAffected == 0 {
			return tx.Error
		}

	}

	return nil
}

func (mr *MissionRepository) UpdateUploadMissionTask(id string, images []*multipart.FileHeader, data entity.UploadMissionTaskCore) error {
	dataUploadMission := model.UploadMissionTask{}
	request := entity.UploadMissionTaskCoreToUploadMissionTaskModel(data)

	tx := mr.db.Where("id = ?", id).First(&dataUploadMission)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}
	request.Status = constanta.PERLU_TINJAUAN
	errUpdate := mr.db.Model(&dataUploadMission).Updates(request)
	if errUpdate.Error != nil {
		return errUpdate.Error
	}

	ImageList := []model.ImageUploadMission{}

	tx = mr.db.Unscoped().Where("upload_mission_task_id = ? ", id).Delete(&ImageList)
	if tx.Error != nil {
		return tx.Error
	}

	for _, image := range images {
		Imagedata := entity.ImageUploadMissionCore{}
		imageURL, uploadErr := storage.UploadProof(image)
		if uploadErr != nil {
			return uploadErr
		}

		Imagedata.UploadMissionTaskID = id
		Imagedata.Image = imageURL
		ImageSave := entity.ImageUploadMissionCoreToImageUploadMissionModel(Imagedata)

		if err := mr.db.Create(&ImageSave).Error; err != nil {
			return err
		}

		data.Images = append(data.Images, Imagedata)
	}

	return nil
}

func (mr *MissionRepository) FindUploadById(id string) error {
	dataUploadMission := model.UploadMissionTask{}

	tx := mr.db.Where("id = ? ", id).First(&dataUploadMission)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}
func (mr *MissionRepository) FindMissionApprovalById(UploadMissionTaskID string) (entity.UploadMissionTaskCore, error) {
	data := model.UploadMissionTask{}

	tx := mr.db.Where("id = ? ", UploadMissionTaskID).Preload("Images").First(&data)
	if tx.Error != nil {
		return entity.UploadMissionTaskCore{}, tx.Error
	}

	// if tx.RowsAffected == 0 {
	// 	return entity.UploadMissionTaskCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	// }

	result := entity.UploadMissionTaskModelToUploadMissionTaskCore(data)

	return result, nil
}

func (mr *MissionRepository) FindAllMissionApproval(page, limit int, search, filter string) ([]entity.UploadMissionTaskCore, pagination.PageInfo, helper.CountMissionApproval, error) {
	approvalMission := []model.UploadMissionTask{}
	offsetInt := pagination.Offset(page, limit)
	paginationQuery := mr.db.Limit(limit).Offset(offsetInt)
	counts, _ := mr.GetCountDataMissionApproval(search)

	var totalCount int
	if filter != "" {
		if strings.Contains(filter, constanta.PERLU_TINJAUAN) {
			totalCount = int(counts.CountPending)
		}
		if strings.Contains(filter, constanta.DITOLAK) {
			totalCount = int(counts.CountRejected)
		}
		if strings.Contains(filter, constanta.DISETUJUI) {
			totalCount = int(counts.CountApproved)
		}

		tx := paginationQuery.Where("status LIKE ?", "%"+filter+"%").Preload("Images").Order("created_at DESC").Find(&approvalMission)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, tx.Error
		}
	}

	if search != "" {
		totalCount = int(counts.TotalCount)
		tx := paginationQuery.Model(&model.UploadMissionTask{}).
			Joins("JOIN users ON upload_mission_tasks.user_id = users.id").
			Where("users.fullname LIKE ?", "%"+search+"%").Preload("Images").Order("created_at DESC").
			Find(&approvalMission)

		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, errors.New("error disini")
		}
	}

	if search == "" && filter == "" {
		totalCount = int(counts.TotalCount)
		tx := paginationQuery.Preload("Images").Order("created_at DESC").Find(&approvalMission)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, tx.Error
		}
	}

	paginationInfo := pagination.CalculateData(totalCount, limit, page)
	result := entity.ListUploadMissionTaskModelToUploadMissionTaskCore(approvalMission)
	return result, paginationInfo, counts, nil

}

func (mr *MissionRepository) UpdateStatusMissionApproval(uploadMissionTaskID, status, reason string) error {
	err := mr.db.Model(&model.UploadMissionTask{}).Where("id = ?", uploadMissionTaskID).Updates(map[string]interface{}{
		"status": status,
		"reason": reason,
	}).Error

	if err != nil {
		return err
	}
	return nil

}

func (mr *MissionRepository) FindHistoryById(userID, transactionID string) (entity.UploadMissionTaskCore, error) {
	data := model.UploadMissionTask{}
	tx := mr.db.Where("user_id = ? AND id = ?", userID, transactionID).Preload("Images").First(&data)
	if tx.Error != nil {
		return entity.UploadMissionTaskCore{}, tx.Error
	}

	result := entity.UploadMissionTaskModelToUploadMissionTaskCore(data)
	return result, nil

}

func (mr *MissionRepository) FindHistoryByIdTransaction(userID, transactionID string) (map[string]interface{}, error) {
	data := model.UploadMissionTask{}
	tx := mr.db.Where("user_id = ? AND id = ? AND status = ? ", userID, transactionID, constanta.DISETUJUI).First(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataMission, txMission := mr.FindById(data.MissionID)
	if txMission != nil {
		return nil, txMission
	}

	dataMissionResult := entity.MissionCoreToMissionModel(dataMission)
	dataResult := entity.UploadTaskModelToMissionHistoriesCore(data, dataMissionResult)
	result := entity.MissionHistoriesCoreToMapDetail(dataResult)
	return result, nil

}
