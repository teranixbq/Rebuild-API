package service

import (
	"recything/features/dashboard/entity"
	"recything/utils/dashboard"
	"time"
)

type dashboardService struct {
	dashboardRepository entity.DashboardRepositoryInterface
}

func NewDashboardService(dashboardRepo entity.DashboardRepositoryInterface) entity.DashboardServiceInterface {
	return &dashboardService{
		dashboardRepository: dashboardRepo,
	}
}

// CountUserActive implements entity.DashboardUsecaseInterface.
func (ds *dashboardService) DashboardMonthly() (dashboard.GetCountUser, dashboard.GetCountExchangeVoucher, dashboard.GetCountReporting, dashboard.GetCountTrashExchange, dashboard.GetCountScaleType, []dashboard.UserRanking, []dashboard.WeeklyStats, dashboard.GetCountTrashExchangeIncome, error) {
	users, reports, err := ds.dashboardRepository.CountUserActive()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	usersLastMonth, reportsLastMonth, err := ds.dashboardRepository.CountUserActiveLastMonth()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	userResult, err := dashboard.CalculateAndMapUserStats(users, usersLastMonth, reports, reportsLastMonth)
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	voucherThisMonth, voucherLastMonth, err := ds.dashboardRepository.CountVoucherExchanges()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	totalReportsThisMonth, totalReportsLastMonth, err := ds.dashboardRepository.CountReports()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	totalTrashThisMonth, totalTrashLastMonth, err := ds.dashboardRepository.CountTrashExchanges()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	totalLargeScale, totalSmallScale, err := ds.dashboardRepository.CountCategory()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	pointUsers, err := ds.dashboardRepository.GetUserRanking()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	trashAndScalaTypes, err := ds.dashboardRepository.CountWeeklyTrashAndScalaTypes()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	trashIncomeStats, err := ds.dashboardRepository.CountTrashExchangesIncome()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	firstOfMonth := time.Now().AddDate(0, 0, -time.Now().Day()+1)
	startOfMonth := time.Date(firstOfMonth.Year(), firstOfMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	weeklyStats := dashboard.CalculateWeeklyStats(trashAndScalaTypes, startOfMonth)

	userRanking := dashboard.MapUserRanking(pointUsers)
	trashResult := dashboard.MapToGetCountTrashExchange(len(totalTrashThisMonth), len(totalTrashLastMonth))
	reportResult := dashboard.MapToGetCountReporting(len(totalReportsThisMonth), len(totalReportsLastMonth))
	voucherResult := dashboard.MapToGetCountExchangeVoucher(len(voucherThisMonth), len(voucherLastMonth))
	scalaResult := dashboard.MapToGetCountScaleTypePercentage(len(totalLargeScale), len(totalSmallScale))
	incomeResult := dashboard.MapToGetCountIncome(trashIncomeStats.TotalIncomeThisMonth, trashIncomeStats.TotalIncomeLastMonth)
	return userResult, voucherResult, reportResult, trashResult, scalaResult, userRanking, weeklyStats, incomeResult, nil
}

// DashboardYears implements entity.DashboardServiceInterface.
func (ds *dashboardService) DashboardYears() (dashboard.GetCountUser, dashboard.GetCountExchangeVoucher, dashboard.GetCountReporting, dashboard.GetCountTrashExchange, dashboard.GetCountScaleType, []dashboard.UserRanking, []dashboard.MonthlyStats, dashboard.GetCountTrashExchangeIncome, error) {
	users, reports, err := ds.dashboardRepository.CountUserActiveThisYear()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	usersLastYears, reportsLastYears, err := ds.dashboardRepository.CountUserActiveLastYear()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	userResult, err := dashboard.CalculateAndMapUserStats(users, usersLastYears, reports, reportsLastYears)
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	voucherThisYears, voucherLastYears, err := ds.dashboardRepository.CountVoucherExchangesYear()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	totalReportsThisYears, totalReportsLastYears, err := ds.dashboardRepository.CountReportsYear()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	totalTrashThisYears, totalTrashLastYears, err := ds.dashboardRepository.CountTrashExchangesYear()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	totalLargeScale, totalSmallScale, err := ds.dashboardRepository.CountCategoryYear()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	pointUsers, err := ds.dashboardRepository.GetUserRankingYear()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	trashAndScalaTypes, err := ds.dashboardRepository.CountWeeklyTrashAndScalaTypes()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	trashIncomeStats, err := ds.dashboardRepository.CountTrashExchangesIncomeYear()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, nil, dashboard.GetCountTrashExchangeIncome{}, err
	}

	// Menghitung jumlah bulan dalam setahun
	currentTime := time.Now()
	startOfYear := time.Date(currentTime.Year(), 1, 1, 0, 0, 0, 0, currentTime.Location())
	monthsInYear := 12

	monthlyStats := dashboard.CalculateMonthlyStats(trashAndScalaTypes, startOfYear, monthsInYear)
	userRanking := dashboard.MapUserRanking(pointUsers)
	trashResult := dashboard.MapToGetCountTrashExchange(len(totalTrashThisYears), len(totalTrashLastYears))
	reportResult := dashboard.MapToGetCountReporting(len(totalReportsThisYears), len(totalReportsLastYears))
	voucherResult := dashboard.MapToGetCountExchangeVoucher(len(voucherThisYears), len(voucherLastYears))
	scalaResult := dashboard.MapToGetCountScaleTypePercentage(len(totalLargeScale), len(totalSmallScale))
	incomeResult := dashboard.MapToGetCountIncome(trashIncomeStats.TotalIncomeThisMonth, trashIncomeStats.TotalIncomeLastMonth)

	return userResult, voucherResult, reportResult, trashResult, scalaResult, userRanking, monthlyStats, incomeResult, nil
}

