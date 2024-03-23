package repository

import (
	"errors"
	"fmt"
	"mime/multipart"
	"recything/features/community/entity"
	"recything/features/community/model"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/storage"

	"gorm.io/gorm"
)

type communityRepository struct {
	db *gorm.DB
	sb storage.StorageInterface
}

func NewCommunityRepository(db *gorm.DB,sb storage.StorageInterface) entity.CommunityRepositoryInterface {
	return &communityRepository{
		db: db,
		sb: sb,
	}
}

// CreateCommunity implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) CreateCommunity(image *multipart.FileHeader, data entity.CommunityCore) error {
	request := entity.CoreCommunityToModelCommunity(data)

	imageURL, errUpload := cr.sb.Upload(image)
	if errUpload != nil {
		return errUpload
	}
	request.Image = imageURL

	tx := cr.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteCommunityById implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) DeleteCommunityById(id string) error {
	request := model.Community{}

	tx := cr.db.Where("id = ?", id).Delete(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetAllCommunity implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) GetAllCommunity(page int, limit int, search string) ([]entity.CommunityCore, pagination.PageInfo, int, error) {
	communityExchange := []model.Community{}

	offset := (page - 1) * limit
	query := cr.db.Model(&model.Community{})

	if search != "" {
		query = query.Where("name LIKE ? OR location LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Order("members DESC")
	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&communityExchange)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	response := entity.ListModelCommunityToCoreCommunity(communityExchange)
	pageInfo := pagination.CalculateData(int(totalCount), limit, page)
	return response, pageInfo, int(totalCount), nil
}

// GetCommunityById implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) GetCommunityById(id string) (entity.CommunityCore, error) {
	communityExchange := model.Community{}

	tx := cr.db.Where("id = ?", id).First(&communityExchange)
	if tx.Error != nil {
		return entity.CommunityCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.CommunityCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	result := entity.ModelCommunityToCoreCommunity(communityExchange)
	return result, nil
}

// UpdateCommunityById implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) UpdateCommunityById(id string, image *multipart.FileHeader, data entity.CommunityCore) error {
	dataCommunity := model.Community{}

	input := entity.CoreCommunityToModelCommunity(data)

	tx := cr.db.Where("id = ?", id).First(&dataCommunity)
	if tx.Error != nil {
		return tx.Error
	}

	if image != nil {
		imageURL, errUpload := cr.sb.Upload(image)
		if errUpload != nil {
			return errUpload
		}
		input.Image = imageURL
	} else {
		input.Image = dataCommunity.Image
	}

	tx = cr.db.Where("id = ?", id).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (cr *communityRepository) GetByName(name string) (entity.CommunityCore, error) {
	dataCommunity := model.Community{}

	tx := cr.db.Where("name = ?", name).First(&dataCommunity)

	if tx.RowsAffected == 0 {
		return entity.CommunityCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	if tx.Error != nil {
		return entity.CommunityCore{}, tx.Error
	}

	result := entity.ModelCommunityToCoreCommunity(dataCommunity)
	return result, nil
}

// Event

// CreateEvent implements entity.CommunityRepositoryInterface.
func (communityRepo *communityRepository) CreateEvent(communityId string, eventInput entity.CommunityEventCore, image *multipart.FileHeader) error {
	eventData := entity.EventCoreToEventModel(eventInput)

	imageURL, uploadErr := communityRepo.sb.Upload(image)
	if uploadErr != nil {
		return uploadErr
	}

	eventData.Image = imageURL
	eventData.CommunityId = communityId

	tx := communityRepo.db.Create(&eventData)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteEvent implements entity.CommunityRepositoryInterface.
func (communityRepo *communityRepository) DeleteEvent(communityId string, eventId string) error {
	checkId := model.CommunityEvent{}

	tx := communityRepo.db.Where("community_id = ? AND id = ?", communityId, eventId).Delete(&checkId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_RECORD_NOT_FOUND)
	}

	return nil
}

// ReadAllEvent implements entity.CommunityRepositoryInterface.
func (communityRepo *communityRepository) ReadAllEvent(status string, page int, limit int, search string, communityId string) ([]entity.CommunityEventCore, pagination.PageInfo, pagination.CountEventInfo, error) {
	var eventData []model.CommunityEvent

	offset := (page - 1) * limit
	query := communityRepo.db.Model(&model.CommunityEvent{})

	var totalCountWithoutFilter int64
	if err := query.Count(&totalCountWithoutFilter).Error; err != nil {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, err
	}

	totalCountWithFilter := totalCountWithoutFilter

	if search != "" {
		query = query.Where("title LIKE ? OR location LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	tx := query.Where("community_id = ?", communityId).Count(&totalCountWithFilter).Find(&eventData)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, tx.Error
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query = query.Offset(offset).Limit(limit)

	tx = query.Where("community_id = ?", communityId).Find(&eventData)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, tx.Error
	}

	fmt.Println("ini dari read all : ", communityId)
	countBelumBerjalan, err := communityRepo.GetCountByStatus("belum berjalan", communityId, search)
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, err
	}

	countBerjalan, err := communityRepo.GetCountByStatus("berjalan", communityId, search)
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, err
	}

	countSelesai, err := communityRepo.GetCountByStatus("selesai", communityId, search)
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountEventInfo{}, err
	}

	// dataResponse := entity.ListEventModelToEventCore(eventData)
	// pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	// return dataResponse, pageInfo, int(totalCount), nil
	paginationInfo := pagination.CalculateData(int(totalCountWithFilter), limit, page)

	countData := pagination.MapCountEventData(totalCountWithFilter, countBelumBerjalan, countBerjalan, countSelesai)

	return entity.ListEventModelToEventCore(eventData), paginationInfo, countData, nil
}

// ReadEvent implements entity.CommunityRepositoryInterface.
func (communityRepo *communityRepository) ReadEvent(communityId string, eventId string) (entity.CommunityEventCore, error) {
	eventData := model.CommunityEvent{}

	tx := communityRepo.db.Where("community_id = ? AND id = ?", communityId, eventId).First(&eventData)
	if tx.Error != nil {
		return entity.CommunityEventCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return entity.CommunityEventCore{}, errors.New(constanta.ERROR_RECORD_NOT_FOUND)
	}

	dataResponse := entity.EventModelToEventCore(eventData)
	return dataResponse, nil
}

// UpdateEvent implements entity.CommunityRepositoryInterface.
func (communityRepo *communityRepository) UpdateEvent(communityId string, eventId string, eventInput entity.CommunityEventCore, image *multipart.FileHeader) error {
	input := entity.EventCoreToEventModel(eventInput)
	var eventData model.CommunityEvent

	check := communityRepo.db.Where("community_id = ? AND id = ?", communityId, eventId).First(&eventData)
	if check.Error != nil {
		return check.Error
	}

	if image != nil {
		imageURL, errUpload := communityRepo.sb.Upload(image)
		if errUpload != nil {
			return errUpload
		}
		eventData.Image = imageURL

	} else {
		input.Image = eventData.Image
	}

	tx := communityRepo.db.Where("community_id = ? AND id = ?", communityId, eventId).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetCountByStatus implements entity.AdminRepositoryInterface.
func (communityRepo *communityRepository) GetCountByStatus(status, communityId string, search string) (int64, error) {
	var count int64
	fmt.Println("id komunitas : ", communityId)
	query := communityRepo.db.Model(&model.CommunityEvent{}).Where("status = ? AND community_id = ?", status, communityId)

	if search != "" {
		query = query.Where("title LIKE ? OR location LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	tx := query.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}
