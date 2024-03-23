package repository

import (
	"errors"
	"mime/multipart"
	"recything/features/report/entity"
	"recything/features/report/model"
	"recything/utils/constanta"
	"recything/utils/storage"

	"gorm.io/gorm"
)

type reportRepository struct {
	db *gorm.DB
	sb storage.StorageInterface
}

func NewReportRepository(db *gorm.DB,sb storage.StorageInterface) entity.ReportRepositoryInterface {
	return &reportRepository{
		db: db,
		sb: sb,
	}
}

// ReadAllReport implements entity.ReportRepositoryInterface.
func (report *reportRepository) ReadAllReport(idUser string) ([]entity.ReportCore, error) {
	dataReport := []model.Report{}

	tx := report.db.Preload("Images").Where("users_id = ?", idUser).Order("created_at DESC").Find(&dataReport)
	if tx.Error != nil {
		return nil, tx.Error
	}

	mapData := make([]entity.ReportCore, len(dataReport))
	for i, value := range dataReport {
		mapData[i] = entity.ReportModelToReportCore(value)
	}

	return mapData, nil
}

func (report *reportRepository) Insert(reportInput entity.ReportCore, images []*multipart.FileHeader) (entity.ReportCore, error) {
	dataReport := entity.ReportCoreToReportModel(reportInput)

	tx := report.db.Create(&dataReport)
	if tx.Error != nil {
		return entity.ReportCore{}, tx.Error
	}

	for _, image := range images {
		imageURL, uploadErr := report.sb.Upload(image)
		if uploadErr != nil {
			return entity.ReportCore{}, uploadErr
		}

		ImageList := entity.ImageCore{}
		ImageList.ReportID = dataReport.Id
		ImageList.Image = imageURL
		ImageSave := entity.ImageCoreToImageModel(ImageList)

		if err := report.db.Create(&ImageSave).Error; err != nil {
			return entity.ReportCore{}, err
		}

		// Tambahkan informasi file ke laporan
		reportInput.Images = append(reportInput.Images, ImageList)
	}

	ReportCreated := entity.ReportModelToReportCore(dataReport)

	return ReportCreated, nil
}

func (report *reportRepository) SelectById(iDReport string) (entity.ReportCore, error) {
	dataReports := model.Report{}

	tx := report.db.Preload("Images").Where("id = ?", iDReport).First(&dataReports)
	if tx.Error != nil {
		if tx.RowsAffected == 0 {
			return entity.ReportCore{}, errors.New(constanta.ERROR_NOT_FOUND)
		}
		return entity.ReportCore{}, tx.Error
	}

	dataResponse := entity.ReportModelToReportCore(dataReports)
	return dataResponse, nil
}
