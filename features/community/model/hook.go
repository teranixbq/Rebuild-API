package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (community *Community) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	community.Id = newUuid.String()

	return nil
}

func (communityevent *CommunityEvent) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	communityevent.Id = newUuid.String()

	return nil
}