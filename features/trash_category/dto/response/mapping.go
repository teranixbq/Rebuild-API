package response

import (
	"recything/features/trash_category/entity"
)

func CoreTrashCategoryToReponseTrashCategory(trash entity.TrashCategoryCore) TrashCategory {
	return TrashCategory{
		ID:        trash.ID,
		TrashType: trash.TrashType,
		Point:     trash.Point,
		Unit:      trash.Unit,
		CreatedAt: trash.CreatedAt,
		UpdatedAt: trash.UpdatedAt,
	}
}

func ListCoreTrashCategoryToReponseTrashCategory(trash []entity.TrashCategoryCore) []TrashCategory {
	list := []TrashCategory{}
	for _, v := range trash {
		result := CoreTrashCategoryToReponseTrashCategory(v)
		list = append(list, result)
	}
	return list
}

func CoreTrashCategoryToReponseTrashCategoriesList(trash entity.TrashCategoryCore) TrashCategoriesList {
	return TrashCategoriesList{
		ID:        trash.ID,
		TrashType: trash.TrashType,
		Point:     trash.Point,
		Unit:      trash.Unit,
	}
}

func ListCoreTrashCategoryToReponseTrashCategoryCategoriesList(trash []entity.TrashCategoryCore) []TrashCategoriesList {
	list := []TrashCategoriesList{}
	for _, v := range trash {
		result := CoreTrashCategoryToReponseTrashCategoriesList(v)
		list = append(list, result)
	}
	return list
}
