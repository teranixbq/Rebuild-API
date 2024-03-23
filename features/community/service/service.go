package service

import (
	"errors"
	"mime/multipart"
	"recything/features/community/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type communityService struct {
	communityRepository entity.CommunityRepositoryInterface
}

func NewCommunityService(community entity.CommunityRepositoryInterface) entity.CommunityServiceInterface {
	return &communityService{
		communityRepository: community,
	}
}

// CreateCommunity implements entity.CommunityServiceInterface.
func (cs *communityService) CreateCommunity(image *multipart.FileHeader, data entity.CommunityCore) error {
	errEmpty := validation.CheckDataEmpty(data.Name, data.Description, data.Location, data.MaxMembers)
	if errEmpty != nil {
		return errEmpty
	}

	// Mengubah huruf pertama di Name menjadi huruf besar
	titleCase := cases.Title(language.Indonesian)
	data.Name = titleCase.String(data.Name)

	// Mengubah huruf pertama di Location menjadi huruf besar
	data.Location = titleCase.String(data.Location)

	_, err := cs.communityRepository.GetByName(data.Name)
	if err == nil {
		return errors.New("nama community sudah digunakan")
	}

	errCreate := cs.communityRepository.CreateCommunity(image, data)
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// DeleteCommunityById implements entity.CommunityServiceInterface.
func (cs *communityService) DeleteCommunityById(id string) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	err := cs.communityRepository.DeleteCommunityById(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllCommunity implements entity.CommunityServiceInterface.
func (cs *communityService) GetAllCommunity(page, limit, search string) ([]entity.CommunityCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateTypePaginationParameter(limit, page)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	pageValid, limitValid := validation.ValidateCountLimitAndPage(pageInt, limitInt)

	dropPointCores, pageInfo, count, err := cs.communityRepository.GetAllCommunity(pageValid, limitValid, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return dropPointCores, pageInfo, count, nil
}

// GetCommunityById implements entity.CommunityServiceInterface.
func (cs *communityService) GetCommunityById(id string) (entity.CommunityCore, error) {
	if id == "" {
		return entity.CommunityCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	idCommunity, err := cs.communityRepository.GetCommunityById(id)
	if err != nil {
		return entity.CommunityCore{}, err
	}

	return idCommunity, err
}

// UpdateCommunityById implements entity.CommunityServiceInterface.
func (cs *communityService) UpdateCommunityById(id string, image *multipart.FileHeader, data entity.CommunityCore) error {
	errEmpty := validation.CheckDataEmpty(data.Name, data.Description, data.Location, data.MaxMembers)
	if errEmpty != nil {
		return errEmpty
	}

	titleCase := cases.Title(language.Indonesian)
	data.Name = titleCase.String(data.Name)
	data.Location = titleCase.String(data.Location)

	_, err := cs.communityRepository.GetByName(data.Name)
	if err == nil {
		return errors.New("nama community sudah digunakan")
	}

	err = cs.communityRepository.UpdateCommunityById(id, image, data)
	if err != nil {
		return err
	}

	return nil
}

// Event

// CreateEvent implements entity.CommunityServiceInterface.
func (cs *communityService) CreateEvent(communityId string, eventInput entity.CommunityEventCore, image *multipart.FileHeader) error {
	_, errFind := cs.communityRepository.GetCommunityById(communityId)
	if errFind != nil {
		return errFind
	}

	errEmpty := validation.CheckDataEmpty(eventInput.Title, eventInput.Description, eventInput.Date,
		eventInput.Quota, eventInput.Location, eventInput.MapLink, eventInput.FormLink)
	if errEmpty != nil {
		return errEmpty
	}

	if _, parseErr := time.Parse("2006/01/02", eventInput.Date); parseErr != nil {
		return errors.New("error, tanggal harus dalam format 'yyyy/mm/dd'")
	}

	if image != nil && image.Size > 5*1024*1024 {
		return errors.New("ukuran file tidak boleh lebih dari 5 MB")
	}

	errInsert := cs.communityRepository.CreateEvent(communityId, eventInput, image)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

// DeleteEvent implements entity.CommunityServiceInterface.
func (cs *communityService) DeleteEvent(communityId string, eventId string) error {
	if eventId == "" {
		return errors.New("id event tidak ditemukan")
	}

	errEvent := cs.communityRepository.DeleteEvent(communityId, eventId)
	if errEvent != nil {
		return errEvent
	}

	return nil
}

// ReadAllEvent implements entity.CommunityServiceInterface.
func (cs *communityService) ReadAllEvent(status string, page string, limit string, search string, communityId string) ([]entity.CommunityEventCore, pagination.PageInfo, pagination.CountEventInfo, error) {
	pageInt, limitInt, validationErr := validation.ValidateTypePaginationParameter(limit, page)

	if validationErr != nil {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, validationErr
	}

	pageValid, limitValid := validation.ValidateCountLimitAndPage(pageInt, limitInt)

	validStatus := map[string]bool{
		"belum berjalan": true,
		"berjalan":       true,
		"selesai":        true,
	}

	if _, ok := validStatus[status]; status != "" && !ok {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, errors.New("status tidak valid")
	}

	data, paginationInfo, count, err := cs.communityRepository.ReadAllEvent(status, pageValid, limitValid, search, communityId)
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, err
	}

	return data, paginationInfo, count, nil	
}

// ReadEvent implements entity.CommunityServiceInterface.
func (cs *communityService) ReadEvent(communityId string, eventId string) (entity.CommunityEventCore, error) {
	if eventId == "" {
		return entity.CommunityEventCore{}, errors.New("event tidak ditemukan")
	}

	eventData, err := cs.communityRepository.ReadEvent(communityId, eventId)
	if err != nil {
		return entity.CommunityEventCore{}, err
	}

	return eventData, nil
}

// UpdateEvent implements entity.CommunityServiceInterface.
func (cs *communityService) UpdateEvent(communityId string, eventId string, eventInput entity.CommunityEventCore, image *multipart.FileHeader) error {
	if eventId == "" {
		return errors.New("event tidak ditemukan")
	}

	dataStatus, errEqual := validation.CheckEqualData(eventInput.Status, constanta.STATUS_EVENT)
	if errEqual != nil {
		return errors.New("error : status input tidak valid")
	}
	errEmpty := validation.CheckDataEmpty(eventInput.Title, eventInput.Description, eventInput.Date,
		eventInput.Quota, eventInput.Location, eventInput.MapLink, eventInput.FormLink)
	if errEmpty != nil {
		return errEmpty
	}

	if _, parseErr := time.Parse("2006/01/02", eventInput.Date); parseErr != nil {
		return errors.New("error, tanggal harus dalam format 'yyyy/mm/dd'")
	}

	if image != nil && image.Size > 5*1024*1024 {
		return errors.New("ukuran file tidak boleh lebih dari 5 MB")
	}

	eventInput.Status = dataStatus
	errInsert := cs.communityRepository.UpdateEvent(communityId, eventId, eventInput, image)
	if errInsert != nil {
		return errInsert
	}

	return nil
}
