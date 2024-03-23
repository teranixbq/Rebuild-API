package request

type TrashExchangeRequest struct {
	Name                 string                       `json:"name"`
	EmailUser            string                       `json:"email"`
	DropPointName        string                       `json:"drop_point_name"`
	TrashExchangeDetails []TrashExchangeDetailRequest `json:"trash_exchange_details"`
}

type TrashExchangeDetailRequest struct {
	TrashType string  `json:"trash_type"`
	Amount    float64 `json:"amount"`
}
