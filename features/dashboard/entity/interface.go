package entity

import (
	report "recything/features/report/entity"
	trash "recything/features/trash_exchange/entity"
	user "recything/features/user/entity"
	voucher "recything/features/voucher/entity"
	"recything/utils/dashboard"
)

type DashboardRepositoryInterface interface {
	CountUserActive() ([]user.UsersCore, []report.ReportCore, error)
	CountUserActiveLastMonth() ([]user.UsersCore, []report.ReportCore, error)
	CountVoucherExchanges() ([]voucher.ExchangeVoucherCore, []voucher.ExchangeVoucherCore, error)
	CountReports() ([]report.ReportCore, []report.ReportCore, error)
	CountTrashExchanges() ([]trash.TrashExchangeCore, []trash.TrashExchangeCore, error)
	CountCategory() ([]report.ReportCore, []report.ReportCore, error)
	GetUserRanking() ([]user.UsersCore, error)
	CountWeeklyTrashAndScalaTypes() ([]report.ReportCore, error)
	CountTrashExchangesIncome() (dashboard.TrashIncomeStats, error)

	// Years
	CountUserActiveThisYear() ([]user.UsersCore, []report.ReportCore, error)
	CountUserActiveLastYear() ([]user.UsersCore, []report.ReportCore, error)
	CountVoucherExchangesYear() ([]voucher.ExchangeVoucherCore, []voucher.ExchangeVoucherCore, error)
	CountReportsYear() ([]report.ReportCore, []report.ReportCore, error)
	CountTrashExchangesYear() ([]trash.TrashExchangeCore, []trash.TrashExchangeCore, error)
	CountCategoryYear() ([]report.ReportCore, []report.ReportCore, error)
	GetUserRankingYear() ([]user.UsersCore, error)
	CountWeeklyTrashAndScalaTypesYear() ([]report.ReportCore, error)
	CountTrashExchangesIncomeYear() (dashboard.TrashIncomeStats, error)
}

type DashboardServiceInterface interface {
	DashboardMonthly() (dashboard.GetCountUser, dashboard.GetCountExchangeVoucher, dashboard.GetCountReporting, dashboard.GetCountTrashExchange, dashboard.GetCountScaleType, []dashboard.UserRanking, []dashboard.WeeklyStats, dashboard.GetCountTrashExchangeIncome, error)
	// CountWeeklyTrashAndScalaTypes() ([]dashboard.WeeklyStats, error)

	DashboardYears() (dashboard.GetCountUser, dashboard.GetCountExchangeVoucher, dashboard.GetCountReporting, dashboard.GetCountTrashExchange, dashboard.GetCountScaleType, []dashboard.UserRanking, []dashboard.MonthlyStats, dashboard.GetCountTrashExchangeIncome, error)
	// CountMonthlyTrashAndScalaTypesYear() ([]dashboard.MonthlyStats, error)
}
