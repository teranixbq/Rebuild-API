package entity

import (
	"recything/features/faq/model"
)

func FaqsCoreToFaqsModel(data FaqCore) model.Faq {
	return model.Faq{
		Title:             data.Title,
		Description:       data.Description,
 }
}

func ListFaqCoreToFaqModel(data []FaqCore) []model.Faq {
	listFaq := []model.Faq{}
	for _, faq := range data {
		faqModel := FaqsCoreToFaqsModel(faq)
		listFaq = append(listFaq, faqModel)
	}
	return listFaq
}

func FaqsModelToFaqsCore(data model.Faq) FaqCore {
	return FaqCore{
		Id:                data.Id,
		Title:             data.Title,
		Description:       data.Description,
	}
}

func ListFaqModelToFaqCore(data []model.Faq) []FaqCore {
	listFaq := []FaqCore{}
	for _, faq := range data {
		faqModel := FaqsModelToFaqsCore(faq)
		listFaq = append(listFaq, faqModel)
	}
	return listFaq
}
