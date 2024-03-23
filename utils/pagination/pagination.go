package pagination

import (
	"fmt"

	"math"
)

type PageInfo struct {
	Limit       int `json:"limit"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

type CountDataInfo struct {
	TotalCount         int `json:"total_count"`
	CountPerluDitinjau int `json:"count_pending"`
	CountDiterima      int `json:"count_approved"`
	CountDitolak       int `json:"count_rejected"`
}

type CountEventInfo struct {
	TotalCount         int `json:"total_count"`
	CountBelumBerjalan int `json:"count_pending"`
	CountBerjalan      int `json:"count_active"`
	CountSelesai       int `json:"count_finished"`
}

func MapCountData(totalCount, countPerluDitinjau, countDiterima, countDitolak int64) CountDataInfo {
	return CountDataInfo{
		TotalCount:         int(totalCount),
		CountPerluDitinjau: int(countPerluDitinjau),
		CountDiterima:      int(countDiterima),
		CountDitolak:       int(countDitolak),
	}
}

func MapCountEventData(totalCount, countBelumBerjalan, countBerjalan, countSelesai int64) CountEventInfo {
	return CountEventInfo{
		TotalCount:         int(totalCount),
		CountBelumBerjalan: int(countBelumBerjalan),
		CountBerjalan:      int(countBerjalan),
		CountSelesai:       int(countSelesai),
	}
}

func CalculateData(totalCount, limitInt, pageInt int) PageInfo {
	lastPage := int(math.Ceil(float64(totalCount) / float64(limitInt)))

	paginationInfo := PageInfo{
		Limit:       limitInt,
		CurrentPage: pageInt,
		LastPage:    lastPage,
	}
	return paginationInfo
}

func PaginationMessage(paginationInfo PageInfo, totalData int) string {
	limit := paginationInfo.Limit
	currentPage := paginationInfo.CurrentPage

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	startIndex := (currentPage-1)*limit + 1
	endIndex := min(startIndex+limit-1, totalData)

	responseMessage := fmt.Sprintf("menampilkan data %d sampai %d dari %d data", startIndex, endIndex, totalData)
	return responseMessage
}

func Offset(page, limit int) int {
	offset := (page - 1) * limit
	return offset
}
