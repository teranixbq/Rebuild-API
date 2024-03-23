package entity

import "recything/features/daily_point/model"

func DailyPointModelToDailyPointCore(dailyPoint model.DailyPoint) DailyPointCore{
	return DailyPointCore{
		ID: dailyPoint.Id,
		Point: dailyPoint.Point,
		Description: dailyPoint.Description,
	}
}

func ListDailyPointModelToDailyPointCore(dailyPoint []model.DailyPoint) []DailyPointCore{
	coreDaily := []DailyPointCore{}
	for _, v := range dailyPoint {
		daily := DailyPointModelToDailyPointCore(v)
		coreDaily = append(coreDaily, daily)
	}
	return coreDaily
}

func DailyPointCoreToDailyPointModel(dailyPoint DailyPointCore) model.DailyPoint{
	return model.DailyPoint{
		Id: dailyPoint.ID,
		Point: dailyPoint.Point,
		Description: dailyPoint.Description,
	}
}

func ListDailyPointCoreToDailyPointModel(dailyPoint []DailyPointCore) []model.DailyPoint{
	coreDaily := []model.DailyPoint{}
	for _, v := range dailyPoint {
		daily := DailyPointCoreToDailyPointModel(v)
		coreDaily = append(coreDaily, daily)
	}
	return coreDaily
}