package request

type ReportRubbishRequest struct {
	ReportType           string  `json:"report_type" form:"report_type"`
	Longitude            float64 `json:"longitude" form:"longitude"`
	Latitude             float64 `json:"latitude" form:"latitude"`
	Location             string  `json:"location" form:"location"`
	AddressPoint         string  `json:"address_point" form:"address_point"`
	Status               string  `json:"status" form:"status"`
	TrashType            string  `json:"trash_type" form:"trash_type"`
	ScaleType            string  `json:"scale_type" form:"scale_type"`
	InsidentDate         string  `json:"insident_date" form:"insident_date"`
	InsidentTime         string  `json:"insident_time" form:"insident_time"`
	CompanyName          string  `json:"company_name" form:"company_name"`
	DangerousWaste       bool    `json:"dangerous_waste" form:"dangerous_waste"`
	RejectionDescription string  `json:"rejection_description" form:"rejection_description"`
	Description          string  `json:"description" form:"description"`
	Images               []ImageRequest `json:"images" form:"images"`
}

type ImageRequest struct {
	Image string `json:"image" form:"image"`
}

type UpdateStatusReportRubbish struct {
	Status               string `json:"status"`
	RejectionDescription string `json:"rejection_description"`
}
