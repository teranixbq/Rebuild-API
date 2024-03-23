package entity

import (
	"recything/features/report/model"
)

func ImageModelToImageCore(image model.Image) ImageCore {
	return ImageCore{
		ID:        image.ID,
		ReportID:  image.ReportId,
		Image:     image.Image,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}

func ListImageModelToImageCore(images []model.Image) []ImageCore {
	coreImages := []ImageCore{}
	for _, v := range images {
		image := ImageModelToImageCore(v)
		coreImages = append(coreImages, image)
	}
	return coreImages
}

func ReportModelToReportCore(report model.Report) ReportCore {
	reportCore := ReportCore{
		ID:                   report.Id,
		ReportType:           report.ReportType,
		UserId:               report.UsersId,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		Description:          report.Description,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		CompanyName:          report.CompanyName,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CreatedAt:            report.CreatedAt,
		UpdatedAt:            report.UpdatedAt,
	}
	image := ListImageModelToImageCore(report.Images)
	reportCore.Images = image
	return reportCore

}

func ImageCoreToImageModel(image ImageCore) model.Image {
	return model.Image{
		ID:        image.ID,
		ReportId:  image.ReportID,
		Image:     image.Image,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}

func ListImageCoreToImageModel(images []ImageCore) []model.Image {
	coreImages := []model.Image{}
	for _, v := range images {
		image := ImageCoreToImageModel(v)
		coreImages = append(coreImages, image)
	}
	return coreImages
}

func ReportCoreToReportModel(report ReportCore) model.Report {
	reportModel := model.Report{
		Id:                   report.ID,
		ReportType:           report.ReportType,
		UsersId:              report.UserId,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		Description:          report.Description,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		CompanyName:          report.CompanyName,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CreatedAt:            report.CreatedAt,
		UpdatedAt:            report.UpdatedAt,
	}
	image := ListImageCoreToImageModel(report.Images)
	reportModel.Images = image
	return reportModel

}

func ListReportModelToReportCore(mainData []model.Report) []ReportCore {
    listReport := []ReportCore{}
    for _, report := range mainData {
        reportModel := ReportModelToReportCore(report)
        listReport = append(listReport, reportModel)
    }
    return listReport
}
