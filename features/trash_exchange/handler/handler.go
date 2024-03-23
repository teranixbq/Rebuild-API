package handler

import (
	"net/http"
	"recything/features/trash_exchange/dto/request"
	"recything/features/trash_exchange/dto/response"
	"recything/features/trash_exchange/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type trashExchangeHandler struct {
	trashExchangeService entity.TrashExchangeServiceInterface
}

func NewTrashExchangeHandler(trashExchange entity.TrashExchangeServiceInterface) *trashExchangeHandler {
	return &trashExchangeHandler{
		trashExchangeService: trashExchange,
	}
}

func (teh *trashExchangeHandler) CreateTrashExchange(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.TrashExchangeRequest{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.TrashExchangeRequestToTrashExchangeCore(input)
	_, errCreate := teh.trashExchangeService.CreateTrashExchange(request)
	if errCreate != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SUCCESS_CREATE_DATA))
}

func (dph *trashExchangeHandler) GetAllTrashExchange(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	search := e.QueryParam("search")
	page := e.QueryParam("page")
	limit := e.QueryParam("limit")

	trashExchange, paginationInfo, count, err := dph.trashExchangeService.GetAllTrashExchange(page, limit, search)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(trashExchange) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := response.ListTrashExchangeCoreToTrashExchangeResponse(trashExchange)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("berhasil mendapatkan data", response, paginationInfo, count))

}

func (dph *trashExchangeHandler) GetTrashExchangeById(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	idtrashExchange := e.Param("id")
	result, err := dph.trashExchangeService.GetTrashExchangeById(idtrashExchange)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	var reportResponse = response.TrashExchangeCoreToTrashExchangeResponse(result)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, reportResponse))
}

func (dph *trashExchangeHandler) DeleteTrashExchange(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	idtrashExchange := e.Param("id")
	err = dph.trashExchangeService.DeleteTrashExchangeById(idtrashExchange)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_DATA_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_DELETE_DATA))
}
