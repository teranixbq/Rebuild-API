package handler

import (
	"net/http"
	"recything/features/voucher/dto/request"
	"recything/features/voucher/dto/response"
	"recything/features/voucher/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type voucherHandler struct {
	VoucherService entity.VoucherServiceInterface
}

func NewVoucherHandler(voucher entity.VoucherServiceInterface) *voucherHandler {
	return &voucherHandler{VoucherService: voucher}
}

func (vh *voucherHandler) CreateVoucher(e echo.Context) error {
	input := request.VoucherRequest{}
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err := helper.BindFormData(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_EMPTY_FILE))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal upload file"))
	}

	request := request.RequestVoucherToCoreVoucher(input)
	err = vh.VoucherService.Create(image, request)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_MESSAGE...) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SUCCESS_CREATE_DATA))
}

func (vh *voucherHandler) GetAllVoucher(e echo.Context) error {
	idUser, _, errExtract := jwt.ExtractToken(e)

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	search := e.QueryParam("search")

	result, pagination, count, err := vh.VoucherService.GetAll(page, limit, search)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_MESSAGE...) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := response.ListCoreVoucherToCoreVoucher(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount(constanta.SUCCESS_GET_DATA, response, pagination, count))
}

func (vh *voucherHandler) GetVoucherById(e echo.Context) error {
	idUser, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	id := e.Param("id")
	result, err := vh.VoucherService.GetById(id)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreVoucherToResponVoucher(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, response))
}

func (vh *voucherHandler) UpdateVoucher(e echo.Context) error {
	input := request.VoucherRequest{}

	id := e.Param("id")

	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err := helper.BindFormData(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, _ := e.FormFile("image")

	request := request.RequestVoucherToCoreVoucher(input)
	err = vh.VoucherService.UpdateData(id, image, request)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_DATA_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		if helper.HttpResponseCondition(err, constanta.ERROR_MESSAGE...) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil mengupdate data"))
}

func (vh *voucherHandler) DeleteVoucherById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	err := vh.VoucherService.DeleteData(id)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_DATA_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_DELETE_DATA))
}

func (vh *voucherHandler) CreateExchangeVoucher(e echo.Context) error {
	input := request.VoucherExchangeRequest{}

	idUser, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(errExtract.Error()))
	}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.RequestVoucherExchangeToCoreVoucherExchange(input)
	err := vh.VoucherService.CreateExchangeVoucher(idUser, request)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil menukarkan voucher"))
}

func (vh *voucherHandler) GetAllExchange(e echo.Context) error {
	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	search := e.QueryParam("search")
	filter := e.QueryParam("filter")

	idUser, role, errExtract := jwt.ExtractToken(e)

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))

	}

	result, pagination, counts, errGet := vh.VoucherService.GetAllExchange(page, limit, search, filter)
	if errGet != nil {
		if strings.Contains(errGet.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		if strings.Contains(errGet.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errGet.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errGet.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := response.ListCoreExchangeVoucherToExchangeVoucheResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCountAll("Berhasil mendapatkan data", response, pagination, counts))

}

func (vh *voucherHandler) GetByIdExchange(e echo.Context) error {
	idUser, role, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))

	}

	id := e.Param("id")
	result, err := vh.VoucherService.GetByIdExchange(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreExchangeVoucherToExchangeVoucheResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, response))
}

func (vh *voucherHandler) UpdateStatusExchange(e echo.Context) error {

	input := request.ExchangeVoucherRequest{}
	id := e.Param("id")

	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err = helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	err = vh.VoucherService.UpdateStatusExchange(id, input.Status)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil memperbarui status"))
}
