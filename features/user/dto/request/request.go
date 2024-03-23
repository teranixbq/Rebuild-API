package request

type UserRegister struct {
	Fullname        string `json:"fullname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Fullname    string `json:"fullname"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Purpose     string `json:"purpose"`
}

type UserNewPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserUpdatePassword struct {
	Password        string `json:"password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
type UserSendOTP struct {
	Email string `json:"email"`
}

type UserVerifyOTP struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type UserJoinCommunity struct{
	Communitiy string `json:"community"`
}