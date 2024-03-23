package entity

import (
	"mime/multipart"
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	"recything/utils/pagination"
)

type AdminRepositoryInterface interface {
	Create(image *multipart.FileHeader, data AdminCore) (AdminCore, error)
	SelectAll(page, limit int, search string) ([]AdminCore, pagination.PageInfo, int, error)
	SelectById(adminId string) (AdminCore, error)
	Update(image *multipart.FileHeader, adminId string, data AdminCore) error
	Delete(adminId string) error
	FindByEmail(email string) error
	FindByEmailANDPassword(data AdminCore) (AdminCore, error)
	GetCount(fullName, role string) (int, error)
	//Manage Users
	GetAllUsers(search string, page, limit int) ([]user.UsersCore, pagination.PageInfo, int, error)
	GetByIdUser(userId string) (user.UsersCore, error)
	DeleteUsers(adminId string) error
	// Manage Reporting
	GetAllReport(status, search string, page, limit int) ([]report.ReportCore, pagination.PageInfo, pagination.CountDataInfo, error)
	UpdateStatusReport(id, status, reason string) (report.ReportCore, error)
	GetReportById(id string) (report.ReportCore, error)
	GetCountByStatus(status, search string) (int64, error)
}

type AdminServiceInterface interface {
	Create(image *multipart.FileHeader, data AdminCore) (AdminCore, error)
	GetAll(page, limit, search string) ([]AdminCore, pagination.PageInfo, int, error)
	GetById(adminId string) (AdminCore, error)
	UpdateById(image *multipart.FileHeader, adminId string, data AdminCore) error
	DeleteById(adminId string) error
	FindByEmailANDPassword(data AdminCore) (AdminCore, string, error)
	//Manage Users
	GetAllUsers(search, page, limit string) ([]user.UsersCore, pagination.PageInfo, int, error)
	GetByIdUsers(adminId string) (user.UsersCore, error)
	DeleteUsers(adminId string) error
	// Manage Reporting
	GetAllReport(status, search, page, limit string) ([]report.ReportCore, pagination.PageInfo, pagination.CountDataInfo, error)
	UpdateStatusReport(id, status, reason string) (report.ReportCore, error)
	GetReportById(id string) (report.ReportCore, error)
}
