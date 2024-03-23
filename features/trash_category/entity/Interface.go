package entity

import (
	"recything/utils/pagination"
)

type TrashCategoryRepositoryInterface interface {
	Create(data TrashCategoryCore) error
	Update(idTrash string, data TrashCategoryCore) (TrashCategoryCore, error)
	Delete(idTrash string) error
	GetById(idTrash string) (TrashCategoryCore, error)
	FindAll(page, limit int, search string) ([]TrashCategoryCore, pagination.PageInfo, int, error)
	GetCount(search string) (int, error)
	GetByType(trashType string) (TrashCategoryCore, error)
	FindAllFetch() ([]TrashCategoryCore, error)
}

type TrashCategoryServiceInterface interface {
	CreateCategory(data TrashCategoryCore) error
	UpdateCategory(idTrash string, data TrashCategoryCore) (TrashCategoryCore, error)
	DeleteCategory(idTrash string) error
	GetAllCategory(page, limit, search string) ([]TrashCategoryCore, pagination.PageInfo, int, error)
	GetById(idTrash string) (TrashCategoryCore, error)
	FindAllFetch() ([]TrashCategoryCore, error)
}
