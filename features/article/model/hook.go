package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (i *Article) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	i.Id = newUuid.String()
	return nil
}

// func (i *ArticleTrashCategory) BeforeCreate(tx *gorm.DB) (err error) {
// 	newUuid := uuid.New()
// 	i.ID = newUuid.String()
// 	return nil
// }