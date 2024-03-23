package repository

import (
	"log"
	"recything/features/dashboard/entity"
	report "recything/features/report/entity"
	modelReport "recything/features/report/model"
	trash "recything/features/trash_exchange/entity"
	modelTrash "recything/features/trash_exchange/model"
	user "recything/features/user/entity"
	modelUser "recything/features/user/model"
	voucher "recything/features/voucher/entity"
	modelVoucher "recything/features/voucher/model"
	"recything/utils/dashboard"
	"time"

	"gorm.io/gorm"
)

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) entity.DashboardRepositoryInterface {
	return &dashboardRepository{
		db: db,
	}
}

// CountUserActive implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountUserActive() ([]user.UsersCore, []report.ReportCore, error) {
	now := time.Now()

	// Cari pengguna yang diupdate dalam bulan ini
	users := []modelUser.Users{}
	err := dr.db.Where("MONTH(updated_at) = ? AND YEAR(updated_at) = ?", now.Month(), now.Year()).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	// Cari laporan yang dibuat dalam bulan ini
	reports := []modelReport.Report{}
	err = dr.db.Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", now.Month(), now.Year()).Find(&reports).Error
	if err != nil {
		return nil, nil, err
	}

	// Memetakan data model ke core
	mappedUsers := user.ListUserModelToUserCore(users)
	mappedReports := report.ListReportModelToReportCore(reports)

	return mappedUsers, mappedReports, nil
}

// CountUserActiveLastMonth implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountUserActiveLastMonth() ([]user.UsersCore, []report.ReportCore, error) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Cari pengguna yang diupdate dalam bulan ini
	users := []modelUser.Users{}
	err := dr.db.Where("MONTH(updated_at) = ? AND YEAR(updated_at) = ?", lastMonth.Month(), lastMonth.Year()).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	// Cari laporan yang dibuat dalam bulan ini
	reports := []modelReport.Report{}
	err = dr.db.Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", lastMonth.Month(), lastMonth.Year()).Find(&reports).Error
	if err != nil {
		return nil, nil, err
	}

	// Memetakan data model ke core
	mappedUsers := user.ListUserModelToUserCore(users)
	mappedReports := report.ListReportModelToReportCore(reports)

	return mappedUsers, mappedReports, nil
}

