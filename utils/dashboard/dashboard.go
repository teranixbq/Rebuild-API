package dashboard

import (
	"fmt"
	"log"
	"math"
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	"recything/utils/helper"
	"time"
)

type GetCountUser struct {
	TotalPenggunaAktif string
	Persentase         string
	Status             string
}

type GetCountExchangeVoucher struct {
	TotalPenukaran string
	Persentase     string
	Status         string
}

type GetCountReporting struct {
	TotalReporting string
	Persentase     string
	Status         string
}

type GetCountTrashExchange struct {
	TotalTrashExchange string
	Persentase         string
	Status             string
}

type GetCountTrashExchangeIncome struct {
	TotalIncome int
	Persentase  string
	Status      string
}

type GetCountScaleType struct {
	Company string
	Person  string
}

type UserRanking struct {
	Name  string
	Email string
	Point int
}

type WeeklyStats struct {
	Week  int
	Trash int
	Scala int
}

type MonthlyStats struct {
	Month int
	Trash int
	Scala int
}

type TrashIncomeStats struct {
	TotalIncomeThisMonth int
	TotalIncomeLastMonth int
}

func CalculateAndMapUserStats(users, usersLastMonth []user.UsersCore, reports, reportsLastMonth []report.ReportCore) (GetCountUser, error) {
	penggunaAktif := make(map[string]struct{})
	for _, u := range users {
		penggunaAktif[u.Id] = struct{}{}
	}

	for _, r := range reports {
		if _, exist := penggunaAktif[r.UserId]; !exist {
			penggunaAktif[r.UserId] = struct{}{}
		}
	}

	totalAktifBulanIni := len(penggunaAktif)

	penggunaAktifBulanLalu := make(map[string]struct{})
	for _, u := range usersLastMonth {
		penggunaAktifBulanLalu[u.Id] = struct{}{}
	}

	for _, r := range reportsLastMonth {
		if _, exist := penggunaAktifBulanLalu[r.UserId]; !exist {
			penggunaAktifBulanLalu[r.UserId] = struct{}{}
		}
	}

	totalAktifBulanLalu := len(penggunaAktifBulanLalu)

	var persentasePerubahan float64
	if totalAktifBulanLalu > 0 {
		persentasePerubahan = float64(totalAktifBulanIni-totalAktifBulanLalu) / float64(totalAktifBulanLalu) * 100
	} else {
		persentasePerubahan = 0
	}

	var status string
	if persentasePerubahan > 0 {
		status = "naik"
	} else if persentasePerubahan < 0 {
		status = "turun"
	} else {
		status = "tetap"
	}

	persentasePerubahanInt := int(math.Round(persentasePerubahan))

	result := GetCountUser{
		TotalPenggunaAktif: fmt.Sprintf("%d", totalAktifBulanIni),
		Persentase:         fmt.Sprintf("%d", persentasePerubahanInt),
		Status:             status,
	}

	return result, nil
}

func MapToGetCountExchangeVoucher(totalThisMonth, totalLastMonth int) GetCountExchangeVoucher {
	var persentasePerubahanVoucher float64
	if totalLastMonth > 0 {
		persentasePerubahanVoucher = float64(totalThisMonth-totalLastMonth) / float64(totalLastMonth) * 100
	} else {
		persentasePerubahanVoucher = 0
	}

	var statusVoucher string
	if persentasePerubahanVoucher > 0 {
		statusVoucher = "naik"
	} else if persentasePerubahanVoucher < 0 {
		statusVoucher = "turun"
	} else {
		statusVoucher = "tetap"
	}

	persentasePerubahanInt := int(math.Round(persentasePerubahanVoucher))

	// Buat map hasil untuk pertukaran voucher
	result := GetCountExchangeVoucher{
		TotalPenukaran: fmt.Sprintf("%d", totalThisMonth),
		Persentase:     fmt.Sprintf("%d", persentasePerubahanInt),
		Status:         statusVoucher,
	}

	return result
}

