package response

import "time"

type UserCreateResponse struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserLoginResponse struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserResponseProfile struct {
	Id           string                  `json:"id"`
	Fullname     string                  `json:"fullname"`
	Email        string                  `json:"email"`
	DateOfBirth  string                  `json:"date_of_birth"`
	Point        int                     `json:"point"`
	Phone        string                  `json:"phone"`
	Address      string                  `json:"address"`
	Purpose      string                  `json:"purpose"`
	Badge        string                  `json:"badge"`
	Communities  []UserCommunityResponse `json:"communities"`
	Community_id []string                `json:"community_id"`
}

type UserResponseManageUsers struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Point    int    `json:"point"`
}

type UserResponseDetailManageUsers struct {
	Id          string    `json:"id"`
	Fullname    string    `json:"fullname"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
	Point       int       `json:"point"`
	Purpose     string    `json:"purpose"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserCommunityResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	Location string `json:"location"`
}

type UserDailyPointsResponse struct {
	Claim        bool      `json:"claim"`
	DailyPointID int       `json:"daily_point_id"`
	CreatedAt    time.Time `json:"created_at"`
}
