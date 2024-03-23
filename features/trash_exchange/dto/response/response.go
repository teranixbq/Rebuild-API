package response

import "time"

type TrashExchangeResponse struct {
	Id                   string                       `json:"id"`
	Name                 string                       `json:"name"`
	EmailUser            string                       `json:"email"`
	DropPointName        string                       `json:"drop_point_name"`
	TotalUnit            float64                      `json:"total_unit"`
	TotalPoint           int                          `json:"total_point"`
	CreatedAt            time.Time                    `json:"created_at"`
	TrashExchangeDetails []TrashExchangeDetailRespose `json:"trash_exchange_details"`
}

type TrashExchangeDetailRespose struct {
	TrashType   string  `json:"trash_type"`
	Amount      float64 `json:"amount"`
	Unit        string  `json:"unit"`
	TotalPoints int     `json:"total_points"`
}
