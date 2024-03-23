package entity

import "recything/features/drop-point/model"
// real

func CoreDropPointsToModelDropPoints(data DropPointsCore) model.DropPoints {
	return model.DropPoints{
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Schedule:  ListCoreScheduleToModelSchedule(data.Schedule),
	}
}

func CoreScheduleToModelSchedule(data ScheduleCore) model.Schedules {
	return model.Schedules{
		DropPointsID: data.DropPointsID,
		Day:          data.Day,
		OpenTime:     data.OpenTime,
		CloseTime:    data.CloseTime,
		Closed:       data.Closed,
	}
}

func ListCoreScheduleToModelSchedule(data []ScheduleCore) []model.Schedules {
	list := []model.Schedules{}
	for _, value := range data {
		result := CoreScheduleToModelSchedule(value)
		list = append(list, result)
	}
	return list
}

func ListCoreDropPointsToModelDropPoints(data []DropPointsCore) []model.DropPoints {
	list := []model.DropPoints{}
	for _, value := range data {
		result := CoreDropPointsToModelDropPoints(value)
		list = append(list, result)
	}
	return list
}

// Model To Core
func ModelDropPointsToCoreDropPoints(data model.DropPoints) DropPointsCore {
	return DropPointsCore{
		Id:        data.Id,
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Schedule:  ListModelScheduleToCoreSchedule(data.Schedule),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ModelScheduleToCoreSchedule(data model.Schedules) ScheduleCore {
	return ScheduleCore{
		Id:           data.Id,
		DropPointsID: data.DropPointsID,
		Day:          data.Day,
		OpenTime:     data.OpenTime,
		CloseTime:    data.CloseTime,
		Closed:       data.Closed,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ListModelScheduleToCoreSchedule(data []model.Schedules) []ScheduleCore {
	list := []ScheduleCore{}
	for _, value := range data {
		result := ModelScheduleToCoreSchedule(value)
		list = append(list, result)
	}
	return list
}

func ListModelDropPointsToCoreDropPoints(data []model.DropPoints) []DropPointsCore {
	list := []DropPointsCore{}
	for _, value := range data {
		result := ModelDropPointsToCoreDropPoints(value)
		list = append(list, result)
	}
	return list
}
