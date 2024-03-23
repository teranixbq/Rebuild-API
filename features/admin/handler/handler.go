package handler

import (
	"net/http"
	"recything/features/admin/dto/request"
	"recything/features/admin/dto/response"
	"recything/features/admin/entity"
	reportRequest "recything/features/report/dto/request"
	reportDto "recything/features/report/dto/response"
	userDto "recything/features/user/dto/response"
	user "recything/features/user/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService entity.AdminServiceInterface
	UserService  user.UsersUsecaseInterface
}

func NewAdminHandler(as entity.AdminServiceInterface, us user.UsersUsecaseInterface) *AdminHandler {
	return &AdminHandler{
		AdminService: as,
		UserService:  us,
	}
}

// membuat admin, hanya untuk super admin
func (ah *AdminHandler) Create(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.AdminRequest{}
	err = e.Bind(&input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, _ := e.FormFile("image")

	request := request.AdminRequestToAdminCore(input)

	_, err = ah.AdminService.Create(image, request)
	if err != nil {
		
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil membuat data admin"))

}

// login untuk admin dan juga super admin
func (ah *AdminHandler) Login(e echo.Context) error {
	input := request.AdminLogin{}

	err := helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.RequestLoginToAdminCore(input)

	result, token, err := ah.AdminService.FindByEmailANDPassword(request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	response := response.AdminCoreToAdminResponseLogin(result, token)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_LOGIN, response))
}

// mendapatkan semua data admin yang active maupun yang tidak active
func (ah *AdminHandler) GetAll(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	fullName := e.QueryParam("search")

	result, pagnationInfo, count, err := ah.AdminService.GetAll(page, limit, fullName)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("data admin belum ada"))
	}

	response := response.ListAdminCoreToAdminResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("berhasil mengambil semua data admin", response, pagnationInfo, count))

}

// mendapatkan data admin detail lengkap
func (ah *AdminHandler) GetById(e echo.Context) error {

	adminId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	result, err := ah.AdminService.GetById(adminId)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// if len(result.Id) == 0 {
	// 	return e.JSON(http.StatusOK, helper.SuccessResponse("data admin belum ada"))
	// }

	response := response.AdminCoreToAdminResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mengambil data admin", response))
}

// menghapus data admin
func (ah *AdminHandler) Delete(e echo.Context) error {
	adminId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err = ah.AdminService.DeleteById(adminId)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus data admin"))

}

// melakukan pembaruan atau edit data admin
func (ah *AdminHandler) UpdateById(e echo.Context) error {
	adminId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.AdminRequest{}

	err = e.Bind(&input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, _ := e.FormFile("image")
	request := request.AdminRequestToAdminCore(input)
	err = ah.AdminService.UpdateById(image, adminId, request)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))

		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil melakukan pembaruan data admin"))
}

// Manage User
func (ah *AdminHandler) GetAllUser(e echo.Context) error {
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

	result, pagination, count, err := ah.AdminService.GetAllUsers(search, page, limit)
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}
	response := userDto.UsersCoreToResponseManageUsersList(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("berhasil mendapatkan data user", response, pagination, count))

}

func (ah *AdminHandler) GetByIdUsers(e echo.Context) error {
	userId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	UsersData, err := ah.AdminService.GetByIdUsers(userId)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	userResponse := userDto.UsersCoreToResponseDetailManageUsers(UsersData)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan data user", userResponse))
}

func (ah *AdminHandler) DeleteUsers(e echo.Context) error {
	userId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err = ah.AdminService.DeleteUsers(userId)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus data user"))
}

// Manage Reporting
func (ah *AdminHandler) GetByStatusReport(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	status := e.QueryParam("status")
	search := e.QueryParam("search")
	page := e.QueryParam("page")
	limit := e.QueryParam("limit")

	result, paginationInfo, count, err := ah.AdminService.GetAllReport(status, search, page, limit)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_INVALID_TYPE, constanta.ERROR_INVALID_STATUS, constanta.ERROR_LIMIT) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := reportDto.ListReportCoresToReportResponseForDataReporting(result, ah.UserService)
	return e.JSON(http.StatusOK, helper.SuccessWithPaginationAndCount("berhasil mendapatkan data reporting", response, paginationInfo, count))
}

func (ah *AdminHandler) UpdateStatusReport(e echo.Context) error {

	input := reportRequest.UpdateStatusReportRubbish{}
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")

	err = helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	_, err = ah.AdminService.UpdateStatusReport(id, input.Status, input.RejectionDescription)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if helper.HttpResponseCondition(err, constanta.ALREADY, constanta.NO, constanta.MUST) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil memperbarui status"))
}

func (dph *AdminHandler) GetReportById(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	idParams := e.Param("id")
	result, err := dph.AdminService.GetReportById(idParams)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	user, _ := dph.UserService.GetById(result.UserId)
	var reportResponse = reportDto.ReportCoreToReportResponseForDataReportingId(result, user)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan data", reportResponse))
}
