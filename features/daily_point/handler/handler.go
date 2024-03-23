package handler

import (
	"net/http"
	"recything/features/daily_point/entity"
	user "recything/features/user/dto/response"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"sort"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type dailyPointHandler struct {
	dailyPointService entity.DailyPointServiceInterface
}

func NewDailyPointHandler(daily entity.DailyPointRepositoryInterface) *dailyPointHandler {
	return &dailyPointHandler{
		dailyPointService: daily,
	}
}

func (daily *dailyPointHandler) PostWeekly(e echo.Context) error {
	err := daily.dailyPointService.PostWeekly()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil menambahkan weekly daily point"))
}

func (daily *dailyPointHandler) DailyClaim(e echo.Context) error {
	Id, _, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}

	err := daily.dailyPointService.DailyClaim(Id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil melakukan daily claim"))
}

func (daily *dailyPointHandler) PointHistory(e echo.Context) error {
	Id, _, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}

	result, err := daily.dailyPointService.GetAllHistoryPoint(Id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))

	}

	sort.Slice(result, func(i, j int) bool {
		timeI, errI := time.Parse(time.RFC3339, result[i]["created_at"].(string))
		if errI != nil {
			return false
		}

		timeJ, errJ := time.Parse(time.RFC3339, result[j]["created_at"].(string))
		if errJ != nil {
			return false
		}

		return timeI.After(timeJ)
	})

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menampilkan seluruh data", result))
}

func (daily *dailyPointHandler) PointHistoryById(e echo.Context) error {
	Id, _, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}

	idTransaction := e.Param("idTransaction")

	result, err := daily.dailyPointService.GetByIdHistoryPoint(Id, idTransaction)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menampilkan data", result))
}

func (daily *dailyPointHandler) ClaimPointHistory(e echo.Context) error {
	Id, _, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}

	result, err := daily.dailyPointService.GetAllClaimedDaily(Id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := user.ListUserDailyPointsCoreToUserDailyPointsResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil menampilkan data", response))
}
