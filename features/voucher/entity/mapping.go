package entity

import (
	"recything/features/voucher/model"
	"time"
)

func CoreVoucherToModelVoucher(data VoucherCore) model.Voucher {
	return model.Voucher{
		Image:       data.Image,
		RewardName:  data.RewardName,
		Point:       data.Point,
		Description: data.Description,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
	}
}

func ListCoreVoucherToModelVoucher(data []VoucherCore) []model.Voucher {
	list := []model.Voucher{}
	for _, v := range data {
		result := CoreVoucherToModelVoucher(v)
		list = append(list, result)
	}
	return list
}

func ModelVoucherToCoreVoucher(data model.Voucher) VoucherCore {
	return VoucherCore{
		Id:          data.Id,
		Image:       data.Image,
		RewardName:  data.RewardName,
		Point:       data.Point,
		Description: data.Description,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ListModelVoucherToCoreVoucher(data []model.Voucher) []VoucherCore {
	list := []VoucherCore{}
	for _, v := range data {
		result := ModelVoucherToCoreVoucher(v)
		list = append(list, result)
	}
	return list
}

func CoreExchangeVoucherToModelExchangeVoucher(data ExchangeVoucherCore) model.ExchangeVoucher {
	return model.ExchangeVoucher{
		IdUser:          data.IdUser,
		IdVoucher:       data.IdVoucher,
		Phone:           data.Phone,
		Status:          data.Status,
		TimeTransaction: data.TimeTransaction,
	}
}

func ListCoreExchangeVoucherToModelExchangeVoucher(data []ExchangeVoucherCore) []model.ExchangeVoucher {
	list := []model.ExchangeVoucher{}
	for _, v := range data {
		result := CoreExchangeVoucherToModelExchangeVoucher(v)
		list = append(list, result)
	}
	return list
}

func ModelExchangeVoucherToCoreExchangeVoucher(data model.ExchangeVoucher) ExchangeVoucherCore {
	return ExchangeVoucherCore{
		Id:              data.Id,
		IdUser:          data.IdUser,
		IdVoucher:       data.IdVoucher,
		Phone:           data.Phone,
		Status:          data.Status,
		TimeTransaction: data.TimeTransaction,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}
}

func ModelExchangeVoucherToMapDetail(data model.ExchangeVoucher, point int) map[string]interface{} {
	return map[string]interface{}{
		"id_transaction":   data.Id,
		"voucher":          data.IdVoucher,
		"points":           point,
		"phone":            data.Phone,
		"status":           data.Status,
		"time_transaction": data.TimeTransaction,
		"type_transaction": "tukar poin",
		"created_at":       data.CreatedAt.Format(time.RFC3339),
	}
}

func ModelExchangeVoucherToMap(data model.ExchangeVoucher, point int) map[string]interface{} {
	return map[string]interface{}{
		"id_transaction":   data.Id,
		"points":           point,
		"time_transaction": data.TimeTransaction,
		"type_transaction": "tukar poin",
		"created_at":       data.CreatedAt.Format(time.RFC3339),
	}
}

func ListModelExchangeVoucherToCoreExchangeVoucher(data []model.ExchangeVoucher) []ExchangeVoucherCore {
	list := []ExchangeVoucherCore{}
	for _, v := range data {
		result := ModelExchangeVoucherToCoreExchangeVoucher(v)
		list = append(list, result)
	}
	return list
}
