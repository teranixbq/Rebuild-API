package handler

import (
	"net/http"
	"recything/features/drop-point/dto/request"
	"recything/features/drop-point/dto/response"
	"recything/features/drop-point/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type dropPointHandler struct {
	dropPointService entity.DropPointServiceInterface
}

func NewDropPointHandler(dropPoint entity.DropPointServiceInterface) *dropPointHandler {
	return &dropPointHandler{
		dropPointService: dropPoint,
	}
}

func (dph *dropPointHandler) CreateDropPoint(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.DropPointRequest{}

	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.DropPointRequestToCoreDropPoint(input)

	errCreate := dph.dropPointService.CreateDropPoint(request)
	if errCreate != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SUCCESS_CREATE_DATA))
}

func (dph *dropPointHandler) GetAllDropPoint(e echo.Context) error {
	idUser, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	dropPoints, paginationInfo, count, err := dph.dropPointService.GetAllDropPoint(page, limit, search)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(dropPoints) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}
	   for _, dropPoint := range dropPoints {
        dropPoint.Schedule = helper.SortByDay(dropPoint.Schedule)
    }

	
	
    response := response.ListCoreDropPointToDropPointResponse(dropPoints)

	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("berhasil mendapatkan data", response, paginationInfo, count))

}



func (dph *dropPointHandler) GetDropPointById(e echo.Context) error {
	idUser, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	idParams := e.Param("id")
	result, err := dph.dropPointService.GetDropPointById(idParams)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	var reportResponse = response.CoreDropPointToDropPointResponse(result)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, reportResponse))
}

func (dph *dropPointHandler) UpdateDropPoint(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.DropPointRequest{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.DropPointRequestToCoreDropPoint(input)

	dropPointId := e.Param("id")
	errUpdate := dph.dropPointService.UpdateDropPointById(dropPointId, request)
	if errUpdate != nil {
		if strings.Contains(errUpdate.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errUpdate.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil melakukan update data"))

}

func (dph *dropPointHandler) DeleteDropPoint(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	dropPointId := e.Param("id")
	err = dph.dropPointService.DeleteDropPointById(dropPointId)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_DATA_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_DELETE_DATA))
}
