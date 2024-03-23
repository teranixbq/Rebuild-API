package handler

import (
	"net/http"
	"recything/features/trash_category/dto/request"
	"recything/features/trash_category/dto/response"
	"recything/features/trash_category/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type trashCategoryHandler struct {
	trashCategory entity.TrashCategoryServiceInterface
}

func NewTrashCategoryHandler(trashCategory entity.TrashCategoryServiceInterface) *trashCategoryHandler {
	return &trashCategoryHandler{trashCategory: trashCategory}
}

func (tc *trashCategoryHandler) CreateCategory(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	requestCategory := request.TrashCategory{}
	err = helper.DecodeJSON(e, &requestCategory)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	input := request.RequestTrashCategoryToCoreTrashCategory(requestCategory)
	err = tc.trashCategory.CreateCategory(input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusCreated, helper.SuccessResponse("Berhasil menambahkan kategori sampah"))
}

func (tc *trashCategoryHandler) GetAllCategory(e echo.Context) error {

	id, _, err := jwt.ExtractToken(e)
	if id == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	search := e.QueryParam("search")

	result, pagnation, count, err := tc.trashCategory.GetAllCategory(page, limit, search)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_INVALID_TYPE) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("Belum ada kategori sampah"))
	}

	response := response.ListCoreTrashCategoryToReponseTrashCategory(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("Berhasil mendapatkan seluruh kategori sampah", response, pagnation, count))
}

func (tc *trashCategoryHandler) GetAllCategoriesFetch(e echo.Context) error {

	idUser, _, err := jwt.ExtractToken(e)
	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	result, err := tc.trashCategory.FindAllFetch()
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_INVALID_TYPE) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("Belum ada kategori sampah"))
	}

	response := response.ListCoreTrashCategoryToReponseTrashCategoryCategoriesList(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan seluruh kategori sampah", response))
}

func (tc *trashCategoryHandler) GetById(e echo.Context) error {

	id, _, err := jwt.ExtractToken(e)
	if id == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	idTrash := e.Param("id")
	result, err := tc.trashCategory.GetById(idTrash)

	if err != nil {
		if strings.Contains(constanta.ERROR_DATA_ID, err.Error()) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreTrashCategoryToReponseTrashCategory(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan detail kategori sampah", response))
}

func (tc *trashCategoryHandler) DeleteById(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	err = tc.trashCategory.DeleteCategory(id)
	if err != nil {
		if strings.Contains(constanta.ERROR_DATA_ID, err.Error()) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil menghapus kategori"))
}

func (tc *trashCategoryHandler) UpdateCategory(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	id := e.Param("id")
	requestCategory := request.TrashCategory{}
	err = helper.DecodeJSON(e, &requestCategory)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := request.RequestTrashCategoryToCoreTrashCategory(requestCategory)
	result, err := tc.trashCategory.UpdateCategory(id, input)
	if err != nil {
		if strings.Contains(constanta.ERROR_DATA_ID, err.Error()) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreTrashCategoryToReponseTrashCategory(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mengupdate kategori sampah", response))
}
