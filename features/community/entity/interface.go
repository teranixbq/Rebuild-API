package entity

import (
	"mime/multipart"
	"recything/utils/pagination"
)

type CommunityRepositoryInterface interface {
	CreateCommunity(image *multipart.FileHeader, data CommunityCore) error
	GetAllCommunity(page, limit int, search string) ([]CommunityCore, pagination.PageInfo, int, error)
	GetCommunityById(id string) (CommunityCore, error)
	GetByName(name string) (CommunityCore, error)
	UpdateCommunityById(id string, image *multipart.FileHeader, data CommunityCore) error
	DeleteCommunityById(id string) error

	//Event
	CreateEvent(communityId string, eventInput CommunityEventCore, image *multipart.FileHeader) error
	ReadAllEvent(status string, page, limit int, search string, communityId string) ([]CommunityEventCore, pagination.PageInfo, pagination.CountEventInfo, error)
	ReadEvent(communityId string,eventId string) (CommunityEventCore, error)
	UpdateEvent(communityId string, eventId string, eventInput CommunityEventCore, image *multipart.FileHeader) error
	DeleteEvent(communityId string, eventId string) error
	GetCountByStatus(status, communityid string, search string) (int64, error)
}

type CommunityServiceInterface interface {
	CreateCommunity(image *multipart.FileHeader, data CommunityCore) error
	GetAllCommunity(page, limit, search string) ([]CommunityCore, pagination.PageInfo, int, error)
	GetCommunityById(id string) (CommunityCore, error)
	UpdateCommunityById(id string, image *multipart.FileHeader, data CommunityCore) error
	DeleteCommunityById(id string) error

	//Event
	CreateEvent(communityId string, eventInput CommunityEventCore, image *multipart.FileHeader) error
	ReadAllEvent(status string, page, limit string, search string, communityId string) ([]CommunityEventCore, pagination.PageInfo, pagination.CountEventInfo, error)
	ReadEvent(communityId string,eventId string) (CommunityEventCore, error)
	UpdateEvent(communityId string, eventId string, eventInput CommunityEventCore, image *multipart.FileHeader) error
	DeleteEvent(communityId string, eventId string) error
}
