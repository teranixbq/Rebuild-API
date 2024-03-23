package request

import "recything/features/report/entity"

func ReportRequestToReportCore(report ReportRubbishRequest) entity.ReportCore {
	reportCore := entity.ReportCore{
		ReportType:           report.ReportType,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CompanyName:          report.CompanyName,
		Description:          report.Description,
	}
	image := ListImageRequestToImageCore(report.Images)
	reportCore.Images = image
	return reportCore
}

func ImagerequestToImageCore(image ImageRequest) entity.ImageCore {
	return entity.ImageCore{
		Image: image.Image,
	}
}

func ListImageRequestToImageCore(images []ImageRequest) []entity.ImageCore {
	listImage := []entity.ImageCore{}
	for _, v := range images {
		image := ImagerequestToImageCore(v)
		listImage = append(listImage, image)
	}

	return listImage
}