package request

type AdminRequest struct {
	Fullname        string `json:"fullname" form:"fullname"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Status          string `json:"status" form:"status"`
}

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
