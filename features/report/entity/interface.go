package entity

import "mime/multipart"

type ReportRepositoryInterface interface {
	Insert(reportInput ReportCore, images []*multipart.FileHeader) (ReportCore, error)
	SelectById(idReport string) (ReportCore, error)
	ReadAllReport(idUser string) ([]ReportCore, error)
}

type ReportServiceInterface interface {
	Create(reportInput ReportCore, userId string, images []*multipart.FileHeader) (ReportCore, error)
	ReadAllReport(idUser string) ([]ReportCore, error)
	SelectById(idReport string) (ReportCore, error)
}
