package service

import (
	"recything/features/trash_category/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"
)

type trashCategoryService struct {
	trashCategoryRepo entity.TrashCategoryRepositoryInterface
}

func NewTrashCategoryService(trashCategoryRepo entity.TrashCategoryRepositoryInterface) entity.TrashCategoryServiceInterface {
	return &trashCategoryService{
		trashCategoryRepo: trashCategoryRepo,
	}
}

// CreateData implements entity.trashCategoryServiceInterface.
func (tc *trashCategoryService) CreateCategory(data entity.TrashCategoryCore) error {

	errEmpty := validation.CheckDataEmpty(data.Unit, data.TrashType, data.Point)
	if errEmpty != nil {
		return errEmpty
	}

	validUnit, errCheck := validation.CheckEqualData(data.Unit, constanta.Unit)
	if errCheck != nil {
		return errCheck
	}

	data.Unit = validUnit
	err := tc.trashCategoryRepo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (tc *trashCategoryService) GetAllCategory(page, limit, search string) ([]entity.TrashCategoryCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	data, pagnationInfo, count, err := tc.trashCategoryRepo.FindAll(pageInt, limitInt, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}
	return data, pagnationInfo, count, nil
}

func (tc *trashCategoryService) GetById(idTrash string) (entity.TrashCategoryCore, error) {

	result, err := tc.trashCategoryRepo.GetById(idTrash)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete implements entity.trashCategoryServiceInterface.
func (tc *trashCategoryService) DeleteCategory(idTrash string) error {

	err := tc.trashCategoryRepo.Delete(idTrash)
	if err != nil {
		return err
	}
	return nil
}

// UpdateData implements entity.trashCategoryServiceInterface.
func (tc *trashCategoryService) UpdateCategory(idTrash string, data entity.TrashCategoryCore) (entity.TrashCategoryCore, error) {

	errEmpty := validation.CheckDataEmpty(data.TrashType, data.Unit)
	if errEmpty != nil {
		return entity.TrashCategoryCore{}, errEmpty
	}

	result, err := tc.trashCategoryRepo.Update(idTrash, data)
	if err != nil {
		return result, err
	}
	result.ID = idTrash
	return result, nil
}

func (tc *trashCategoryService) FindAllFetch() ([]entity.TrashCategoryCore, error){
	dataTrash,errFind := tc.trashCategoryRepo.FindAllFetch()
	if errFind != nil {
		return []entity.TrashCategoryCore{},errFind
	}

	return dataTrash,nil
}
