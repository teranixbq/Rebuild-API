package request

import (
	"recything/features/article/entity"
)

func ArticleRequestToArticleCore(article ArticleRequest) entity.ArticleCore {
	articleReq := entity.ArticleCore{
		Title:       article.Title,
		Image:       article.Image,
		Content:     article.Content,
		Category_id: article.Category_id,
	}
	category := ListCategoryRequestToCategoryCore(article.Categories)
	articleReq.Categories = category
	return articleReq
}

func CategotyrequestToCategotyCore(category ArticleTrashCategoryRequest) entity.ArticleTrashCategoryCore {
	return entity.ArticleTrashCategoryCore{
		// TrashCategoryID: category.TrashCategoryID,
		Category: category.Category,
	}
}

func ListCategoryRequestToCategoryCore(categories []ArticleTrashCategoryRequest) []entity.ArticleTrashCategoryCore {
	listCategory := []entity.ArticleTrashCategoryCore{}
	for _, v := range categories {
		category := CategotyrequestToCategotyCore(v)
		listCategory = append(listCategory, category)
	}

	return listCategory
}