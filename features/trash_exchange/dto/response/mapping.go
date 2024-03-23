package response

import "recything/features/trash_exchange/entity"

func TrashExchangeDetailCoreToTrashExchangeDetailResponse(data entity.TrashExchangeDetailCore) TrashExchangeDetailRespose {
	return TrashExchangeDetailRespose{
		TrashType:   data.TrashType,
		Amount:      data.Amount,
		Unit:        data.Unit,
		TotalPoints: data.TotalPoints,
	}
}

func ListTrashExchangeDetailCoreToTrashExchangeDetailResponse(data []entity.TrashExchangeDetailCore) []TrashExchangeDetailRespose {
	responseTrashExchangeDetail := []TrashExchangeDetailRespose{}
	for _, v := range data {
		trashExchangeDetail := TrashExchangeDetailCoreToTrashExchangeDetailResponse(v)
		responseTrashExchangeDetail = append(responseTrashExchangeDetail, trashExchangeDetail)
	}
	return responseTrashExchangeDetail
}

func TrashExchangeCoreToTrashExchangeResponse(data entity.TrashExchangeCore) TrashExchangeResponse {
	trashExchangeResponse := TrashExchangeResponse{
		Id:            data.Id,
		Name:          data.Name,
		EmailUser:     data.EmailUser,
		DropPointName: data.DropPointName,
		TotalPoint:    data.TotalPoint,
		TotalUnit:     data.TotalUnit,
		CreatedAt:     data.CreatedAt,
	}

	trashExchange := ListTrashExchangeDetailCoreToTrashExchangeDetailResponse(data.TrashExchangeDetails)
	trashExchangeResponse.TrashExchangeDetails = trashExchange
	return trashExchangeResponse
}

func ListTrashExchangeCoreToTrashExchangeResponse(data []entity.TrashExchangeCore) []TrashExchangeResponse {
	list := []TrashExchangeResponse{}
	for _, v := range data {
		result := TrashExchangeCoreToTrashExchangeResponse(v)
		list = append(list, result)
	}
	return list
}
