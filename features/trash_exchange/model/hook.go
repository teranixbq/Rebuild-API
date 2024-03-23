package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// func (user *TrashExchange) BeforeCreate(tx *gorm.DB) (err error) {
// 	newUuid := uuid.New()
// 	user.Id = newUuid.String()

// 	return nil
// }

func (user *TrashExchangeDetail) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	user.Id = newUuid.String()

	return nil
}
