package response

import "recything/features/drop-point/entity"

func CoreDropPointToDropPointResponse(data entity.DropPointsCore) DropPointResponse {
	return DropPointResponse{
		Id:        data.Id,
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Schedule:  ListCoreScheduleToScheduleRequest(data.Schedule),
	}
}

func CoreScheduleToScheduleResponse(data entity.ScheduleCore) ScheduleResponse {
	return ScheduleResponse{
		Day:        data.Day,
		Open_Time:  data.OpenTime,
		Close_Time: data.CloseTime,
		Closed:     data.Closed,
	}
}

func ListCoreScheduleToScheduleRequest(data []entity.ScheduleCore) []ScheduleResponse {
	list := []ScheduleResponse{}
	for _, value := range data {
		result := CoreScheduleToScheduleResponse(value)
		list = append(list, result)
	}
	return list
}

func ListCoreDropPointToDropPointResponse(data []entity.DropPointsCore) []DropPointResponse {
	list := []DropPointResponse{}
	for _, value := range data {
		result := CoreDropPointToDropPointResponse(value)
		list = append(list, result)
	}
	return list
}

// Detail

// func CoreDropPointToDropPointDetailResponse(data entity.DropPointsCore) DropPointDetailResponse {
// 	return DropPointDetailResponse{
// 		Id:        data.Id,
// 		Name:      data.Name,
// 		Address:   data.Address,
// 		Latitude:  data.Latitude,
// 		Longitude: data.Longitude,
// 		Schedule:  ListCoreScheduleToScheduleDetailRequest(data.Schedule),
// 	}
// }

// func CoreScheduleToScheduleDetailResponse(data entity.ScheduleCore) ScheduleDetailResponse {
// 	return ScheduleDetailResponse{
// 		Day:        data.Day,
// 		Open_Time:  data.OpenTime,
// 		Close_Time: data.CloseTime,
// 		Closed:     data.Closed,
// 	}
// }

// func ListCoreScheduleToScheduleDetailRequest(data []entity.ScheduleCore) []ScheduleDetailResponse {
// 	list := []ScheduleDetailResponse{}
// 	for _, value := range data {
// 		result := CoreScheduleToScheduleDetailResponse(value)
// 		list = append(list, result)
// 	}
// 	return list
// }

// func ListCoreDropPointToDropPointDetailResponse(data []entity.DropPointsCore) []DropPointDetailResponse {
// 	list := []DropPointDetailResponse{}
// 	for _, value := range data {
// 		result := CoreDropPointToDropPointDetailResponse(value)
// 		list = append(list, result)
// 	}
// 	return list
// }
