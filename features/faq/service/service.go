package service

import (
	"errors"
	"recything/features/faq/entity"
)

type faqService struct {
	faqRepo entity.FaqRepositoryInterface
}

func NewFaqService(faqRepo entity.FaqRepositoryInterface) entity.FaqServiceInterface {
	return &faqService{
		faqRepo: faqRepo,
	}
}

func (fs *faqService) GetFaqsById(id uint) (entity.FaqCore, error) {
	// if id == 0 {
	// 	return entity.FaqCore{}, errors.New(constanta.ERROR_ID_INVALID)
	// }

	dataFaqs, err := fs.faqRepo.GetFaqsById(id)
	if err != nil {
		return entity.FaqCore{}, err
	}
	return dataFaqs, nil
}

func (fs *faqService) GetFaqs() ([]entity.FaqCore, error) {

	data, err := fs.faqRepo.GetFaqs()
	if err != nil {
		return nil, errors.New("")
	}

	return data, nil
}
