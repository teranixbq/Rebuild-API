package response

import "time"

type RecybotResponse struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Question  string    `json:"question"`
	CreatedAt time.Time `json:"created_at"`
}

type RecybotHistoryResponse struct {
	UserId   string `json:"userId"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
