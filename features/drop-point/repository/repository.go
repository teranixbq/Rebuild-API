package repository

import (
	"errors"
	"recything/features/drop-point/entity"
	"recything/features/drop-point/model"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"

	"gorm.io/gorm"
)

type dropPointRepository struct {
	db *gorm.DB
}

func NewDropPointRepository(db *gorm.DB) entity.DropPointRepositoryInterface {
	return &dropPointRepository{
		db: db,
	}
}

func (dpr *dropPointRepository) CreateDropPoint(data entity.DropPointsCore) error {
	request := entity.CoreDropPointsToModelDropPoints(data)

	tx := dpr.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (dpr *dropPointRepository) GetAllDropPoint(page, limit int, search string) ([]entity.DropPointsCore, pagination.PageInfo, int, error) {
	dropPoint := []model.DropPoints{}

	offset := (page - 1) * limit
	query := dpr.db.Model(&model.DropPoints{}).Preload("Schedule")

	if search != "" {
		query = query.Where("name LIKE ? or address LIKE ? ", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&dropPoint)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&dropPoint)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	dataResponse := entity.ListModelDropPointsToCoreDropPoints(dropPoint)
	pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	return dataResponse, pageInfo, int(totalCount), nil

}

func (dpr *dropPointRepository) GetDropPointById(id string) (entity.DropPointsCore, error) {
	dataDropPoint := model.DropPoints{}

	tx := dpr.db.Preload("Schedule").Where("id = ?", id).First(&dataDropPoint)
	if tx.Error != nil {
		return entity.DropPointsCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.DropPointsCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := entity.ModelDropPointsToCoreDropPoints(dataDropPoint)
	return dataResponse, nil
}

func (dpr *dropPointRepository) UpdateDropPointById(id string, data entity.DropPointsCore) error {
	dropPointData := model.DropPoints{}

	tx := dpr.db.Where("id = ?", id).First(&dropPointData)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	if helper.FieldsEqual(data, dropPointData, "Name", "Address", "Latitude", "Longitude") {
	} else {
		errUpdate := dpr.db.Model(&dropPointData).Updates(entity.CoreDropPointsToModelDropPoints(data))
		if errUpdate.Error != nil {
			return errUpdate.Error
		}
	}

	for _, schedule := range data.Schedule {
		dataSchedule := model.Schedules{}

		tx := dpr.db.Where("drop_points_id = ? AND day = ?", id, schedule.Day).Take(&dataSchedule)
		if tx.Error != nil {
			return tx.Error
		}

		if helper.FieldsEqual(schedule, dataSchedule, "Day", "OpenTime", "CloseTime", "Closed") {
			continue
		}

		dataSchedule.Day = schedule.Day
		dataSchedule.OpenTime = schedule.OpenTime
		dataSchedule.CloseTime = schedule.CloseTime
		dataSchedule.Closed = schedule.Closed

		tx = dpr.db.Save(&dataSchedule)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func (dpr *dropPointRepository) DeleteDropPointById(id string) error {
	dropPointData := model.DropPoints{}

	tx := dpr.db.Unscoped().Where("id = ?", id).Delete(&dropPointData)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetDropPointByAddress implements entity.DropPointRepositoryInterface.
func (dpr *dropPointRepository) GetDropPointByName(name string) (entity.DropPointsCore, error) {
	dropPoint := model.DropPoints{}
	tx := dpr.db.Where("name = ?", name).First(&dropPoint)

	if tx.RowsAffected == 0 {
		return entity.DropPointsCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	if tx.Error != nil {
		return entity.DropPointsCore{}, tx.Error
	}

	result := entity.ModelDropPointsToCoreDropPoints(dropPoint)
	return result, nil
}