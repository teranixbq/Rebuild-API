package response

import (
	"time"
)

type ArticleCreateResponse struct {
	Id          string                  `json:"id,omitempty"`
	Title       string                  `json:"title,omitempty"`
	Image       string                  `json:"image,omitempty"`
	Content     string                  `json:"content,omitempty"`
	Category_id []string                `json:"category_id,omitempty"`
	Categories  []TrashCategoryResponse `json:"categories,omitempty"`
	Like        int                     `json:"like"`
	Share       int                     `json:"share"`
	CreatedAt   time.Time               `json:"created_at,omitempty"`
	UpdatedAt   time.Time               `json:"updated_at,omitempty"`
}

type TrashCategoryResponse struct {
	Category string `json:"category,omitempty"`
	// TrashCategoryID string
}
