package handler

import (
	"fmt"
	"net/http"
	"recything/features/report/dto/request"
	"recything/features/report/dto/response"
	"recything/features/report/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type reportHandler struct {
	reportService entity.ReportServiceInterface
}

func NewReportHandler(report entity.ReportServiceInterface) *reportHandler {
	return &reportHandler{reportService: report}
}

func (report *reportHandler) CreateReport(e echo.Context) error {
	userId, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	newReport := request.ReportRubbishRequest{}
	err = e.Bind(&newReport)

	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	form, err := e.MultipartForm()
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan form multipart"))
	}

	images, ok := form.File["images"]
	if !ok || len(images) == 0 {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("tidak ada file yang di upload"))
	}

	reportInput := request.ReportRequestToReportCore(newReport)
	fmt.Println("handler : ", reportInput.InsidentDate)
	_, errCreate := report.reportService.Create(reportInput, userId, images)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errCreate.Error()))
	}
	// reportResponse := response.ReportCoreToReportResponse(createdReport)
	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil melaporkan"))
}

func (rco *reportHandler) SelectById(e echo.Context) error {
	idParams := e.Param("id")

	result, err := rco.reportService.SelectById(idParams)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	var reportResponse = response.ReportCoreToReportResponse(result)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan data laporan", reportResponse))
}

func (rco *reportHandler) ReadAllReport(e echo.Context) error {
	userId, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	data, err := rco.reportService.ReadAllReport(userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal mendapatkan data laporan"))
	}

	response := response.ListReportCoresToReportResponse(data)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan semua data laporan", response))
}
