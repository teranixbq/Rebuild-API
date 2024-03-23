package response

import "time"

type VoucherResponse struct {
	Id          string `json:"id"`
	Image       string `json:"image"`
	RewardName  string `json:"reward_name"`
	Point       int    `json:"point"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type ExchangeVoucheResponse struct {
	Id              string    `json:"id"`
	IdUser          string    `json:"user"`
	IdVoucher       string    `json:"voucher"`
	Phone           string    `json:"phone"`
	Status          string    `json:"status"`
	TimeTransaction string `json:"time_transaction"`
	CreatedAt       time.Time `json:"created_at"`
}
