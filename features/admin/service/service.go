package service

import (
	"errors"
	"mime/multipart"
	"recything/features/admin/entity"
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"recything/utils/pagination"
	"recything/utils/validation"
)

type AdminService struct {
	AdminRepository entity.AdminRepositoryInterface
}

func NewAdminService(ar entity.AdminRepositoryInterface) entity.AdminServiceInterface {
	return &AdminService{
		AdminRepository: ar,
	}
}

func (as *AdminService) Create(image *multipart.FileHeader, data entity.AdminCore) (entity.AdminCore, error) {

	errEmpty := validation.CheckDataEmpty(data.Fullname, data.Email, data.Password, data.ConfirmPassword, data.Status)
	if errEmpty != nil {
		return entity.AdminCore{}, errors.New(constanta.ERROR_EMPTY)
	}

	status, errEqual := validation.CheckEqualData(data.Status, constanta.STATUS_ADMIN)
	if errEqual != nil {
		return entity.AdminCore{}, errors.New("error : status input tidak valid")
	}

	errEmail := validation.EmailFormat(data.Email)
	if errEmail != nil {
		return entity.AdminCore{}, errors.New(constanta.ERROR_FORMAT_EMAIL)
	}

	errLength := validation.MinLength(data.Password, 8)
	if errLength != nil {
		return entity.AdminCore{}, errors.New(constanta.ERROR_LENGTH_PASSWORD)
	}

	errFind := as.AdminRepository.FindByEmail(data.Email)
	if errFind == nil {
		return entity.AdminCore{}, errors.New(constanta.ERROR_EMAIL_EXIST)
	}

	if data.ConfirmPassword != data.Password {
		return entity.AdminCore{}, errors.New(constanta.ERROR_CONFIRM_PASSWORD)
	}

	data.Status = status

	dataAdmins, errCreate := as.AdminRepository.Create(image, data)
	if errCreate != nil {
		return entity.AdminCore{}, errors.New("gagal membuat data admin")
	}

	return dataAdmins, nil
}

func (as *AdminService) GetAll(page, limit, search string) ([]entity.AdminCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	dataAdmins, pagnationInfo, count, err := as.AdminRepository.SelectAll(pageInt, limitInt, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, errors.New("gagal mengambil semua data admin")
	}
	return dataAdmins, pagnationInfo, count, nil
}

func (as *AdminService) GetById(adminId string) (entity.AdminCore, error) {

	dataAdmins, err := as.AdminRepository.SelectById(adminId)
	if err != nil {
		return entity.AdminCore{}, err
	}

	return dataAdmins, nil
}

func (as *AdminService) UpdateById(image *multipart.FileHeader, adminId string, data entity.AdminCore) error {

	errEmpty := validation.CheckDataEmpty(data.Fullname, data.Email, data.Password, data.ConfirmPassword, data.Status)
	if errEmpty != nil {
		return errors.New(constanta.ERROR_EMPTY)
	}

	status, errEqual := validation.CheckEqualData(data.Status, constanta.STATUS_ADMIN)
	if errEqual != nil {
		return errors.New("error : status input tidak valid")
	}

	errEmail := validation.EmailFormat(data.Email)
	if errEmail != nil {
		return errors.New(constanta.ERROR_FORMAT_EMAIL)
	}

	errLength := validation.MinLength(data.Password, 8)
	if errLength != nil {
		return errors.New(constanta.ERROR_LENGTH_PASSWORD)
	}

	dataAdmins, errAdmins := as.AdminRepository.SelectById(adminId)
	if errAdmins != nil {
		return errAdmins
	}

	if data.Email == dataAdmins.Email {

	} else {
		errFind := as.AdminRepository.FindByEmail(data.Email)
		if errFind == nil {
			return errors.New(constanta.ERROR_EMAIL_EXIST)
		}
	}

	if data.ConfirmPassword != data.Password {
		return errors.New(constanta.ERROR_CONFIRM_PASSWORD)
	}

	HashPassword, errHash := helper.HashPassword(data.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	data.Password = HashPassword
	data.Status = status

	err := as.AdminRepository.Update(image, adminId, data)
	if err != nil {
		return err
	}

	return nil
}

func (as *AdminService) DeleteById(adminId string) error {

	err := as.AdminRepository.Delete(adminId)
	if err != nil {
		return err
	}

	return nil
}

func (as *AdminService) FindByEmailANDPassword(data entity.AdminCore) (entity.AdminCore, string, error) {

	errEmpty := validation.CheckDataEmpty(data.Email, data.Password)
	if errEmpty != nil {
		return entity.AdminCore{}, "", errEmpty
	}

	errEmail := validation.EmailFormat(data.Email)
	if errEmail != nil {
		return entity.AdminCore{}, "", errEmail
	}

	data, err := as.AdminRepository.FindByEmailANDPassword(data)
	if err != nil {
		return entity.AdminCore{}, "", errors.New("error : email atau password salah")
	}

	token, errToken := jwt.CreateToken(data.Id, data.Role)
	if errToken != nil {
		return entity.AdminCore{}, "", errors.New("gagal membuat token session")
	}
	return data, token, nil
}

//Manage Users

func (as *AdminService) GetAllUsers(search, page, limit string) ([]user.UsersCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}
	data, pagnationInfo, count, err := as.AdminRepository.GetAllUsers(search, pageInt, limitInt)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return data, pagnationInfo, count, nil
}

