package response

import (
	"recything/features/article/entity"
)

func CategoryCoreToCategoryResponse(category entity.ArticleTrashCategoryCore) TrashCategoryResponse {
	return TrashCategoryResponse{
		// TrashCategoryID: category.TrashCategoryID,
		Category: category.Category,
	}
}

func ListCategoryCoreToCategoryResponse(categories []entity.ArticleTrashCategoryCore) []TrashCategoryResponse {
	ResponseCategory := []TrashCategoryResponse{}
	for _, v := range categories {
		category := CategoryCoreToCategoryResponse(v)
		ResponseCategory = append(ResponseCategory, category)
	}
	return ResponseCategory
}

func ArticleCoreToArticleResponse(article entity.ArticleCore) ArticleCreateResponse {
	articleResp := ArticleCreateResponse{
		Id:          article.ID,
		Title:       article.Title,
		Image:       article.Image,
		Content:     article.Content,
		Like:        article.Like,
		Share:       article.Share,
		Category_id: article.Category_id,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
	category := ListCategoryCoreToCategoryResponse(article.Categories)
	articleResp.Categories = category
	return articleResp
}

func ListArticleCoreToListArticleResponse(articles []entity.ArticleCore) []ArticleCreateResponse {
	articleResp := []ArticleCreateResponse{}
	for _, article := range articles {
		articlesData := ArticleCoreToArticleResponse(article)
		articleResp = append(articleResp, articlesData)
	}
	return articleResp
}
