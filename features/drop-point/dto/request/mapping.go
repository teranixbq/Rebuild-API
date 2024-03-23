package request

import (
	"recything/features/drop-point/entity"
)

func DropPointRequestToCoreDropPoint(data DropPointRequest) entity.DropPointsCore {
	return entity.DropPointsCore{
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Schedule:  ListScheduleRequestToCoreSchedule(data.Schedule),
	}
}

func ScheduleRequestToCoreSchedule(data ScheduleRequest) entity.ScheduleCore {
	return entity.ScheduleCore{
		Day:       data.Day,
		OpenTime:  data.Open_Time,
		CloseTime: data.Close_Time,
		Closed:    data.Closed,
	}
}

func ListScheduleRequestToCoreSchedule(data []ScheduleRequest) []entity.ScheduleCore {
	list := []entity.ScheduleCore{}
	for _, value := range data {
		result := ScheduleRequestToCoreSchedule(value)
		list = append(list, result)
	}
	return list
}