func (as *AdminService) GetByIdUsers(userId string) (user.UsersCore, error) {

	data, err := as.AdminRepository.GetByIdUser(userId)

	if err != nil {
		return user.UsersCore{}, err
	}

	return data, nil
}

func (as *AdminService) DeleteUsers(userId string) error {

	err := as.AdminRepository.DeleteUsers(userId)
	if err != nil {
		return err
	}

	return nil
}

// Manage Reporting
func (as *AdminService) GetAllReport(status, search, page, limit string) ([]report.ReportCore, pagination.PageInfo, pagination.CountDataInfo, error) {
	pageInt, limitInt, validationErr := validation.ValidateTypePaginationParameter(limit, page)
	if validationErr != nil {
		return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, validationErr
	}

	pageValid, limitValid := validation.ValidateCountLimitAndPage(pageInt, limitInt)

	validStatus := map[string]bool{
		"perlu ditinjau": true,
		"diterima":       true,
		"ditolak":        true,
	}

	if _, ok := validStatus[status]; status != "" && !ok {
		return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, errors.New("status tidak valid")
	}

	data, paginationInfo, count, err := as.AdminRepository.GetAllReport(status, search, pageValid, limitValid)
	if err != nil {
		return nil, pagination.PageInfo{}, pagination.CountDataInfo{}, err
	}

	return data, paginationInfo, count, nil
}

// UpdateStatusReport implements entity.AdminServiceInterface.
func (as *AdminService) UpdateStatusReport(id string, status string, reason string) (report.ReportCore, error) {

	errEmpty := validation.CheckDataEmpty(status)
	if errEmpty != nil {
		return report.ReportCore{}, errEmpty
	}

	if status == "diterima" && reason != "" {
		return report.ReportCore{}, errors.New("tidak perlu memberikan alasan laporan")
	}

	if status == "ditolak" && reason == "" {
		return report.ReportCore{}, errors.New("alasan harus dilengkapi saat menolak laporan")
	}

	dataStatus, err := as.AdminRepository.GetReportById(id)
	if err != nil {
		return report.ReportCore{}, err
	}

	if dataStatus.Status == "diterima" || dataStatus.Status == "ditolak" {
		return report.ReportCore{}, errors.New("status sudah diterima atau ditolak")
	}

	data, err := as.AdminRepository.UpdateStatusReport(id, status, reason)
	if err != nil {
		return report.ReportCore{}, errors.New("gagal update status")
	}

	return data, nil
}

// GetReportById implements entity.AdminServiceInterface.
func (as *AdminService) GetReportById(id string) (report.ReportCore, error) {
	idReport, err := as.AdminRepository.GetReportById(id)
	if err != nil {
		return report.ReportCore{}, err
	}
	return idReport, err
}
