package entity

type UsersRepositoryInterface interface {
	Register(data UsersCore) (UsersCore, error)
	GetById(id string) (UsersCore, error)
	FindByEmail(email string) (UsersCore,error)
	UpdateById(id string, data UsersCore) error 
	UpdatePassword(id string, data UsersCore) error
	GetByVerificationToken(token string) (UsersCore, error)
	UpdateIsVerified(id string, isVerified bool) error
	SendOTP(emailUser string, otp string, expiry int64) (UsersCore, error)
	VerifyOTP(email, otp string) (UsersCore, error)
	ResetOTP(otp string) (UsersCore, error)
	NewPassword(email string, data UsersCore) (UsersCore, error)
	UpdateUserPoint(id string, point int)error
	JoinCommunity(communityId string, userId string) error
	FindById(userID string) (UsersCore,error)
	UpdateBadge(id string) error
}

type UsersUsecaseInterface interface {
	Register(data UsersCore) (UsersCore,error)
	Login(email, password string) (UsersCore, string, error)
	GetById(id string) (UsersCore, error)
	UpdateById(id string, data UsersCore) error
	VerifyUser(token string) (bool, error)
	UpdateIsVerified(id string, isVerified bool) error
	UpdatePassword(id string, data UsersCore)  error 
	SendOTP(emailUser string) error
	VerifyOTP(email, otp string) (string, error)
	NewPassword(email string, data UsersCore) error
	JoinCommunity(communityId string, userId string) error
}
