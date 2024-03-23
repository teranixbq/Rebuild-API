package entity

import (
	"time"

	"gorm.io/gorm"
)

type UsersCore struct {
	Id                string
	Email             string
	Password          string
	Badge             string
	ConfirmPassword   string
	Fullname          string
	Phone             string
	Address           string
	DateOfBirth       string
	Purpose           string
	Point             int
	Communities       []UserCommunityCore
	Community_id      []string
	IsVerified        bool
	VerificationToken string
	Otp               string
	NewPassword       string
	OtpExpiration     int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeleteAt          gorm.DeletedAt
}

type UserCommunityCore struct {
	Id       string
	Name     string
	Image    string
	Location string
}

type UserDailyPointsCore struct {
	UsersID      string
	DailyPointID int
	Claim        bool
	CreatedAt    time.Time
}
