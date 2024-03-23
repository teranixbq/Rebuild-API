package helper

type ErrorResponseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseJson struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessResponseJsonWithPagenation struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}
type SuccessResponseJsonWithPagenationAndCount struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Count      int         `json:"count_data,omitempty"`
}

type SuccessResponseJsonWithPagenationAndCountAll struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Count      interface{} `json:"count,omitempty"`
}

type SuccessResponseJsonWithPaginationAndCount struct {
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	DataMessage string      `json:"data_message,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Pagination  interface{} `json:"pagination,omitempty"`
	Count       interface{} `json:"count,omitempty"`
}

func ErrorResponse(message string) ErrorResponseJson {
	return ErrorResponseJson{
		Status:  false,
		Message: message,
	}
}

func SuccessResponse(message string) SuccessResponseJson {
	return SuccessResponseJson{
		Status:  true,
		Message: message,
	}
}

func SuccessWithDataResponse(message string, data interface{}) SuccessResponseJson {
	return SuccessResponseJson{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func SuccessWithPaginationAndCount(message string, data interface{}, pagnation interface{}, count interface{}) SuccessResponseJsonWithPaginationAndCount {
	return SuccessResponseJsonWithPaginationAndCount{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagnation,
		Count:      count,
	}
}

func SuccessWithPagnationAndCount(message string, data interface{}, pagnation interface{}, count int) SuccessResponseJsonWithPagenationAndCount {
	return SuccessResponseJsonWithPagenationAndCount{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagnation,
		Count:      count,
	}
}
func SuccessWithPagnationAndCountAll(message string, data interface{}, pagnation interface{}, count interface{}) SuccessResponseJsonWithPagenationAndCountAll {
	return SuccessResponseJsonWithPagenationAndCountAll{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagnation,
		Count:      count,
	}
}