func MapToGetCountReporting(totalThisMonth int, totalLastMonth int) GetCountReporting {
	var persentasePerubahanReporting float64
	if totalLastMonth > 0 {
		persentasePerubahanReporting = float64(totalThisMonth-totalLastMonth) / float64(totalLastMonth) * 100
	} else {
		persentasePerubahanReporting = 0
	}

	var statusReporting string
	if persentasePerubahanReporting > 0 {
		statusReporting = "naik"
	} else if persentasePerubahanReporting < 0 {
		statusReporting = "turun"
	} else {
		statusReporting = "tetap"
	}

	persentasePerubahanInt := int(math.Round(persentasePerubahanReporting))

	// Buat map hasil untuk pertukaran voucher
	result := GetCountReporting{
		TotalReporting: fmt.Sprintf("%d", totalThisMonth),
		Persentase:     fmt.Sprintf("%d", persentasePerubahanInt),
		Status:         statusReporting,
	}

	return result
}

// MapToGetCountTrashExchange membuat objek GetCountTrashExchange dari total TrashExchange.
func MapToGetCountTrashExchange(totalThisMonth int, totalLastMonth int) GetCountTrashExchange {
	var persentasePerubahanTrash float64
	if totalLastMonth > 0 {
		persentasePerubahanTrash = float64(totalThisMonth-totalLastMonth) / float64(totalLastMonth) * 100
	} else {
		persentasePerubahanTrash = 0
	}

	var statusTrash string
	if persentasePerubahanTrash > 0 {
		statusTrash = "naik"
	} else if persentasePerubahanTrash < 0 {
		statusTrash = "turun"
	} else {
		statusTrash = "tetap"
	}

	persentasePerubahanInt := int(math.Round(persentasePerubahanTrash))

	// Buat map hasil untuk pertukaran voucher
	result := GetCountTrashExchange{
		TotalTrashExchange: fmt.Sprintf("%d", totalThisMonth),
		Persentase:         fmt.Sprintf("%d", persentasePerubahanInt),
		Status:             statusTrash,
	}

	return result
}

// MapToGetCountScaleTypePercentage membuat objek GetCountScaleType dengan persentase pelaporan skala besar dan skala kecil.
func MapToGetCountScaleTypePercentage(totalLargeScale int, totalSmallScale int) GetCountScaleType {
	totalReports := totalLargeScale + totalSmallScale

	var percentageLargeScale float64
	if totalReports > 0 {
		percentageLargeScale = float64(totalLargeScale) / float64(totalReports) * 100
	}

	var percentageSmallScale float64
	if totalReports > 0 {
		percentageSmallScale = float64(totalSmallScale) / float64(totalReports) * 100
	}

	persentaseLargeScalePerubahanInt := int(math.Round(percentageLargeScale))
	persentaseSmallScalePerubahanInt := int(math.Round(percentageSmallScale))

	// Buat map hasil untuk persentase pelaporan skala besar dan skala kecil
	result := GetCountScaleType{
		Company: fmt.Sprintf("%d", persentaseLargeScalePerubahanInt),
		Person:  fmt.Sprintf("%d", persentaseSmallScalePerubahanInt),
	}

	return result
}

func MapUserRanking(users []user.UsersCore) []UserRanking {
	var userRanking []UserRanking
	for _, user := range users {
		userRanking = append(userRanking, UserRanking{
			Name:  user.Fullname,
			Email: user.Email,
			Point: user.Point,
		})
	}
	return userRanking
}

// Function untuk memfilter data berdasarkan range tanggal
func FilterDataByDate(data []report.ReportCore, startDate, endDate time.Time) []report.ReportCore {
	var filteredData []report.ReportCore
	for _, entry := range data {
		if entry.CreatedAt.After(startDate) && entry.CreatedAt.Before(endDate) {
			filteredData = append(filteredData, entry)
		}
	}
	return filteredData
}

// Function untuk menghitung jumlah data trash_type dan scala_type
func CountTrashAndScalaTypes(data []report.ReportCore) (int, int) {
	var trashCount, scalaCount int
	for _, entry := range data {
		if entry.TrashType != "" {
			trashCount++
		}
		if entry.ScaleType != "" {
			scalaCount++
		}
	}
	return trashCount, scalaCount
}