// CountVoucherExchanges implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountVoucherExchanges() ([]voucher.ExchangeVoucherCore, []voucher.ExchangeVoucherCore, error) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Ambil data pertukaran voucher bulan ini
	var exchangesThisMonth []modelVoucher.ExchangeVoucher
	if err := dr.db.Model(&modelVoucher.ExchangeVoucher{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", now.Month(), now.Year()).
		Find(&exchangesThisMonth).Error; err != nil {
		return nil, nil, err
	}

	// Ambil data pertukaran voucher bulan lalu
	var exchangesLastMonth []modelVoucher.ExchangeVoucher
	if err := dr.db.Model(&modelVoucher.ExchangeVoucher{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", lastMonth.Month(), lastMonth.Year()).
		Find(&exchangesLastMonth).Error; err != nil {
		return nil, nil, err
	}

	// Konversi dari model ke core
	coreThisMonth := voucher.ListModelExchangeVoucherToCoreExchangeVoucher(exchangesThisMonth)
	coreLastMonth := voucher.ListModelExchangeVoucherToCoreExchangeVoucher(exchangesLastMonth)

	return coreThisMonth, coreLastMonth, nil
}

// CountReports implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountReports() ([]report.ReportCore, []report.ReportCore, error) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Hitung total pelaporan bulan ini
	var reportThisMonth []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", now.Month(), now.Year()).
		Find(&reportThisMonth).Error; err != nil {
		return nil, nil, err
	}

	// Hitung total pelaporan bulan lalu
	var reportLastMonth []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", lastMonth.Month(), lastMonth.Year()).
		Find(&reportLastMonth).Error; err != nil {
		return nil, nil, err
	}

	coreThisMonth := report.ListReportModelToReportCore(reportThisMonth)
	coreLastMonth := report.ListReportModelToReportCore(reportLastMonth)

	return coreThisMonth, coreLastMonth, nil
}

// CountTrashExchanges implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountTrashExchanges() ([]trash.TrashExchangeCore, []trash.TrashExchangeCore, error) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Hitung total TrashExchange bulan ini
	var totalThisMonth []modelTrash.TrashExchange
	if err := dr.db.Model(&modelTrash.TrashExchange{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", now.Month(), now.Year()).
		Find(&totalThisMonth).Error; err != nil {
		return nil, nil, err
	}

	// Hitung total TrashExchange bulan lalu
	var totalLastMonth []modelTrash.TrashExchange
	if err := dr.db.Model(&modelTrash.TrashExchange{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", lastMonth.Month(), lastMonth.Year()).
		Find(&totalLastMonth).Error; err != nil {
		return nil, nil, err
	}

	coreThisMonth := trash.ListTrashExchangeModelToTrashExchangeCoreForGetData(totalThisMonth)
	coreLastMonth := trash.ListTrashExchangeModelToTrashExchangeCoreForGetData(totalLastMonth)

	return coreThisMonth, coreLastMonth, nil
}

// CountScaleTypes implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountCategory() ([]report.ReportCore, []report.ReportCore, error) {
	now := time.Now()

	// Hitung total pelaporan dengan company_name terisi bulan ini
	var totalWithCompanyName []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ? AND company_name IS NOT NULL AND company_name != ''", now.Month(), now.Year()).
		Find(&totalWithCompanyName).Error; err != nil {
		return nil, nil, err
	}

	// Hitung total pelaporan dengan company_name tidak terisi bulan ini
	var totalWithoutCompanyName []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ? AND (company_name IS NULL OR company_name = '')", now.Month(), now.Year()).
		Find(&totalWithoutCompanyName).Error; err != nil {
		return nil, nil, err
	}

	coreWithCompanyName := report.ListReportModelToReportCore(totalWithCompanyName)
	coreWithoutCompanyName := report.ListReportModelToReportCore(totalWithoutCompanyName)

	return coreWithCompanyName, coreWithoutCompanyName, nil
}

// GetUserRanking implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) GetUserRanking() ([]user.UsersCore, error) {
	now := time.Now()

	var userPoints []modelUser.Users
	limit := 3

	// Ambil peringkat pengguna untuk bulan ini
	err := dr.db.Model(&modelUser.Users{}).
		Where("MONTH(updated_at) = ? AND YEAR(updated_at) = ?", now.Month(), now.Year()).
		Order("point DESC").
		Limit(limit).
		Find(&userPoints).Error
	if err != nil {
		return nil, err
	}

	mappedUsers := user.ListUserModelToUserCore(userPoints)
	return mappedUsers, nil
}

// CountWeeklyTrashAndScalaTypes implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountWeeklyTrashAndScalaTypes() ([]report.ReportCore, error) {
	var trashAndScalaTypes []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Find(&trashAndScalaTypes).Error; err != nil {
		return nil, err
	}

	coreThisMonth := report.ListReportModelToReportCore(trashAndScalaTypes)
	return coreThisMonth, nil
}

// Years

// CountUserActiveLastYear implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountUserActiveLastYear() ([]user.UsersCore, []report.ReportCore, error) {
	now := time.Now()
	lastYear := now.AddDate(-1, 0, 0)

	// Cari pengguna yang diupdate dalam tahun lalu
	users := []modelUser.Users{}
	err := dr.db.Where("YEAR(updated_at) = ?", lastYear.Year()).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	// Cari laporan yang dibuat dalam tahun lalu
	reports := []modelReport.Report{}
	err = dr.db.Where("YEAR(created_at) = ?", lastYear.Year()).Find(&reports).Error
	if err != nil {
		return nil, nil, err
	}

	mappedUsers := user.ListUserModelToUserCore(users)
	mappedReports := report.ListReportModelToReportCore(reports)

	return mappedUsers, mappedReports, nil
}

// CountUserActiveThisYears implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountUserActiveThisYear() ([]user.UsersCore, []report.ReportCore, error) {
	now := time.Now()

	// Cari pengguna yang diupdate dalam tahun ini
	users := []modelUser.Users{}
	err := dr.db.Where("YEAR(updated_at) = ?", now.Year()).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	// Cari laporan yang dibuat dalam tahun ini
	reports := []modelReport.Report{}
	err = dr.db.Where("YEAR(created_at) = ?", now.Year()).Find(&reports).Error
	if err != nil {
		return nil, nil, err
	}

	mappedUsers := user.ListUserModelToUserCore(users)
	mappedReports := report.ListReportModelToReportCore(reports)

	return mappedUsers, mappedReports, nil
}

