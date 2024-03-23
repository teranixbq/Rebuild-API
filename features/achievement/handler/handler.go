package handler

import (
	"net/http"
	"recything/features/achievement/dto/request"
	"recything/features/achievement/dto/response"
	"recything/features/achievement/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type achievementHandler struct {
	achievementService entity.AchievementServiceInterface
}

func NewAchievementHandler(achievement entity.AchievementServiceInterface) *achievementHandler {
	return &achievementHandler{
		achievementService: achievement,
	}
}

func (ah *achievementHandler) GetAllAchievement(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	result, err := ah.achievementService.GetAllAchievement()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("data achievement belum ada"))
	}

	if role != "" {
		response := response.ListAchievementCoreToAchievementResponse(result)
		return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mengambil semua data", response))
	} else {
		response := response.ListAchievementCoreToAchievementResponseUser(result)
		return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mengambil semua data", response))
	}
}

func (ah *achievementHandler) UpdateById(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	input := request.AchievementRequest{}

	err = helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.AchievementRequestToAchievementCore(input)
	err = ah.achievementService.UpdateById(id, request.TargetPoint)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_DATA_ID) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil melakukan pembaruan data"))
}
