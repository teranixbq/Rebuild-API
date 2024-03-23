package handler

import (
	"net/http"
	"recything/features/article/dto/request"
	"recything/features/article/dto/response"
	"recything/features/article/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	articleService entity.ArticleServiceInterface
}

func NewArticleHandler(article entity.ArticleServiceInterface) *articleHandler {
	return &articleHandler{
		articleService: article,
	}
}

func (article *articleHandler) CreateArticle(e echo.Context) error {
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

	newArticle := request.ArticleRequest{}
	err := e.Bind(&newArticle)
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

	articleInput := request.ArticleRequestToArticleCore(newArticle)
	_, errCreate := article.articleService.CreateArticle(articleInput, image)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse("kategori tidak ditemukan"))

		}
		if strings.Contains(errCreate.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errCreate.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil menambahkan artikel"))
}

func (article *articleHandler) GetAllArticle(e echo.Context) error {
	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	Id, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_ID_INVALID))
	}

	articleData, paginationInfo, count, err := article.articleService.GetAllArticle(page, limit, search,"")
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan artikel"))
	}

	var articleResponse = response.ListArticleCoreToListArticleResponse(articleData)

	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("berhasil mendapatkan semua article", articleResponse, paginationInfo, count))
}


func (article *articleHandler) GetAllArticleUser(e echo.Context) error {
	filter := e.QueryParam("filter")
	search := e.QueryParam("search")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	Id, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_ID_INVALID))
	}

	articleData, paginationInfo, count, err := article.articleService.GetAllArticle(page, limit, search,filter)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if len(articleData) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	var articleResponse = response.ListArticleCoreToListArticleResponse(articleData)

	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("berhasil mendapatkan semua article", articleResponse, paginationInfo, count))
}


func (article *articleHandler) GetSpecificArticle(e echo.Context) error {
	idParams := e.Param("id")

	Id, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_ID_INVALID))
	}

	articleData, err := article.articleService.GetSpecificArticle(idParams)
	if err != nil {
		return e.JSON(http.StatusNotFound, helper.ErrorResponse("gagal membaca data"))
	}

	var articleResponse = response.ArticleCoreToArticleResponse(articleData)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan artikel", articleResponse))
}

func (article *articleHandler) UpdateArticle(e echo.Context) error {
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

	idParams := e.Param("id")

	updatedData := request.ArticleRequest{}
	errBind := e.Bind(&updatedData)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	image, _ := e.FormFile("image")

	articleInput := request.ArticleRequestToArticleCore(updatedData)
	updateArticle, errUpdate := article.articleService.UpdateArticle(idParams, articleInput, image)
	if errUpdate != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errUpdate.Error()))
	}

	articleResponse := response.ArticleCoreToArticleResponse(updateArticle)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("berhasil", articleResponse))
}

func (article *articleHandler) DeleteArticle(e echo.Context) error {
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

	idParams := e.Param("id")

	errDelete := article.articleService.DeleteArticle(idParams)
	if errDelete != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errDelete.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus artikel"))
}

func (article *articleHandler) PostLike(e echo.Context) error{
	Id, _, errId := jwt.ExtractToken(e)
	if errId != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(errId.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_ID_INVALID))
	}

	idParams := e.Param("id")

	err := article.articleService.PostLike(idParams,Id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil melakukan like"))
}

func (article *articleHandler) PostShare(e echo.Context) error{
	Id, _, errId := jwt.ExtractToken(e)
	if errId != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(errId.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_ID_INVALID))
	}

	idParams := e.Param("id")

	err := article.articleService.PostShare(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil melakukan share"))
}

func (article *articleHandler) GetPopularArticle(e echo.Context) error{
	Id, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_ID_INVALID))
	}

	search := e.QueryParam("search")
	
	articleData, errData := article.articleService.GetPopularArticle(search)
	if errData != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan artikel"))
	}

	var articleResponse = response.ListArticleCoreToListArticleResponse(articleData)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan artikel populer", articleResponse))
}