// CountVoucherExchangesYear implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountVoucherExchangesYear() ([]voucher.ExchangeVoucherCore, []voucher.ExchangeVoucherCore, error) {
	now := time.Now()

	// Ambil data pertukaran voucher tahun ini
	var exchangesThisYear []modelVoucher.ExchangeVoucher
	if err := dr.db.Model(&modelVoucher.ExchangeVoucher{}).
		Where("YEAR(created_at) = ?", now.Year()).
		Find(&exchangesThisYear).Error; err != nil {
		return nil, nil, err
	}

	// Ambil data pertukaran voucher tahun lalu
	lastYear := now.AddDate(-1, 0, 0)
	var exchangesLastYear []modelVoucher.ExchangeVoucher
	if err := dr.db.Model(&modelVoucher.ExchangeVoucher{}).
		Where("YEAR(created_at) = ?", lastYear.Year()).
		Find(&exchangesLastYear).Error; err != nil {
		return nil, nil, err
	}

	coreThisYear := voucher.ListModelExchangeVoucherToCoreExchangeVoucher(exchangesThisYear)
	coreLastYear := voucher.ListModelExchangeVoucherToCoreExchangeVoucher(exchangesLastYear)

	return coreThisYear, coreLastYear, nil
}

// CountReportsYear implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountReportsYear() ([]report.ReportCore, []report.ReportCore, error) {
	now := time.Now()

	// Hitung total pelaporan tahun ini
	var reportThisYear []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Where("YEAR(created_at) = ?", now.Year()).
		Find(&reportThisYear).Error; err != nil {
		return nil, nil, err
	}

	// Hitung total pelaporan tahun lalu
	lastYear := now.AddDate(-1, 0, 0)
	var reportLastYear []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Where("YEAR(created_at) = ?", lastYear.Year()).
		Find(&reportLastYear).Error; err != nil {
		return nil, nil, err
	}

	coreThisYear := report.ListReportModelToReportCore(reportThisYear)
	coreLastYear := report.ListReportModelToReportCore(reportLastYear)

	return coreThisYear, coreLastYear, nil
}

// CountTrashExchangesYear implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountTrashExchangesYear() ([]trash.TrashExchangeCore, []trash.TrashExchangeCore, error) {
	now := time.Now()

	// Hitung total TrashExchange tahun ini
	var totalThisYear []modelTrash.TrashExchange
	if err := dr.db.Model(&modelTrash.TrashExchange{}).
		Where("YEAR(created_at) = ?", now.Year()).
		Find(&totalThisYear).Error; err != nil {
		return nil, nil, err
	}

	// Hitung total TrashExchange tahun lalu
	lastYear := now.AddDate(-1, 0, 0)
	var totalLastYear []modelTrash.TrashExchange
	if err := dr.db.Model(&modelTrash.TrashExchange{}).
		Where("YEAR(created_at) = ?", lastYear.Year()).
		Find(&totalLastYear).Error; err != nil {
		return nil, nil, err
	}

	coreThisYear := trash.ListTrashExchangeModelToTrashExchangeCoreForGetData(totalThisYear)
	coreLastYear := trash.ListTrashExchangeModelToTrashExchangeCoreForGetData(totalLastYear)

	return coreThisYear, coreLastYear, nil
}

// CountScaleTypesYear implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountCategoryYear() ([]report.ReportCore, []report.ReportCore, error) {
	var totalWithCompanyNameThisYear []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Where("YEAR(created_at) = ? AND company_name IS NOT NULL AND company_name != ''", time.Now().Year()).
		Find(&totalWithCompanyNameThisYear).Error; err != nil {
		return nil, nil, err
	}

	// Hitung total pelaporan dengan company_name tidak terisi tahun ini
	var totalWithoutCompanyNameThisYear []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Where("YEAR(created_at) = ? AND (company_name IS NULL OR company_name = '')", time.Now().Year()).
		Find(&totalWithoutCompanyNameThisYear).Error; err != nil {
		return nil, nil, err
	}

	coreWithCompanyNameThisYear := report.ListReportModelToReportCore(totalWithCompanyNameThisYear)
	coreWithoutCompanyNameThisYear := report.ListReportModelToReportCore(totalWithoutCompanyNameThisYear)

	return coreWithCompanyNameThisYear, coreWithoutCompanyNameThisYear, nil
}

