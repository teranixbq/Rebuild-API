package service

import (
	"errors"
	"recything/features/user/entity"
	"recything/utils/constanta"
	"recything/utils/email"
	"recything/utils/helper"
	"recything/utils/jwt"
	"recything/utils/validation"
	"time"
)

type userService struct {
	userRepo entity.UsersRepositoryInterface
}

func NewUserService(userRepo entity.UsersRepositoryInterface) entity.UsersUsecaseInterface {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements entity.UsersUsecaseInterface.
func (us *userService) Register(data entity.UsersCore) (entity.UsersCore, error) {

	errEmpty := validation.CheckDataEmpty(data.Fullname, data.Email, data.Password, data.ConfirmPassword)
	if errEmpty != nil {
		return entity.UsersCore{}, errEmpty
	}

	errEmail := validation.EmailFormat(data.Email)
	if errEmail != nil {
		return entity.UsersCore{}, errEmail
	}

	errLength := validation.MinLength(data.Password, 8)
	if errLength != nil {
		return entity.UsersCore{}, errLength
	}

	_, err := us.userRepo.FindByEmail(data.Email)
	if err == nil {
		return entity.UsersCore{},errors.New(constanta.ERROR_EMAIL_EXIST)
	}

	if data.Password != data.ConfirmPassword {
		return entity.UsersCore{}, errors.New(constanta.ERROR_CONFIRM_PASSWORD)
	}

	hashedPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		return entity.UsersCore{}, errors.New(constanta.ERROR_HASH_PASSWORD)
	}

	data.Password = hashedPassword
	uniqueToken := email.GenerateUniqueToken()
	data.VerificationToken = uniqueToken

	dataUsers, err := us.userRepo.Register(data)
	if err != nil {
		return entity.UsersCore{}, err
	}

	email.SendVerificationEmail(data.Email, uniqueToken)

	return dataUsers, nil
}

// Login implements entity.UsersUsecaseInterface.
func (us *userService) Login(email, password string) (entity.UsersCore, string, error) {

	errEmpty := validation.CheckDataEmpty(email, password)
	if errEmpty != nil {
		return entity.UsersCore{}, "", errEmpty
	}

	errEmail := validation.EmailFormat(email)
	if errEmail != nil {
		return entity.UsersCore{}, "", errEmail
	}

	dataUser, errEmail := us.userRepo.FindByEmail(email)
	if errEmail != nil {
		return entity.UsersCore{}, "", errors.New(constanta.EMAIL_NOT_REGISTER)
	}

	if !dataUser.IsVerified {
		return entity.UsersCore{}, "", errors.New("akun belum terverifikasi")
	}

	comparePass := helper.CompareHash(dataUser.Password, password)
	if !comparePass {
		return entity.UsersCore{}, "", errors.New("email atau password salah")
	}

	token, err := jwt.CreateToken(dataUser.Id, "")
	if err != nil {
		return entity.UsersCore{}, "", errors.New("gagal mendapatkan generate token")
	}
	return dataUser, token, nil
}

// GetById implements entity.UsersUsecaseInterface.
func (us *userService) GetById(id string) (entity.UsersCore, error) {
	if id == "" {
		return entity.UsersCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	loc, err := time.LoadLocation(constanta.ASIABANGKOK)
	if err != nil {
		return entity.UsersCore{}, err
	}

	now := time.Now().In(loc)
	if now.Day() == 1 {
		errBadge := us.userRepo.UpdateBadge(id)
		if errBadge != nil {
			return entity.UsersCore{}, errBadge
		}
	}

	dataUser, err := us.userRepo.GetById(id)
	if err != nil {
		return entity.UsersCore{}, errors.New("data user tidak ada")
	}
	return dataUser, nil
}

// UpdateById implements entity.UsersUsecaseInterface.
func (us *userService) UpdateById(id string, data entity.UsersCore) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	_, errGet := us.userRepo.GetById(id)
	if errGet != nil {
		return errGet
	}

	errEmpty := validation.CheckDataEmpty(data.Fullname, data.Phone, data.Address, data.DateOfBirth, data.Purpose)
	if errEmpty != nil {
		return errEmpty
	}

	errPhone := validation.PhoneNumber(data.Phone)
	if errPhone != nil {
		return errPhone
	}

	if data.DateOfBirth != "" {
		if _, errParse := time.Parse("2006-01-02", data.DateOfBirth); errParse != nil {
			return errors.New("error, tanggal harus dalam format 'yyyy-mm-dd'")
		}
	}

	err := us.userRepo.UpdateById(id, data)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePassword implements entity.UsersUsecaseInterface.
func (us *userService) UpdatePassword(id string, data entity.UsersCore) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	result, err := us.GetById(id)
	if err != nil {
		return err
	}

	errEmpty := validation.CheckDataEmpty(data.Password, data.NewPassword, data.ConfirmPassword)
	if errEmpty != nil {
		return errEmpty
	}

	errLength := validation.MinLength(data.NewPassword, 8)
	if errLength != nil {
		return errLength
	}

	ComparePass := helper.CompareHash(result.Password, data.Password)
	if !ComparePass {
		return errors.New(constanta.ERROR_PASSWORD)
	}

	if data.NewPassword != data.ConfirmPassword {
		return errors.New(constanta.ERROR_CONFIRM_PASSWORD)
	}

	HashPassword, errHash := helper.HashPassword(data.NewPassword)
	if errHash != nil {
		return errors.New(constanta.ERROR_HASH_PASSWORD)
	}
	data.Password = HashPassword

	err = us.userRepo.UpdatePassword(id, data)
	if err != nil {
		return err
	}

	return nil
}

