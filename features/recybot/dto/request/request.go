package request

type RecybotManageRequest struct {
	Category string `json:"category"`
	Question string `json:"question"`
}

type RecybotRequest struct {
	Question string `json:"question"`
}