// GetUserRankingYear implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) GetUserRankingYear() ([]user.UsersCore, error) {
	now := time.Now()
	var userPointsThisYear []modelUser.Users
	limit := 3
	err := dr.db.Model(&modelUser.Users{}).
		Where("YEAR(updated_at) = ?", now.Year()).
		Order("point DESC").
		Limit(limit).
		Find(&userPointsThisYear).Error
	if err != nil {
		return nil, err
	}

	mappedUsers := user.ListUserModelToUserCore(userPointsThisYear)
	return mappedUsers, nil
}

func (dr *dashboardRepository) CountWeeklyTrashAndScalaTypesYear() ([]report.ReportCore, error) {
	var trashAndScalaTypes []modelReport.Report
	if err := dr.db.Model(&modelReport.Report{}).
		Find(&trashAndScalaTypes).Error; err != nil {
		return nil, err
	}

	coreThisMonth := report.ListReportModelToReportCore(trashAndScalaTypes)
	return coreThisMonth, nil
}

// CountTrashExchangesIncome implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountTrashExchangesIncome() (dashboard.TrashIncomeStats, error) {
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	// Ambil awal bulan
	firstOfNextMonth := firstOfMonth.AddDate(0, 1, 0)
	log.Printf("Now: %s, First of Month: %s, First of Next Month: %s\n", now, firstOfMonth, firstOfNextMonth)

	var totalIncomeThisMonth int
	if err := dr.db.Model(&modelTrash.TrashExchange{}).
		Where("created_at >= ? AND created_at < ?", firstOfMonth, firstOfNextMonth).
		Pluck("COALESCE(SUM(total_income), 0)", &totalIncomeThisMonth).
		Error; err != nil {
		return dashboard.TrashIncomeStats{}, err
	}

	// Bulan lalu
	var totalIncomeLastMonth int
	if err := dr.db.Model(&modelTrash.TrashExchange{}).
		Where("created_at >= ? AND created_at < ?", firstOfMonth.AddDate(0, -1, 0), firstOfMonth).
		Pluck("COALESCE(SUM(total_income), 0)", &totalIncomeLastMonth).
		Error; err != nil {
		return dashboard.TrashIncomeStats{}, err
	}

	data := dashboard.MapTrashIncomeStats(totalIncomeThisMonth, totalIncomeLastMonth)
	return data, nil
}

// CountTrashExchangesIncomeYear implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountTrashExchangesIncomeYear() (dashboard.TrashIncomeStats, error) {
	now := time.Now()
    firstOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)

    // Awal Tahun
    firstOfNextYear := firstOfYear.AddDate(1, 0, 0)
    log.Printf("Now: %s, First of Year: %s, First of Next Year: %s\n", now, firstOfYear, firstOfNextYear)

    var totalIncomeThisYear int
    if err := dr.db.Model(&modelTrash.TrashExchange{}).
        Where("created_at >= ? AND created_at < ?", firstOfYear, firstOfNextYear).
        Pluck("COALESCE(SUM(total_income), 0)", &totalIncomeThisYear).
        Error; err != nil {
        return dashboard.TrashIncomeStats{}, err
    }

    // Tahun lalu
    var totalIncomeLastYear int
    if err := dr.db.Model(&modelTrash.TrashExchange{}).
        Where("created_at >= ? AND created_at < ?", firstOfYear.AddDate(-1, 0, 0), firstOfYear).
        Pluck("COALESCE(SUM(total_income), 0)", &totalIncomeLastYear).
        Error; err != nil {
        return dashboard.TrashIncomeStats{}, err
    }

    data := dashboard.MapTrashIncomeStats(totalIncomeThisYear, totalIncomeLastYear)
    return data, nil
}