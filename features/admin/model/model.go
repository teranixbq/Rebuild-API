package model

import (
	"recything/utils/helper"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	Id              string `gorm:"primary key"`
	Image			string
	Fullname        string
	Role            string `gorm:"type:enum('admin', 'super_admin');default:'admin'"`
	Email           string
	Password        string
	ConfirmPassword string
	Status          string         `gorm:"type:enum('aktif', 'tidak aktif');default:'aktif'"`
	CreatedAt       time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt       time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	a.Id = newUuid.String()

	a.Password, _ = helper.HashPassword(a.Password)
	return nil
}
