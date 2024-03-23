package dto

import "recything/utils/dashboard"

func MapToGetCountUserResponse(input dashboard.GetCountUser) GetCountUserResponse {
	return GetCountUserResponse{
		TotalUserActive: input.TotalPenggunaAktif,
		Percentage:      input.Persentase,
		Status:          input.Status,
	}
}

func MapToGetCountExchangeVoucherResponse(input dashboard.GetCountExchangeVoucher) GetCountExchangeVoucherResponse {
	return GetCountExchangeVoucherResponse{
		TotalExchange: input.TotalPenukaran,
		Percentage:    input.Persentase,
		Status:        input.Status,
	}
}

func MapToGetCountReportingResponse(input dashboard.GetCountReporting) GetCountReportingResponse {
	return GetCountReportingResponse{
		TotalReporting: input.TotalReporting,
		Percentage:     input.Persentase,
		Status:         input.Status,
	}
}

func MapToGetCountTrashExchangeResponse(input dashboard.GetCountTrashExchange) GetCountTrashExchangeResponse {
	return GetCountTrashExchangeResponse{
		TotalTrashExchange: input.TotalTrashExchange,
		Percentage:         input.Persentase,
		Status:             input.Status,
	}
}

func MapToGetCountTrashExchangeIncomeResponse(input dashboard.GetCountTrashExchangeIncome) GetCountTrashExchangeIncomeResponse {
	return GetCountTrashExchangeIncomeResponse{
		TotalIncome: input.TotalIncome,
		Percentage:  input.Persentase,
		Status:      input.Status,
	}
}

func MapToGetCountPersentaseScalaReportingResponse(input dashboard.GetCountScaleType) GetCountScaleTypeResponse {
	return GetCountScaleTypeResponse{
		Company: input.Company,
		Person:  input.Person,
	}
}

func MapToGetWeeklyStatsResponse(input dashboard.WeeklyStats) WeeklyStatsResponse {
	return WeeklyStatsResponse{
		Week:      input.Week,
		TrashType: input.Trash,
		ScaleType: input.Scala,
	}
}

func ListMapToWeeklyStatsResponses(stats []dashboard.WeeklyStats) []WeeklyStatsResponse {
	var responses []WeeklyStatsResponse
	for _, stat := range stats {
		responses = append(responses, MapToGetWeeklyStatsResponse(stat))
	}
	return responses
}

func MapToGetMonthlyStatsResponse(input dashboard.MonthlyStats) MonthlyStatsResponse {
	return MonthlyStatsResponse{
		Month:     input.Month,
		TrashType: input.Trash,
		ScaleType: input.Scala,
	}
}

func ListMapToMonthlyStatsResponses(stats []dashboard.MonthlyStats) []MonthlyStatsResponse {
	var responses []MonthlyStatsResponse
	for _, stat := range stats {
		responses = append(responses, MapToGetMonthlyStatsResponse(stat))
	}
	return responses
}

func MapToGetUserRankingResponse(input dashboard.UserRanking) UserRankingResponse {
	return UserRankingResponse{
		Name:  input.Name,
		Email: input.Email,
		Point: input.Point,
	}
}

func ListMapToGetUserRankingResponse(rankingResult []dashboard.UserRanking) []UserRankingResponse {
	var rankingResponse []UserRankingResponse
	for _, userRanking := range rankingResult {
		rankingResponse = append(rankingResponse, MapToGetUserRankingResponse(userRanking))
	}
	return rankingResponse
}

func MapToCombinedResponseYears(
	userActiveResult dashboard.GetCountUser,
	voucherResult dashboard.GetCountExchangeVoucher,
	reportResult dashboard.GetCountReporting,
	trashExchangeResult dashboard.GetCountTrashExchange,
	scalaResult dashboard.GetCountScaleType,
	rankingResult []dashboard.UserRanking,
	monthResult []dashboard.MonthlyStats,
	incomeResult dashboard.GetCountTrashExchangeIncome,
) map[string]interface{} {
	userActiveResponse := MapToGetCountUserResponse(userActiveResult)
	voucherResponse := MapToGetCountExchangeVoucherResponse(voucherResult)
	reportResponse := MapToGetCountReportingResponse(reportResult)
	trashExchangeResponse := MapToGetCountTrashExchangeResponse(trashExchangeResult)
	scalaResponse := MapToGetCountPersentaseScalaReportingResponse(scalaResult)
	rankingResponse := ListMapToGetUserRankingResponse(rankingResult)
	monthlyResponse := ListMapToMonthlyStatsResponses(monthResult)
	incomeResponse := MapToGetCountTrashExchangeIncomeResponse(incomeResult)
	combinedResponse := map[string]interface{}{
		"user_active": userActiveResponse,
		"exchange":    voucherResponse,
		"report":      reportResponse,
		"recycle":     trashExchangeResponse,
		"scale":       scalaResponse,
		"ranking":     rankingResponse,
		"years":       monthlyResponse,
		"income":      incomeResponse,
	}

	return combinedResponse
}

func MapToCombinedResponseMonthly(
	userActiveResult dashboard.GetCountUser,
	voucherResult dashboard.GetCountExchangeVoucher,
	reportResult dashboard.GetCountReporting,
	trashExchangeResult dashboard.GetCountTrashExchange,
	scalaResult dashboard.GetCountScaleType,
	rankingResult []dashboard.UserRanking,
	weekResult []dashboard.WeeklyStats,
	incomeResult dashboard.GetCountTrashExchangeIncome,
) map[string]interface{} {
	userActiveResponse := MapToGetCountUserResponse(userActiveResult)
	voucherResponse := MapToGetCountExchangeVoucherResponse(voucherResult)
	reportResponse := MapToGetCountReportingResponse(reportResult)
	trashExchangeResponse := MapToGetCountTrashExchangeResponse(trashExchangeResult)
	scalaResponse := MapToGetCountPersentaseScalaReportingResponse(scalaResult)
	rankingResponse := ListMapToGetUserRankingResponse(rankingResult)
	weekResponse := ListMapToWeeklyStatsResponses(weekResult)
	incomeResponse := MapToGetCountTrashExchangeIncomeResponse(incomeResult)
	combinedResponse := map[string]interface{}{
		"user_active": userActiveResponse,
		"exchange":    voucherResponse,
		"report":      reportResponse,
		"recycle":     trashExchangeResponse,
		"scale":       scalaResponse,
		"ranking":     rankingResponse,
		"monthly":     weekResponse,
		"income":      incomeResponse,
	}

	return combinedResponse
}