func CalculateMonthlyStats(data []report.ReportCore, startOfYear time.Time, monthsInYear int) []MonthlyStats {
	monthlyStats := make([]MonthlyStats, monthsInYear)

	for i := 0; i < monthsInYear; i++ {
		// Menghitung awal dan akhir bulan
		monthStartDate := time.Date(startOfYear.Year(), startOfYear.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, i, 0)
		monthEndDate := time.Date(monthStartDate.Year(), monthStartDate.Month()+1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond)

		log.Printf("Month %d: StartDate: %s, EndDate: %s\n", i+1, monthStartDate, monthEndDate)

		// Filter data yang berada dalam rentang waktu bulan ini
		filteredData := FilterDataByDate(data, monthStartDate, monthEndDate)

		// Hitung jumlah data trash_type dan scala_type
		trashCount, scalaCount := CountTrashAndScalaTypes(filteredData)

		// Set nilai MonthlyStats
		monthlyStats[i].Month = i + 1
		monthlyStats[i].Trash = trashCount
		monthlyStats[i].Scala = scalaCount
	}

	return monthlyStats
}

func CalculateWeeklyStats(data []report.ReportCore, startOfMonth time.Time) []WeeklyStats {
	year, month, _ := startOfMonth.Date()
	weeksInMonth := helper.GetWeeksInMonth(year, month)
	weeklyStats := make([]WeeklyStats, weeksInMonth)

	for i := 0; i < weeksInMonth; i++ {
		// Menghitung awal dan akhir minggu
		weekStartDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 7*i)
		weekEndDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, 7*(i+1)-1)

		// Jika tanggal 28 termasuk dalam minggu ini
		if weekStartDate.Day() <= 28 && weekEndDate.Day() >= 28 {
			weekEndDate = time.Date(year, month, 28, 23, 59, 59, 999999999, time.UTC)
		}

		if i == weeksInMonth-1 {
			lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
			weekEndDate = time.Date(year, month, lastDayOfMonth, 23, 59, 59, 999999999, time.UTC)
		}
		log.Printf("Week %d: StartDate: %s, EndDate: %s\n", i+1, weekStartDate, weekEndDate)

		// Filter data yang berada dalam rentang waktu minggu ini
		filteredData := FilterDataByDate(data, weekStartDate, weekEndDate)

		// Hitung jumlah data trash_type dan scala_type
		trashCount, scalaCount := CountTrashAndScalaTypes(filteredData)

		// Set nilai WeeklyStats
		weeklyStats[i].Week = i + 1
		weeklyStats[i].Trash = trashCount
		weeklyStats[i].Scala = scalaCount
	}

	return weeklyStats
}

func MapTrashIncomeStats(totalIncomeThisMonth, totalIncomeLastMonth int) TrashIncomeStats {
	return TrashIncomeStats{
		TotalIncomeThisMonth: totalIncomeThisMonth,
		TotalIncomeLastMonth: totalIncomeLastMonth,
	}
}

func MapToGetCountIncome(totalThisMonth int, totalLastMonth int) GetCountTrashExchangeIncome {
	var persentasePerubahanIncome float64
	if totalLastMonth > 0 {
		persentasePerubahanIncome = float64(totalThisMonth-totalLastMonth) / float64(totalLastMonth) * 100
	} else {
		persentasePerubahanIncome = 0
	}

	var statusIncome string
	if persentasePerubahanIncome > 0 {
		statusIncome = "naik"
	} else if persentasePerubahanIncome < 0 {
		statusIncome = "turun"
	} else {
		statusIncome = "tetap"
	}

	persentasePerubahanInt := int(math.Round(persentasePerubahanIncome))

	// Buat map hasil untuk total income
	result := GetCountTrashExchangeIncome{
		TotalIncome: totalThisMonth,
		Persentase:  fmt.Sprintf("%d", persentasePerubahanInt),
		Status:      statusIncome,
	}

	return result
}