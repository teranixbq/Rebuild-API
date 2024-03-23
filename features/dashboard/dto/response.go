package dto

type GetCountUserResponse struct {
	TotalUserActive string `json:"total_user_active"`
	Percentage      string `json:"percentage"`
	Status          string `json:"status"`
}

type GetCountExchangeVoucherResponse struct {
	TotalExchange string `json:"total_exchange"`
	Percentage    string `json:"percentage"`
	Status        string `json:"status"`
}

type GetCountReportingResponse struct {
	TotalReporting string `json:"total_report"`
	Percentage     string `json:"percentage"`
	Status         string `json:"status"`
}

type GetCountTrashExchangeResponse struct {
	TotalTrashExchange string `json:"total_recycle"`
	Percentage         string `json:"percentage"`
	Status             string `json:"status"`
}

type GetCountTrashExchangeIncomeResponse struct {
	TotalIncome int    `json:"total_income"`
	Percentage  string `json:"percentage"`
	Status      string `json:"status"`
}

type GetCountScaleTypeResponse struct {
	Company string `json:"company"`
	Person  string `json:"person"`
}

type UserRankingResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Point int    `json:"point"`
}

type WeeklyStatsResponse struct {
	Week      int `json:"week"`
	TrashType int `json:"trash_type"`
	ScaleType int `json:"scale_type"`
}

type MonthlyStatsResponse struct {
	Month     int `json:"month"`
	TrashType int `json:"trash_type"`
	ScaleType int `json:"scale_type"`
}
