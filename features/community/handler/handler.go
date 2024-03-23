package handler

import (
	"net/http"
	"recything/features/community/dto/request"
	"recything/features/community/dto/response"
	"recything/features/community/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type communityHandler struct {
	communityService entity.CommunityServiceInterface
}

func NewCommunityHandler(community entity.CommunityServiceInterface) *communityHandler {
	return &communityHandler{
		communityService: community,
	}
}

func (ch *communityHandler) CreateCommunity(e echo.Context) error {

	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.CommunityRequest{}
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

	request := request.RequestCommunityToCoreCommunity(input)
	err = ch.communityService.CreateCommunity(image, request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SUCCESS_CREATE_DATA))
}

func (ch *communityHandler) GetAllCommunity(e echo.Context) error {
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

	result, pagination, count, err := ch.communityService.GetAllCommunity(page, limit, search)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_INVALID_TYPE) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := response.ListCoreCommunityToResponseCommunity(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount(constanta.SUCCESS_GET_DATA, response, pagination, count))
}

func (ch *communityHandler) GetCommunityById(e echo.Context) error {
	idUser, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	id := e.Param("id")
	result, err := ch.communityService.GetCommunityById(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreCommunityToResponCommunityForDetails(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, response))
}

func (ch *communityHandler) DeleteCommunityById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	err := ch.communityService.DeleteCommunityById(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_DELETE_DATA))
}

func (ch *communityHandler) UpdateCommunityById(e echo.Context) error {
	input := request.CommunityRequest{}

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

	id := e.Param("id")
	image, _ := e.FormFile("image")

	request := request.RequestCommunityToCoreCommunity(input)
	err = ch.communityService.UpdateCommunityById(id, image, request)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil update data"))
}

// Event

func (ch *communityHandler) CreateEvent(e echo.Context) error {
	Id, role, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}

	if role != "admin" && role != "super_admin" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("akses ditolak"))
	}

	idParams := e.Param("idkomunitas")

	newEvent := request.EventRequest{}
	err := e.Bind(&newEvent)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse("tidak ada file yang di upload"))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal upload file"))
	}

	eventInput := request.EventRequestToEventCore(newEvent)
	errCreate := ch.communityService.CreateEvent(idParams, eventInput, image)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(errCreate.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errCreate.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil menambahkan event"))
}

func (ch *communityHandler) DeleteEvent(e echo.Context) error {
	Id, role, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}

	if role != "admin" && role != "super_admin" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("akses ditolak"))
	}

	idKom := e.Param("idkomunitas")
	idEve := e.Param("idevent")

	errDelete := ch.communityService.DeleteEvent(idKom, idEve)
	if errDelete != nil {
		if strings.Contains(errDelete.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(errDelete.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errDelete.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errDelete.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus event"))
}

func (ch *communityHandler) ReadAllEvent(e echo.Context) error {

	filter := e.QueryParam("filter")
	search := e.QueryParam("search")
	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	idKom := e.Param("idkomunitas")

	Id, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_ID_INVALID))
	}

	eventData, paginationInfo, count, err := ch.communityService.ReadAllEvent(filter, page, limit, search, idKom)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(eventData) == 0 {
		return e.JSON(http.StatusNotFound, helper.SuccessResponse(constanta.ERROR_DATA_NOT_FOUND))
	}

	var eventResponse = response.ListEventCoreToListEventRessponse(eventData)
	return e.JSON(http.StatusOK, helper.SuccessWithPaginationAndCount("berhasil mendapatkan semua event", eventResponse, paginationInfo, count))
}

func (ch *communityHandler) ReadEvent(e echo.Context) error {
	idKom := e.Param("idkomunitas")
	idEve := e.Param("idevent")

	Id, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_ID_INVALID))
	}

	eventData, errRead := ch.communityService.ReadEvent(idKom, idEve)
	if errRead != nil {
		if strings.Contains(errRead.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(errRead.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errRead.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errRead.Error()))
	}

	var eventResponse = response.EventCoreToEventResponseDetail(eventData)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan event", eventResponse))
}

func (ch *communityHandler) UpdateEvent(e echo.Context) error {
	Id, role, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}

	if role != "admin" && role != "super_admin" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("akses ditolak"))
	}

	idKom := e.Param("idkomunitas")
	idEve := e.Param("idevent")

	updateData := request.EventRequest{}
	errBind := e.Bind(&updateData)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	image, _ := e.FormFile("image")

	eventInput := request.EventRequestToEventCore(updateData)
	errUpdate := ch.communityService.UpdateEvent(idKom, idEve, eventInput, image)
	if errUpdate != nil {
		if strings.Contains(errUpdate.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(errUpdate.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errUpdate.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errUpdate.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil update event"))
}
