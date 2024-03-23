package service

import (
	"errors"
	re "recything/features/report/entity"
	te "recything/features/trash_exchange/entity"
	ue "recything/features/user/entity"
	ve "recything/features/voucher/entity"
	"recything/mocks"
	"recything/utils/dashboard"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashboardMonthlySuccess(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	// Set up the expected behavior for CountUserActiveLastMonth
	repoData.On("CountUserActiveLastMonth").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	// Set up the expected behavior for GetUserRanking
	repoData.On("GetUserRanking").Return(
		[]ue.UsersCore{
			{Id: "user1", Point: 100},
			{Id: "user2", Point: 90},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchanges
	repoData.On("CountVoucherExchanges").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	// Set up the expected behavior for CountReports
	repoData.On("CountReports").Return(
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		[]re.ReportCore{
			{ID: "report7"},
			{ID: "report8"},
		},
		nil,
	)

	// Set up the expected behavior for CountTrashExchanges
	repoData.On("CountTrashExchanges").Return(
		[]te.TrashExchangeCore{
			{Id: "trash1"},
			{Id: "trash2"},
		},
		[]te.TrashExchangeCore{
			{Id: "trash3"},
			{Id: "trash4"},
		},
		nil,
	)

	// Set up the expected behavior for CountCategory
	repoData.On("CountCategory").Return(
		[]re.ReportCore{
			{ID: "category1"},
			{ID: "category2"},
		},
		[]re.ReportCore{
			{ID: "category3"},
			{ID: "category4"},
		},
		nil,
	)

	// Set up the expected behavior for CountWeeklyTrashAndScalaTypes
	repoData.On("CountWeeklyTrashAndScalaTypes").Return(
		[]re.ReportCore{
			{ID: "weekly1"},
			{ID: "weekly2"},
		},
		nil,
	)

	// Set up the expected behavior for CountTrashExchangesIncome
	repoData.On("CountTrashExchangesIncome").Return(
		dashboard.TrashIncomeStats{
			TotalIncomeThisMonth: 200,
			TotalIncomeLastMonth: 180,
		},
		nil,
	)

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardMonthly()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users.TotalPenggunaAktif))
	assert.Equal(t, 1, len(vouchers.TotalPenukaran))
	assert.Equal(t, 1, len(reports.TotalReporting)) // Sesuaikan panjang dengan jumlah laporan
	assert.Equal(t, 1, len(trash.TotalTrashExchange))
	assert.Equal(t, 2, len(scaleTypes.Company))
	assert.Equal(t, 2, len(userRanking))
	assert.Equal(t, 5, len(weeklyStats)) // Sesuaikan panjang dengan jumlah minggu
	assert.Equal(t, 200, income.TotalIncome)
}

func TestDashboardMonthlyCountUserActiveLastMonthError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastMonth").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardMonthly()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardMonthlyCountVoucherActiveError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastMonth").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	repoData.On("CountVoucherExchanges").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardMonthly()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardMonthlyCountReportsError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastMonth").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchanges
	repoData.On("CountVoucherExchanges").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	repoData.On("CountReports").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardMonthly()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardMonthlyCountCategoryError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastMonth").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchanges
	repoData.On("CountVoucherExchanges").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	// Set up the expected behavior for CountReports
	repoData.On("CountReports").Return(
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		[]re.ReportCore{
			{ID: "report7"},
			{ID: "report8"},
		},
		nil,
	)

	// Set up the expected behavior for CountTrashExchanges
	repoData.On("CountTrashExchanges").Return(
		[]te.TrashExchangeCore{
			{Id: "trash1"},
			{Id: "trash2"},
		},
		[]te.TrashExchangeCore{
			{Id: "trash3"},
			{Id: "trash4"},
		},
		nil,
	)

	repoData.On("CountCategory").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardMonthly()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardMonthlyCountWeeklyTrashAndScalaTypesError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	// Set up the expected behavior for GetUserRanking
	repoData.On("GetUserRanking").Return(
		[]ue.UsersCore{
			{Id: "user1", Point: 100},
			{Id: "user2", Point: 90},
		},
		nil,
	)

	repoData.On("CountUserActiveLastMonth").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchanges
	repoData.On("CountVoucherExchanges").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	// Set up the expected behavior for CountReports
	repoData.On("CountReports").Return(
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		[]re.ReportCore{
			{ID: "report7"},
			{ID: "report8"},
		},
		nil,
	)

	// Set up the expected behavior for CountTrashExchanges
	repoData.On("CountTrashExchanges").Return(
		[]te.TrashExchangeCore{
			{Id: "trash1"},
			{Id: "trash2"},
		},
		[]te.TrashExchangeCore{
			{Id: "trash3"},
			{Id: "trash4"},
		},
		nil,
	)

	// Set up the expected behavior for CountCategory
	repoData.On("CountCategory").Return(
		[]re.ReportCore{
			{ID: "category1"},
			{ID: "category2"},
		},
		[]re.ReportCore{
			{ID: "category3"},
			{ID: "category4"},
		},
		nil,
	)

	// Set up the expected behavior for CountTrashExchangesIncome
	repoData.On("CountTrashExchangesIncome").Return(
		dashboard.TrashIncomeStats{
			TotalIncomeThisMonth: 200,
			TotalIncomeLastMonth: 180,
		},
		nil,
	)

	repoData.On("CountWeeklyTrashAndScalaTypes").Return(nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardMonthly()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardMonthlyCountTotalTrashActiveError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastMonth").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchanges
	repoData.On("CountVoucherExchanges").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	// Set up the expected behavior for CountReports
	repoData.On("CountReports").Return(
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		[]re.ReportCore{
			{ID: "report7"},
			{ID: "report8"},
		},
		nil,
	)

	repoData.On("CountTrashExchanges").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardMonthly()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardMonthlyCountUserActiveError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardMonthly()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}


