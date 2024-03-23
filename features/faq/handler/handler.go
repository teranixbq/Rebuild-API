package handler

import (
	"net/http"
	"recything/features/faq/dto/response"
	"recything/features/faq/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type faqHandler struct {
	faqService entity.FaqServiceInterface
}

func NewFaqHandlers(fc entity.FaqServiceInterface) *faqHandler {
	return &faqHandler{
		faqService: fc,
	}
}

func (fc *faqHandler) GetFaqsById(e echo.Context) error {
	faqUser, err := strconv.Atoi(e.Param("id"))

	if err != nil{
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("Tipe ID tidak sesuai"))
	}

	result, errGet := fc.faqService.GetFaqsById(uint(faqUser))
	if errGet != nil {
		if strings.Contains(errGet.Error(), constanta.ERROR_DATA_ID) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(errGet.Error()))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errGet.Error()))
	}

	response := response.FaqsCoreToFaqsResponse(result)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan faq", response))
}

func (fc *faqHandler) GetAllFaqs(e echo.Context) error {
	result, err := fc.faqService.GetFaqs()
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}
	response := response.FaqsCoreToResponseFaqsList(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan data faq", response))

}