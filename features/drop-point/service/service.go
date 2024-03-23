package service

import (
	"errors"
	"recything/features/drop-point/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"
)

type dropPointService struct {
	dropPointRepository entity.DropPointRepositoryInterface
}

func NewDropPointService(dropPoint entity.DropPointRepositoryInterface) entity.DropPointServiceInterface {
	return &dropPointService{
		dropPointRepository: dropPoint,
	}
}

func (dps *dropPointService) CreateDropPoint(data entity.DropPointsCore) error {
	errEmpty := validation.CheckDataEmpty(data.Name, data.Address, data.Latitude, data.Longitude)
	if errEmpty != nil {
		return errEmpty
	}

	errLatLong := validation.CheckLatLong(data.Latitude, data.Longitude)
	if errLatLong != nil {
		return errLatLong
	}

	existingDays := make(map[string]bool)

	for i, schedule := range data.Schedule {
		errEmpty := validation.CheckDataEmpty(schedule.Day, schedule.OpenTime, schedule.CloseTime)
		if errEmpty != nil {
			return errEmpty
		}

		_, errDay := validation.CheckEqualData(schedule.Day, constanta.Days)
		if errDay != nil {
			return errDay
		}

		errTime := validation.ValidateTime(schedule.OpenTime, schedule.CloseTime)
		if errTime != nil {
			return errTime
		}

		for j := i + 1; j < len(data.Schedule); j++ {
			if schedule.Day == data.Schedule[j].Day {
				return errors.New("hari tidak boleh sama")
			}
		}

		existingDays[schedule.Day] = true

		dayExists := false
		for j := range data.Schedule {
			if schedule.Day == data.Schedule[j].Day {
				dayExists = true
				//data.Schedule[j].Closed = false
				break
			}
		}

		if !dayExists {
			data.Schedule = append(data.Schedule, entity.ScheduleCore{
				Day:    schedule.Day,
				Closed: false,
			})
		}
	}

	for _, day := range constanta.Days {
		if !existingDays[day] {
			data.Schedule = append(data.Schedule, entity.ScheduleCore{
				Day:    day,
				Closed: true,
			})
		}
	}

	err := dps.dropPointRepository.CreateDropPoint(data)
	if err != nil {
		return err
	}

	return nil
}

func (dps *dropPointService) GetAllDropPoint(page, limit int, search string) ([]entity.DropPointsCore, pagination.PageInfo,int, error) {
	if limit > 10 {
		return nil, pagination.PageInfo{},0, errors.New("limit tidak boleh lebih dari 10")
	}

	page, limit = validation.ValidateCountLimitAndPage(page, limit)

	dropPointCores, pageInfo, count,err := dps.dropPointRepository.GetAllDropPoint(page, limit, search)
	if err != nil {
		return nil, pagination.PageInfo{},0, err
	}

	return dropPointCores, pageInfo, count,nil
}

func (dps *dropPointService) GetDropPointById(id string) (entity.DropPointsCore, error) {
	if id == "" {
		return entity.DropPointsCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	idDropPoint, err := dps.dropPointRepository.GetDropPointById(id)
	if err != nil {
		return entity.DropPointsCore{}, err
	}

	return idDropPoint, err
}

func (dps *dropPointService) UpdateDropPointById(id string, data entity.DropPointsCore) error {
	errEmpty := validation.CheckDataEmpty(data.Name, data.Address, data.Latitude, data.Longitude)
	if errEmpty != nil {
		return errEmpty
	}

	errLatLong := validation.CheckLatLong(data.Latitude, data.Longitude)
	if errLatLong != nil {
		return errLatLong
	}

	for i, schedule := range data.Schedule {
		errEmpty := validation.CheckDataEmpty(schedule.Day, schedule.OpenTime, schedule.CloseTime)
		if errEmpty != nil {
			return errEmpty
		}

		day, errDay := validation.CheckEqualData(schedule.Day, constanta.Days)
		if errDay != nil {
			return errDay
		}
		schedule.Day = day

		errTime := validation.ValidateTime(schedule.OpenTime, schedule.CloseTime)
		if errTime != nil {
			return errTime
		}

		for j := i + 1; j < len(data.Schedule); j++ {
			if schedule.Day == data.Schedule[j].Day {
				return errors.New("hari tidak boleh sama")
			}
		}
	}

	err := dps.dropPointRepository.UpdateDropPointById(id, data)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDropPoint implements entity.DropPointServiceInterface.
func (dps *dropPointService) DeleteDropPointById(id string) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	err := dps.dropPointRepository.DeleteDropPointById(id)
	if err != nil {
		return err
	}

	return nil
}
