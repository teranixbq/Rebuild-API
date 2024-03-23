package model

import (
	"recything/features/trash_category/model"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Id          string `gorm:"primary key"`
	Title       string
	Image       string
	Content     string
	Categories  []model.TrashCategory `gorm:"many2many:ArticleTrashCategory"`
	category_id []string              `gorm:"-"`
	Like        int                   `gorm:"default:0"`
	Share       int                   `gorm:"default:0"`
	CreatedAt   time.Time             `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time             `gorm:"type:DATETIME(0)"`
	DeletedAt   gorm.DeletedAt        `gorm:"index"`
}

type ArticleTrashCategory struct{
	TrashCategoryID string
	ArticleID string
}