package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (dropPoint *DropPoints) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	dropPoint.Id = newUuid.String()

	return nil
}

func (schedule *Schedules) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	schedule.Id = newUuid.String()

	return nil
}
