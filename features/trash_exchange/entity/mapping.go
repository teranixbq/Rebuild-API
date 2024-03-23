package entity

import (
	"recything/features/trash_exchange/model"
	"recything/utils/constanta"
	"time"
)

func TrashExchangeDetailModelToTrashExchangeDetailCore(data model.TrashExchangeDetail) TrashExchangeDetailCore {
	return TrashExchangeDetailCore{
		Id:              data.Id,
		TrashExchangeId: data.TrashExchangeId,
		TrashType:       data.Type,
		Amount:          data.Amount,
		Unit:            data.Unit,
		TotalPoints:     data.TotalPoints,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}
}

func ListTrashExchangeDetailModelToTrashExchangeDetailCore(data []model.TrashExchangeDetail) []TrashExchangeDetailCore {
	coreTrashExchange := []TrashExchangeDetailCore{}
	for _, v := range data {
		trashExchange := TrashExchangeDetailModelToTrashExchangeDetailCore(v)
		coreTrashExchange = append(coreTrashExchange, trashExchange)
	}
	return coreTrashExchange
}

func TrashExchangeModelToTrashExchangeCore(data model.TrashExchange) TrashExchangeCore {
	coreTrashExchange := TrashExchangeCore{
		Id:            data.Id,
		Name:          data.Name,
		EmailUser:     data.EmailUser,
		DropPointName: data.DropPointName,
		DropPointId:   data.DropPointId,
		TotalIncome:   data.TotalIncome,
		TotalPoint:    data.TotalPoint,
		TotalUnit:     data.TotalUnit,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
	return coreTrashExchange
}

func TrashExchangeModelToTrashExchangeCoreForGetData(data model.TrashExchange) TrashExchangeCore {
	coreTrashExchange := TrashExchangeCore{
		Id:            data.Id,
		Name:          data.Name,
		EmailUser:     data.EmailUser,
		DropPointName: data.DropPointName,
		DropPointId:   data.DropPointId,
		TotalPoint:    data.TotalPoint,
		TotalUnit:     data.TotalUnit,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
	trashExchange := ListTrashExchangeDetailModelToTrashExchangeDetailCore(data.TrashExchangeDetails)
	coreTrashExchange.TrashExchangeDetails = trashExchange
	return coreTrashExchange
}

func ListTrashExchangeModelToTrashExchangeCoreForGetData(data []model.TrashExchange) []TrashExchangeCore {
	trashExchangeCores := []TrashExchangeCore{}
	for _, v := range data {
		trashExchangeCore := TrashExchangeModelToTrashExchangeCoreForGetData(v)
		trashExchangeCores = append(trashExchangeCores, trashExchangeCore)
	}
	return trashExchangeCores
}

func TrashExchangeDetailCoreToTrashExchangeDetailModel(data TrashExchangeDetailCore) model.TrashExchangeDetail {
	return model.TrashExchangeDetail{
		Id:              data.Id,
		TrashExchangeId: data.TrashExchangeId,
		Type:            data.TrashType,
		Amount:          data.Amount,
		Unit:            data.Unit,
		TotalPoints:     data.TotalPoints,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}
}

func ListTrashExchangeDetailCoreToTrashExchangeDetailModel(data []TrashExchangeDetailCore) []model.TrashExchangeDetail {
	coreTrashExchange := []model.TrashExchangeDetail{}
	for _, v := range data {
		trashExchange := TrashExchangeDetailCoreToTrashExchangeDetailModel(v)
		coreTrashExchange = append(coreTrashExchange, trashExchange)
	}
	return coreTrashExchange
}

func TrashExchangeCoreToTrashExchangeModel(data TrashExchangeCore) model.TrashExchange {
	trashExchangeModel := model.TrashExchange{
		Id:            data.Id,
		Name:          data.Name,
		EmailUser:     data.EmailUser,
		DropPointName: data.DropPointName,
		DropPointId:   data.DropPointId,
		TotalIncome:   data.TotalIncome,
		TotalPoint:    data.TotalPoint,
		TotalUnit:     data.TotalUnit,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
	return trashExchangeModel
}

func TrashExchangeModelToMapTrash(data model.TrashExchange) map[string]interface{} {
	loc, _ := time.LoadLocation(constanta.ASIABANGKOK)
	return map[string]interface{}{
		"id_transaction":   data.Id,
		"created_at":       data.CreatedAt.Format(time.RFC3339),
		"time_transaction": data.CreatedAt.In(loc).Format("15:04:05.000"),
		"type_transaction": "drop sampah",
		"points":           data.TotalPoint,
	}
}

func TrashExchangeModelToMapTrashDetail(data model.TrashExchange) map[string]interface{} {
	loc, _ := time.LoadLocation(constanta.ASIABANGKOK)
	return map[string]interface{}{
		"id_transaction":   data.Id,
		"drop_point":       data.DropPointId,
		"created_at":       data.CreatedAt.Format(time.RFC3339),
		"time_transaction": data.CreatedAt.In(loc).Format("15:04:05.000"),
		"type_transaction": "reward penukaran sampah",
		"points":           data.TotalPoint,
		"trash_detail":     ListTrashExchangeDetailCoreToMapTrash(data.TrashExchangeDetails),
	}
}

func TrashExchangeDetailCoreToMapTrash(data model.TrashExchangeDetail) map[string]interface{} {
	return map[string]interface{}{
		"type":   data.Type,
		"amount": data.Amount,
		"unit":   data.Unit,
	}
}

func ListTrashExchangeDetailCoreToMapTrash(data []model.TrashExchangeDetail) []map[string]interface{} {
	coreTrashExchange := []map[string]interface{}{}
	for _, v := range data {
		trashExchange := TrashExchangeDetailCoreToMapTrash(v)
		coreTrashExchange = append(coreTrashExchange, trashExchange)
	}
	return coreTrashExchange
}