//years

func TestDashboardYearsSuccess(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActiveThisYear
	repoData.On("CountUserActiveThisYear").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	// Set up the expected behavior for CountUserActiveLastYear
	repoData.On("CountUserActiveLastYear").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	// Set up the expected behavior for GetUserRankingYear
	repoData.On("GetUserRankingYear").Return(
		[]ue.UsersCore{
			{Id: "user1", Point: 100},
			{Id: "user2", Point: 90},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchangesYear
	repoData.On("CountVoucherExchangesYear").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	// Set up the expected behavior for CountReportsYear
	repoData.On("CountReportsYear").Return(
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		[]re.ReportCore{
			{ID: "report7"},
			{ID: "report8"},
		},
		nil,
	)

	// Set up the expected behavior for CountTrashExchangesYear
	repoData.On("CountTrashExchangesYear").Return(
		[]te.TrashExchangeCore{
			{Id: "trash1"},
			{Id: "trash2"},
		},
		[]te.TrashExchangeCore{
			{Id: "trash3"},
			{Id: "trash4"},
		},
		nil,
	)

	// Set up the expected behavior for CountCategoryYear
	repoData.On("CountCategoryYear").Return(
		[]re.ReportCore{
			{ID: "category1"},
			{ID: "category2"},
		},
		[]re.ReportCore{
			{ID: "category3"},
			{ID: "category4"},
		},
		nil,
	)

	// Set up the expected behavior for GetUserRankingYear
	repoData.On("GetUserRankingYear").Return(
		[]ue.UsersCore{
			{Id: "user1", Point: 100},
			{Id: "user2", Point: 90},
		},
		nil,
	)

	// Set up the expected behavior for CountWeeklyTrashAndScalaTypes
	repoData.On("CountWeeklyTrashAndScalaTypes").Return(
		[]re.ReportCore{
			{ID: "weekly1"},
			{ID: "weekly2"},
		},
		nil,
	)

	// Set up the expected behavior for CountTrashExchangesIncomeYear
	repoData.On("CountTrashExchangesIncomeYear").Return(
		dashboard.TrashIncomeStats{
			TotalIncomeThisMonth: 200,
			TotalIncomeLastMonth: 180,
		},
		nil,
	)

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, monthlyStats, income, err := dashboardService.DashboardYears()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users.TotalPenggunaAktif))
	assert.Equal(t, 1, len(vouchers.TotalPenukaran))
	assert.Equal(t, 1, len(reports.TotalReporting))
	assert.Equal(t, 1, len(trash.TotalTrashExchange))
	assert.Equal(t, 2, len(scaleTypes.Company))
	assert.Equal(t, 2, len(userRanking))
	assert.Equal(t, 12, len(monthlyStats)) // Assuming 12 months
	assert.Equal(t, 200, income.TotalIncome)
}