// GetByVerificationToken implements entity.UsersUsecaseInterface.
func (us *userService) VerifyUser(token string) (bool, error) {
	if token == "" {
		return false, errors.New("invalid token")
	}

	user, err := us.userRepo.GetByVerificationToken(token)
	if err != nil {
		return false, errors.New("gagal mendapatkan data")
	}

	if user.IsVerified {
		return true, nil
	}

	err = us.userRepo.UpdateIsVerified(user.Id, true)
	if err != nil {
		return false, errors.New("gagal untuk mengaktifkan user")
	}

	return false, nil
}

// UpdateIsVerified implements entity.UsersUsecaseInterface.
func (us *userService) UpdateIsVerified(id string, isVerified bool) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	return us.userRepo.UpdateIsVerified(id, isVerified)
}

// SendOTP implements entity.UsersUsecaseInterface.
func (us *userService) SendOTP(emailUser string) error {

	errEmpty := validation.CheckDataEmpty(emailUser)
	if errEmpty != nil {
		return errEmpty
	}

	errEmail := validation.EmailFormat(emailUser)
	if errEmail != nil {
		return errEmail
	}

	otp, errGenerate := email.GenerateOTP(4)
	if errGenerate != nil {
		return errors.New("generate otp gagal")
	}

	expired := time.Now().Add(5 * time.Minute).Unix()

	_, errSend := us.userRepo.SendOTP(emailUser, otp, expired)
	if errSend != nil {
		return errSend
	}

	if email.ContainsLowerCase(otp) {
		return errors.New("otp tidak boleh mengandung huruf kecil")
	}

	email.SendOTPEmail(emailUser, otp)
	return nil
}

// VerifyOTP implements entity.UsersUsecaseInterface.
func (us *userService) VerifyOTP(email, otp string) (string, error) {

	errEmpty := validation.CheckDataEmpty(email, otp)
	if errEmpty != nil {
		return "", errEmpty
	}

	dataUsers, err := us.userRepo.VerifyOTP(email, otp)
	if err != nil {
		return "", errors.New("email atau otp salah")
	}

	if dataUsers.OtpExpiration <= time.Now().Unix() {
		return "", errors.New("otp sudah kadaluwarsa")
	}

	if dataUsers.Otp != otp {
		return "", errors.New("otp tidak valid")
	}

	token, err := jwt.CreateTokenVerifikasi(email)
	if err != nil {
		return "", errors.New("token gagal dibuat")
	}

	_, errReset := us.userRepo.ResetOTP(otp)
	if errReset != nil {
		return "", errors.New("gagal mengatur ulang OTP")
	}

	return token, nil
}

// ForgetPassword implements entity.UsersUsecaseInterface.
func (us *userService) NewPassword(email string, data entity.UsersCore) error {

	errEmpty := validation.CheckDataEmpty(email, data.Password, data.ConfirmPassword)
	if errEmpty != nil {
		return errEmpty
	}

	errEmail := validation.EmailFormat(email)
	if errEmail != nil {
		return errEmail
	}

	errLength := validation.MinLength(data.Password, 8)
	if errLength != nil {
		return errLength
	}

	if data.Password != data.ConfirmPassword {
		return errors.New(constanta.ERROR_PASSWORD)
	}

	HashPassword, errHash := helper.HashPassword(data.Password)
	if errHash != nil {
		return errors.New(constanta.ERROR_HASH_PASSWORD)
	}
	data.Password = HashPassword

	_, errNew := us.userRepo.NewPassword(email, data)
	if errNew != nil {
		return errNew
	}

	return nil
}

// JoinCommunity implements entity.UsersUsecaseInterface.
func (us *userService) JoinCommunity(communityId string, userId string) error {
	if communityId == "" || userId == "" {
		return errors.New("id tidak boleh kosong")
	}

	tx := us.userRepo.JoinCommunity(communityId, userId)
	if tx != nil {
		return tx
	}

	return nil
}
