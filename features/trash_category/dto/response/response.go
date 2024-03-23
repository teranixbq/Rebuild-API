package response

import "time"

type TrashCategory struct {
	ID        string    `json:"id"`
	TrashType string    `json:"trash_type"`
	Point     int       `json:"point"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TrashCategoriesList struct {
	ID        string `json:"id"`
	TrashType string `json:"trash_type"`
	Point     int    `json:"point"`
	Unit      string `json:"unit"`
}
