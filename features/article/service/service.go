package service

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"recything/features/article/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"
)

type articleService struct {
	ArticleRepository entity.ArticleRepositoryInterface
}

func NewArticleService(article entity.ArticleRepositoryInterface) entity.ArticleServiceInterface {
	return &articleService{
		ArticleRepository: article,
	}
}

// DeleteArticle implements entity.ArticleServiceInterface.
func (ac *articleService) DeleteArticle(id string) error {
	if id == "" {
		return errors.New("id artikel tidak ditemukan")
	}

	errArticle := ac.ArticleRepository.DeleteArticle(id)
	if errArticle != nil {
		return errors.New("gagal menghapus artikel " + errArticle.Error())
	}

	return nil
}

// GetSpecificArticle implements entity.ArticleServiceInterface.
func (ac *articleService) GetSpecificArticle(idArticle string) (entity.ArticleCore, error) {
	if idArticle == "" {
		return entity.ArticleCore{}, errors.New("id tidak cocok")
	}

	articleData, err := ac.ArticleRepository.GetSpecificArticle(idArticle)
	if err != nil {
		fmt.Println("service", err)
		return entity.ArticleCore{}, errors.New("gagal membaca data")
	}

	return articleData, nil
}

// UpdateArticle implements entity.ArticleServiceInterface.
func (article *articleService) UpdateArticle(idArticle string, articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {

	if idArticle == "" {
		return entity.ArticleCore{}, errors.New("id tidak ditemukan")
	}

	if articleInput.Title == "" || articleInput.Content == "" {
		return entity.ArticleCore{}, errors.New("artikel tidak boleh kosong")
	}

	if len(articleInput.Category_id) == 0 {
		return entity.ArticleCore{}, errors.New("kategori tidak boleh kosong")
	}

	if image != nil && image.Size > 5*1024*1024 {
		return entity.ArticleCore{}, errors.New("ukuran file tidak boleh lebih dari 5 MB")
	}

	articleUpdate, errinsert := article.ArticleRepository.UpdateArticle(idArticle, articleInput, image)
	if errinsert != nil {
		return entity.ArticleCore{}, errinsert
	}

	return articleUpdate, nil
}

// GetAllArticle implements entity.ArticleServiceInterface.
func (ac *articleService) GetAllArticle(page, limit int, search, filter string) ([]entity.ArticleCore, pagination.PageInfo, int, error) {

	if limit > 10 {
		return nil, pagination.PageInfo{}, 0, errors.New("limit tidak boleh lebih dari 10")
	}

	page, limit = validation.ValidateCountLimitAndPage(page, limit)

	if filter != "" {
		category, errEqual := validation.CheckEqualData(filter, constanta.CATEGORY_ARTICLE)
		if errEqual != nil {
			return []entity.ArticleCore{}, pagination.PageInfo{}, 0, errors.New("error : kategori tidak valid")
		}
		filter = category
	}
	log.Println("ini:", filter)

	article, pageInfo, count, err := ac.ArticleRepository.GetAllArticle(page, limit, search, filter)
	if err != nil {
		return []entity.ArticleCore{}, pagination.PageInfo{}, 0, err
	}

	return article, pageInfo, count, nil
}

// CreateArticle implements entity.ArticleServiceInterface.
func (article *articleService) CreateArticle(articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {

	if articleInput.Title == "" || articleInput.Content == "" {
		return entity.ArticleCore{}, errors.New("judul dan konten artikel tidak boleh kosong")
	}

	if len(articleInput.Category_id) == 0 {
		return entity.ArticleCore{}, errors.New("kategori tidak boleh kosong")
	}

	if image != nil && image.Size > 5*1024*1024 {
		return entity.ArticleCore{}, errors.New("ukuran file tidak boleh lebih dari 5 MB")
	}

	articleCreate, errinsert := article.ArticleRepository.CreateArticle(articleInput, image)
	if errinsert != nil {
		return entity.ArticleCore{}, errinsert
	}

	return articleCreate, nil
}

// PostLike implements entity.ArticleServiceInterface.
func (article *articleService) PostLike(idArticle string, idUser string) error {
	if idArticle == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	if idUser == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	postLike := article.ArticleRepository.PostLike(idArticle, idUser)
	if postLike != nil {
		return postLike
	}
	return nil
}

// PostShare implements entity.ArticleServiceInterface.
func (article *articleService) PostShare(idArticle string) error {
	if idArticle == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	postShare := article.ArticleRepository.PostShare(idArticle)
	if postShare != nil {
		return postShare
	}
	return nil
}

// GetPopularArticle implements entity.ArticleServiceInterface.
func (article *articleService) GetPopularArticle(search string) ([]entity.ArticleCore, error) {
	articleData, err := article.ArticleRepository.GetPopularArticle(search)
	if err != nil {
		return []entity.ArticleCore{}, err
	}

	return articleData, nil
}

