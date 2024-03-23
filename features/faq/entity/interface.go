package entity

type FaqRepositoryInterface interface {
	GetFaqsById(id uint) (FaqCore, error)
	GetFaqs() ([]FaqCore, error)
}

type FaqServiceInterface interface {
	GetFaqsById(id uint) (FaqCore, error)
	GetFaqs() ([]FaqCore, error)
}
