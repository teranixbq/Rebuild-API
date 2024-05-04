package entity

import (
	"recything/utils/helper"
	"recything/utils/pagination"
)

type RecybotRepositoryInterface interface {
	Create(recybot RecybotCore) (RecybotCore, error)
	Update(idData string, data RecybotCore) (RecybotCore, error)
	Delete(idData string) error
	GetAll() ([]RecybotCore, error)
	GetById(idData string) (RecybotCore, error)
	FindAll(page, limit int, filter, search string) ([]RecybotCore, pagination.PageInfo, helper.CountPrompt, error)
	GetCountAllData(search, filter string) (helper.CountPrompt, error)
	InsertHistory(userId, answer, question string) error
	GetAllHistory(userId string) (RecybbotHistories, error)
}

type RecybotServiceInterface interface {
	CreateData(recybot RecybotCore) (RecybotCore, error)
	UpdateData(idData string, data RecybotCore) (RecybotCore, error)
	DeleteData(idData string) error
	GetById(idData string) (RecybotCore, error)
	GetPrompt(userId, question string) (string, error)
	FindAllData(filter, search, page, limit string) ([]RecybotCore, pagination.PageInfo, helper.CountPrompt, error)
}
