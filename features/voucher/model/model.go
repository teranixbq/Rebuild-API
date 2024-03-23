package model

import (
	user "recything/features/user/model"
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	Id          string `gorm:"primary key"`
	Image       string
	RewardName  string
	Point       int
	Description string
	StartDate   string
	EndDate     string
	CreatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ExchangeVoucher struct {
	Id              string     `gorm:"primary key"`
	IdUser          string     `gorm:"index"`
	Users           user.Users `gorm:"foreignKey:IdUser"`
	IdVoucher       string     `gorm:"index"`
	Vouchers        Voucher    `gorm:"foreignKey:IdVoucher"`
	Phone           string
	Status          string `gorm:"type:enum('terbaru', 'diproses', 'selesai');default:terbaru"`
	TimeTransaction string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
