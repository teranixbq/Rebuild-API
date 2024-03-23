package entity

import "recything/features/trash_category/model"

func CoreTrashCategoryToModelTrashCategory(trash TrashCategoryCore) model.TrashCategory {
	return model.TrashCategory{
		ID:        trash.ID,
		TrashType: trash.TrashType,
		Point:     trash.Point,
		Unit:      trash.Unit,
		CreatedAt: trash.CreatedAt,
		UpdatedAt: trash.UpdatedAt,
	}
}

func ModelTrashCategoryToCoreTrashCategory(trash model.TrashCategory) TrashCategoryCore {
	return TrashCategoryCore{
		ID:        trash.ID,
		TrashType: trash.TrashType,
		Point:     trash.Point,
		Unit:      trash.Unit,
		CreatedAt: trash.CreatedAt,
		UpdatedAt: trash.UpdatedAt,
	}
}

func ListCoreTrashCategoryToModelTrashCategory(trash []TrashCategoryCore) []model.TrashCategory {
	list := []model.TrashCategory{}
	for _, v := range trash {
		result := CoreTrashCategoryToModelTrashCategory(v)
		list = append(list, result)
	}
	return list
}

func ListModelTrashCategoryToCoreTrashCategory(trash []model.TrashCategory) []TrashCategoryCore {
	list := []TrashCategoryCore{}
	for _, v := range trash {
		result := ModelTrashCategoryToCoreTrashCategory(v)
		list = append(list, result)
	}
	return list
}