func TestDashboardYearsCountUserActiveLastYearError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActiveThisYear").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastYear").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardYears()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardYearCountVoucherActiveError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveThisYear").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastYear").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		nil,
	)

	repoData.On("CountVoucherExchangesYear").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardYears()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardYearsCountReportsError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActiveThisYear").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastYear").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchanges
	repoData.On("CountVoucherExchangesYear").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	repoData.On("CountReportsYear").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardYears()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardYearsCountCategoryError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActiveThisYear").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastYear").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchanges
	repoData.On("CountVoucherExchangesYear").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	// Set up the expected behavior for CountReports
	repoData.On("CountReportsYear").Return(
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		[]re.ReportCore{
			{ID: "report7"},
			{ID: "report8"},
		},
		nil,
	)

	// Set up the expected behavior for CountTrashExchanges
	repoData.On("CountTrashExchangesYear").Return(
		[]te.TrashExchangeCore{
			{Id: "trash1"},
			{Id: "trash2"},
		},
		[]te.TrashExchangeCore{
			{Id: "trash3"},
			{Id: "trash4"},
		},
		nil,
	)
	
	repoData.On("CountCategoryYear").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardYears()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardYearCountTotalTrashActiveError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActive
	repoData.On("CountUserActive").Return(
		[]ue.UsersCore{
			{Id: "user1"},
			{Id: "user2"},
		},
		[]re.ReportCore{
			{ID: "report1"},
			{ID: "report2"},
		},
		nil,
	)

	repoData.On("CountUserActiveThisYear").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report3"},
			{ID: "report4"},
		},
		nil,
	)

	repoData.On("CountUserActiveLastYear").Return(
		[]ue.UsersCore{
			{Id: "user3"},
			{Id: "user4"},
		},
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		nil,
	)

	// Set up the expected behavior for CountVoucherExchanges
	repoData.On("CountVoucherExchangesYear").Return(
		[]ve.ExchangeVoucherCore{
			{Id: "voucher1"},
			{Id: "voucher2"},
		},
		[]ve.ExchangeVoucherCore{
			{Id: "voucher3"},
			{Id: "voucher4"},
		},
		nil,
	)

	// Set up the expected behavior for CountReports
	repoData.On("CountReportsYear").Return(
		[]re.ReportCore{
			{ID: "report5"},
			{ID: "report6"},
		},
		[]re.ReportCore{
			{ID: "report7"},
			{ID: "report8"},
		},
		nil,
	)

	repoData.On("CountTrashExchangesYear").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, weeklyStats, income, err := dashboardService.DashboardYears()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t,"", users.TotalPenggunaAktif)
	assert.Equal(t,"", vouchers.TotalPenukaran)
	assert.Equal(t,"", reports.TotalReporting)
	assert.Equal(t,"", trash.TotalTrashExchange)
	assert.Equal(t,"", scaleTypes.Company)
	assert.Nil(t,userRanking)
	assert.Nil(t,weeklyStats)
	assert.Equal(t,"", income.Persentase)
}

func TestDashboardYearsCountUserActiveError(t *testing.T) {
	repoData := new(mocks.DashboardRepositoryInterface)
	dashboardService := NewDashboardService(repoData)

	// Set up the expected behavior for CountUserActiveThisYear with an error
	repoData.On("CountUserActiveThisYear").Return(nil, nil, errors.New("some error"))

	// Call the method
	users, vouchers, reports, trash, scaleTypes, userRanking, monthlyStats, income, err := dashboardService.DashboardYears()

	// Assertions for error case
	assert.Error(t, err)
	assert.Equal(t, "", users.TotalPenggunaAktif)
	assert.Equal(t, "", vouchers.TotalPenukaran)
	assert.Equal(t, "", reports.TotalReporting)
	assert.Equal(t, "", trash.TotalTrashExchange)
	assert.Equal(t, "", scaleTypes.Company)
	assert.Nil(t, userRanking)
	assert.Nil(t, monthlyStats)
	assert.Equal(t, "", income.Persentase)
}
