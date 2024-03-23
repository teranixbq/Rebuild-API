package repository

import (
	"errors"
	"recything/features/faq/entity"
	"recything/features/faq/model"
	"recything/utils/constanta"

	"gorm.io/gorm"
)

type faqRepository struct {
	db *gorm.DB
}

func NewFaqRepository(db *gorm.DB) entity.FaqRepositoryInterface {
	return &faqRepository{
		db: db,
	}
}

func (fr *faqRepository) GetFaqsById(id uint) (entity.FaqCore, error) {
	dataFaqs := model.Faq{}

	tx := fr.db.Where("id = ?", id).First(&dataFaqs)
	if tx.Error != nil {
		if tx.RowsAffected == 0 {
			return entity.FaqCore{}, errors.New(constanta.ERROR_DATA_ID)
		}
		return entity.FaqCore{}, tx.Error
	}

	dataResponse := entity.FaqsModelToFaqsCore(dataFaqs)
	return dataResponse, nil
}

func (fr *faqRepository) GetFaqs() ([]entity.FaqCore, error) {
	dataFaqs := []model.Faq{}

	tx := fr.db.Find(&dataFaqs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataResponse := entity.ListFaqModelToFaqCore(dataFaqs)
	return dataResponse, nil
}
