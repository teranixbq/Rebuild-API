package entity

import (
	"time"
)

type ArticleCore struct {
	ID          string                     
	Title       string                     
	Image       string                     
	Content     string                     
	Categories  []ArticleTrashCategoryCore 
	Category_id []string
	Like        int       
	Share       int       
	CreatedAt   time.Time 
	UpdatedAt   time.Time 
}

type ArticleTrashCategoryCore struct {
	// TrashCategoryID string 
	Category        string 
}
