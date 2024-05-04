package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recybot struct {
	ID        string `gorm:"primary key"`
	Category  string `gorm:"type:enum('sampah organik', 'sampah anorganik', 'informasi', 'batasan')"`
	Question  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RecybotHistory struct {
	ID        string `gorm:"primary ke by"`
	Question  string `gorm:"not null"`
	UserId    string
	Answer    string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// type RoomChat struct{
// 	Id string
// 	UserId string
// 	CreatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

func (r *Recybot) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	r.ID = newUuid.String()
	return nil
}
