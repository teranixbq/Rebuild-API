package response

import (
	"recything/features/report/entity"
	user "recything/features/user/entity"
)

func ImageCoreToImageResponse(image entity.ImageCore) ImageResponse {
	return ImageResponse{
		ID:        image.ID,
		Image:     image.Image,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}

func ListImageCoreToImageResponse(images []entity.ImageCore) []ImageResponse {
	ResponseImages := []ImageResponse{}
	for _, v := range images {
		image := ImageCoreToImageResponse(v)
		ResponseImages = append(ResponseImages, image)
	}
	return ResponseImages
}

func ReportCoreToReportResponse(report entity.ReportCore) ReportCreateResponse {
	reportResponse := ReportCreateResponse{
		Id:                   report.ID,
		ReportType:           report.ReportType,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		Description:          report.Description,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CompanyName:          report.CompanyName,
		CreatedAt:            report.CreatedAt,
		UpdatedAt:            report.UpdatedAt,
	}
	image := ListImageCoreToImageResponse(report.Images)
	reportResponse.Images = image
	return reportResponse

}

func ReportCoreToReportResponseForDataReporting(report entity.ReportCore, user user.UsersCore) ReportDetails {
	return ReportDetails{
		Id:           report.ID,
		ReportType:   report.ReportType,
		Fullname:     user.Fullname,
		Location:     report.Location,
		InsidentDate: report.InsidentDate,
		Status:       report.Status,
		CreatedAt:    report.CreatedAt,
	}
}

func ListReportCoresToReportResponseForDataReporting(reports []entity.ReportCore, userService user.UsersUsecaseInterface) []ReportDetails {
	responReporting := []ReportDetails{}
	for _, report := range reports {
		user, _ := userService.GetById(report.UserId)
		reports := ReportCoreToReportResponseForDataReporting(report, user)
		responReporting = append(responReporting, reports)
	}
	return responReporting
}

func ListReportCoresToReportResponse(reports []entity.ReportCore) []ReportCreateResponse {
	responReporting := []ReportCreateResponse{}
	for _, report := range reports {
		reports := ReportCoreToReportResponse(report)
		responReporting = append(responReporting, reports)
	}
	return responReporting
}

func ReportCoreToReportResponseForDataReportingId(report entity.ReportCore, user user.UsersCore) ReportDetailsById {
	reportResponse := ReportDetailsById{
		Id:                   report.ID,
		ReportType:           report.ReportType,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		Description:          report.Description,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CompanyName:          report.CompanyName,
	}

	// Menambahkan informasi pengguna ke respons
	reportResponse.Fullname = user.Fullname

	image := ListImageCoreToImageResponse(report.Images)
	reportResponse.Images = image
	return reportResponse
}
