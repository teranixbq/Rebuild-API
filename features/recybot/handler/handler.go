package handler

import (
	"net/http"
	"recything/features/recybot/dto/request"
	"recything/features/recybot/dto/response"
	"recything/features/recybot/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type recybotHandler struct {
	RecybotService entity.RecybotServiceInterface
}

func NewRecybotHandler(recybot entity.RecybotServiceInterface) *recybotHandler {
	return &recybotHandler{RecybotService: recybot}
}

func (rh *recybotHandler) CreateData(e echo.Context) error {
	input := request.RecybotManageRequest{}

	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err := helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.ManageRequestRecybotToCoreRecybot(input)
	result, err := rh.RecybotService.CreateData(request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreRecybotToResponRecybot(result)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("Berhasil menambahkan data", response))
}

func (rh *recybotHandler) GetAllData(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	filter := e.QueryParam("filter")
	search := e.QueryParam("search")

	result, pagnation, count, err := rh.RecybotService.FindAllData(filter, search, page, limit)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_INVALID_TYPE) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("Belum ada data"))
	}

	response := response.ListCoreRecybotToCoreRecybot(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCountAll("Berhasil mendapatkan seluruh data", response, pagnation, count))
}

func (rh *recybotHandler) GetById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	result, err := rh.RecybotService.GetById(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreRecybotToResponRecybot(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan data", response))
}

func (rh *recybotHandler) DeleteById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	err := rh.RecybotService.DeleteData(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil menghapus data"))
}

func (rh *recybotHandler) UpdateData(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	input := request.RecybotManageRequest{}
	err := helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.ManageRequestRecybotToCoreRecybot(input)
	result, err := rh.RecybotService.UpdateData(id, request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreRecybotToResponRecybot(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mengupdate data", response))
}

func (rh *recybotHandler) RecyBotChat(e echo.Context) error {
	input := request.RecybotRequest{}

	idUser, _, errExtract := jwt.ExtractToken(e)
	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err := helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.RequestRecybotToCoreRecybot(input)
	result, err := rh.RecybotService.GetPrompt(idUser, request.Question)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil melakukan request chat", result))

}
