package repository

import (
	"errors"
	"mime/multipart"

	"recything/features/admin/entity"
	"recything/features/admin/model"

	report "recything/features/report/entity"
	reportModel "recything/features/report/model"

	user "recything/features/user/entity"
	userModel "recything/features/user/model"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"recything/utils/storage"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) entity.AdminRepositoryInterface {
	return &AdminRepository{
		db: db,
	}
}

func (ar *AdminRepository) Create(image *multipart.FileHeader, data entity.AdminCore) (entity.AdminCore, error) {
	dataAdmins := entity.AdminCoreToAdminModel(data)

	if image != nil {
		imageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			return entity.AdminCore{}, errUpload
		}
		dataAdmins.Image = imageURL
	} else {
		dataAdmins.Image = constanta.IMAGE_ADMIN + data.Fullname
	}

	tx := ar.db.Create(&dataAdmins)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

func (ar *AdminRepository) SelectAll(page, limit int, search string) ([]entity.AdminCore, pagination.PageInfo, int, error) {
	dataAdmins := []model.Admin{}
	offsetInt := (page - 1) * limit

	totalCount, err := ar.GetCount(search, constanta.ADMIN)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	paginationQuery := ar.db.Limit(limit).Offset(offsetInt)
	if search == "" {
		tx := paginationQuery.Where("role = ? ", constanta.ADMIN).Find(&dataAdmins)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if search != "" {
		tx := paginationQuery.Where("role = ? AND fullname LIKE ?", constanta.ADMIN, "%"+search+"%").Find(&dataAdmins)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	dataResponse := entity.ListAdminModelToAdminCore(dataAdmins)
	paginationInfo := pagination.CalculateData(totalCount, limit, page)

	return dataResponse, paginationInfo, totalCount, nil
}

func (ar *AdminRepository) GetCount(search, role string) (int, error) {
	var totalCount int64
	model := ar.db.Model(&model.Admin{})
	if search == "" {
		tx := model.Where("role = ? ", constanta.ADMIN).Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}

	}

	if search != "" {
		tx := model.Where("role = ? AND fullname LIKE ?", constanta.ADMIN, "%"+search+"%").Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	return int(totalCount), nil
}

func (ar *AdminRepository) SelectById(adminId string) (entity.AdminCore, error) {
	dataAdmins := model.Admin{}

	tx := ar.db.Where("id = ? ", adminId).First(&dataAdmins)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.AdminCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

func (ar *AdminRepository) Update(image *multipart.FileHeader, adminId string, data entity.AdminCore) error {
	dataAdmins := entity.AdminCoreToAdminModel(data)

	dataFind , errFind := ar.SelectById(adminId) 
	if errFind != nil  {
		return errFind
	}

	if image != nil {
		imageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			return errUpload
		}
		dataAdmins.Image = imageURL
	} else {
		dataAdmins.Image = dataFind.Image
	}

	tx := ar.db.Where("id = ?", adminId).Updates(&dataAdmins)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (ar *AdminRepository) Delete(adminId string) error {
	dataAdmins := model.Admin{}

	result, _ := ar.SelectById(adminId)
	if result.Role == constanta.SUPERADMIN {
		return errors.New("tidak bisa menghapus super admin")
	}

	tx := ar.db.Where("id = ? AND role = ?", adminId, constanta.ADMIN).Delete(&dataAdmins)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (ar *AdminRepository) FindByEmail(email string) error {
	dataAdmins := model.Admin{}

	tx := ar.db.Where("email = ?", email).First(&dataAdmins)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (ar *AdminRepository) FindByEmailANDPassword(data entity.AdminCore) (entity.AdminCore, error) {
	dataAdmins := model.Admin{}

	tx := ar.db.Where("email = ?", data.Email).First(&dataAdmins)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.AdminCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	if comparePass := helper.CompareHash(dataAdmins.Password, data.Password); !comparePass {
		return entity.AdminCore{}, errors.New(constanta.ERROR_PASSWORD)
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

// Manage Users
func (ar *AdminRepository) GetAllUsers(search string, page, limit int) ([]user.UsersCore, pagination.PageInfo, int, error) {
	dataUsers := []userModel.Users{}

	offset := (page - 1) * limit
	query := ar.db.Model(&userModel.Users{})

	if search != "" {
		query = query.Where("fullname LIKE ? or point LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	query = query.Offset(offset).Limit(limit)

	if err := query.Find(&dataUsers).Error; err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	dataAllUser := user.ListUserModelToUserCore(dataUsers)
	paginationInfo := pagination.CalculateData(int(totalCount), limit, page)

	return dataAllUser, paginationInfo, int(totalCount), nil
}

func (ar *AdminRepository) GetByIdUser(userId string) (user.UsersCore, error) {
	dataUsers := userModel.Users{}

	tx := ar.db.Where("id = ?", userId).Find(&dataUsers)
	if tx.Error != nil {
		return user.UsersCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return user.UsersCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := user.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

func (ar *AdminRepository) DeleteUsers(userId string) error {
	dataUsers := userModel.Users{}

	tx := ar.db.Where("id = ?", userId).Delete(&dataUsers)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetByStatusReport implements entity.AdminRepositoryInterface.
func (ar *AdminRepository) GetAllReport(status, search string, page, limit int) ([]report.ReportCore, pagination.PageInfo, pagination.CountDataInfo, error) {
	dataReports := []reportModel.Report{}

	offset := (page - 1) * limit
	query := ar.db.Model(&reportModel.Report{})

	var totalCountWithoutFilter int64
	if err := query.Count(&totalCountWithoutFilter).Error; err != nil {
		return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, err
	}

	totalCountWithFilter := totalCountWithoutFilter

	if search != "" {
		if err := query.Joins("JOIN users AS u_search ON reports.users_id = u_search.id").
			Where("u_search.fullname LIKE ? or reports.id LIKE ?", "%"+search+"%", "%"+search+"%").
			Count(&totalCountWithFilter).Error; err != nil {
			return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, err
		}
	}

	if status != "" {
		query = query.Where("reports.status = ?", status)
	}

	query = query.Offset(offset).Limit(limit)

	err := query.Find(&dataReports).Error
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, err
	}

	countPerluDitinjau, err := ar.GetCountByStatus("perlu ditinjau", search)
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, err
	}

	countDiterima, err := ar.GetCountByStatus("diterima", search)
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, err
	}

	countDitolak, err := ar.GetCountByStatus("ditolak", search)
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, err
	}

	paginationInfo := pagination.CalculateData(int(totalCountWithFilter), limit, page)

	countData := pagination.MapCountData(totalCountWithFilter, countPerluDitinjau, countDiterima, countDitolak)

	return report.ListReportModelToReportCore(dataReports), paginationInfo, countData, nil
}

// GetCountByStatus implements entity.AdminRepositoryInterface.
func (ar *AdminRepository) GetCountByStatus(status, search string) (int64, error) {
	var count int64
	query := ar.db.Model(&reportModel.Report{}).Where("reports.status = ?", status)

	if search != "" {
		query = query.Joins("JOIN users AS u_status_filtered ON reports.users_id = u_status_filtered.id").
			Where("u_status_filtered.fullname LIKE ? or reports.id LIKE ? ", "%"+search+"%", "%"+search+"%")
	}

	tx := query.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return count, nil
}

// UpdateStatusReport implements entity.AdminRepositoryInterface.
func (ar *AdminRepository) UpdateStatusReport(id, status, reason string) (report.ReportCore, error) {
	dataReports := reportModel.Report{}

	errData := ar.db.Where("id = ?", id).First(&dataReports)
	if errData.Error != nil {
		return report.ReportCore{}, errData.Error
	}

	dataReports.Status = status
	dataReports.RejectionDescription = reason
	tx := ar.db.Save(&dataReports)
	if tx.Error != nil {
		return report.ReportCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return report.ReportCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := report.ReportModelToReportCore(dataReports)
	return dataResponse, nil
}

func (ar *AdminRepository) GetReportById(id string) (report.ReportCore, error) {
	dataReports := reportModel.Report{}

	tx := ar.db.Preload("Images").Where("id = ?", id).First(&dataReports)
	if tx.Error != nil {
		return report.ReportCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return report.ReportCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := report.ReportModelToReportCore(dataReports)
	return dataResponse, nil
}
