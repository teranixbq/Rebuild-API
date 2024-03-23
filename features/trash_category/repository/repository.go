package repository

import (
	"errors"

	"recything/features/trash_category/entity"
	"recything/features/trash_category/model"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"

	"gorm.io/gorm"
)

type trashCategoryRepository struct {
	db *gorm.DB
}

func NewTrashCategoryRepository(db *gorm.DB) entity.TrashCategoryRepositoryInterface {
	return &trashCategoryRepository{
		db: db,
	}
}

func (tc *trashCategoryRepository) Create(data entity.TrashCategoryCore) error {
	input := entity.CoreTrashCategoryToModelTrashCategory(data)

	tx := tc.db.Create(&input)
	if tx.Error != nil {
		if validation.IsDuplicateError(tx.Error) {
			return errors.New(constanta.ERROR_DATA_EXIST)
		}
		return tx.Error
	}
	return nil
}

func (tc *trashCategoryRepository) FindAll(page, limit int, search string) ([]entity.TrashCategoryCore, pagination.PageInfo, int, error) {
	dataTrashCategories := []model.TrashCategory{}
	offsetInt := (page - 1) * limit

	totalCount, err := tc.GetCount(search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	paginationQuery := tc.db.Limit(limit).Offset(offsetInt)
	if search == "" {
		tx := paginationQuery.Find(&dataTrashCategories)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if search != "" {
		tx := paginationQuery.Where("trash_type LIKE ?", "%"+search+"%").Find(&dataTrashCategories)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	result := entity.ListModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
	paginationInfo := pagination.CalculateData(totalCount, limit, page)
	return result, paginationInfo, totalCount, nil
}

func (tc *trashCategoryRepository) GetCount(search string) (int, error) {
	var totalCount int64
	model := tc.db.Model(&model.TrashCategory{})
	if search == "" {
		tx := model.Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	if search != "" {
		tx := model.Where("trash_type LIKE ?", "%"+search+"%").Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}

	}
	return int(totalCount), nil
}

func (tc *trashCategoryRepository) FindAllFetch() ([]entity.TrashCategoryCore, error) {
	dataTrashCategories := []model.TrashCategory{}
	
		tx := tc.db.Find(&dataTrashCategories)
		if tx.Error != nil {
			return []entity.TrashCategoryCore{},  tx.Error
		}
	
	result := entity.ListModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
	return result, nil
}

func (tc *trashCategoryRepository) GetById(idTrash string) (entity.TrashCategoryCore, error) {

	dataTrashCategories := model.TrashCategory{}
	tx := tc.db.Where("id = ?", idTrash).First(&dataTrashCategories)
	if tx.Error != nil {

		if tx.RowsAffected == 0 {
			return entity.TrashCategoryCore{}, errors.New(constanta.ERROR_DATA_ID)
		}

		return entity.TrashCategoryCore{}, tx.Error
	}

	result := entity.ModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
	return result, nil
}

func (tc *trashCategoryRepository) Update(idTrash string, data entity.TrashCategoryCore) (entity.TrashCategoryCore, error) {
	dataTrashCategories := entity.CoreTrashCategoryToModelTrashCategory(data)

	tx := tc.db.Where("id = ?", idTrash).Updates(&dataTrashCategories)
	if tx.Error != nil {
		return entity.TrashCategoryCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.TrashCategoryCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.ModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
	return result, nil
}

func (tc *trashCategoryRepository) Delete(idTrash string) error {
	data := model.TrashCategory{}

	tx := tc.db.Where("id = ?", idTrash).Delete(&data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}

// GetByType implements entity.TrashCategoryRepositoryInterface.
func (tc *trashCategoryRepository) GetByType(trashType string) (entity.TrashCategoryCore, error) {
	dataTrashCategories := model.TrashCategory{}
	tx := tc.db.Where("trash_type = ?", trashType).First(&dataTrashCategories)

	if tx.RowsAffected == 0 {
		return entity.TrashCategoryCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	if tx.Error != nil {
		return entity.TrashCategoryCore{}, tx.Error
	}

	result := entity.ModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
	return result, nil
}