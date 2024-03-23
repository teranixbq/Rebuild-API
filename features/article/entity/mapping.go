package entity

import (
	"recything/features/article/model"
	tcm "recything/features/trash_category/model"
)

func CategoryModelToCategoryCore(category tcm.TrashCategory) ArticleTrashCategoryCore {
	return ArticleTrashCategoryCore{
		// TrashCategoryID: category.ID,
		Category:        category.TrashType,
	}
}

func ListCategoryModelToCategoryCore(category []tcm.TrashCategory) []ArticleTrashCategoryCore {
	coreCategory := []ArticleTrashCategoryCore{}
	for _, v := range category {
		category := CategoryModelToCategoryCore(v)
		coreCategory = append(coreCategory, category)
	}
	return coreCategory
}

func CategoryCoreToCategoryModel(category ArticleTrashCategoryCore) tcm.TrashCategory {
	return tcm.TrashCategory{
		// ID:        category.TrashCategoryID,
		TrashType: category.Category,
	}
}

func ListCategoryCoreToCategoryModel(category []ArticleTrashCategoryCore) []tcm.TrashCategory {
	coreCategorys := []tcm.TrashCategory{}
	for _, v := range category {
		categorys := CategoryCoreToCategoryModel(v)
		coreCategorys = append(coreCategorys, categorys)
	}
	return coreCategorys
}

func ArticleCoreToArticleModel(article ArticleCore) model.Article {
	articleModel := model.Article{
		Id:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		Like:      article.Like,
		Share:     article.Share,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	category := ListCategoryCoreToCategoryModel(article.Categories)
	articleModel.Categories = category
	return articleModel
}


func ArticleModelToArticleCore(article model.Article) ArticleCore {
	articleCore := ArticleCore{
		ID:        article.Id,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		Like:      article.Like,
		Share:     article.Share,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	category := ListCategoryModelToCategoryCore(article.Categories)
	articleCore.Categories = category
	return articleCore
}

func ListArticleModelToArticleCore(article []model.Article) []ArticleCore{
	coreArticle := []ArticleCore{}
	for _, v := range article {
		articles := ArticleModelToArticleCore(v)
		coreArticle = append(coreArticle, articles)
	}
	return coreArticle
}
