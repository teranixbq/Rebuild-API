package repository

import (
	"errors"
	achievement "recything/features/achievement/entity"
	moco "recything/features/community/model"
	"recything/features/user/entity"
	"recything/features/user/model"
	"recything/utils/constanta"

	"gorm.io/gorm"
)

type userRepository struct {
	db              *gorm.DB
	achievementRepo achievement.AchievementRepositoryInterface
}

func NewUserRepository(db *gorm.DB, achievementRepo achievement.AchievementRepositoryInterface) entity.UsersRepositoryInterface {
	return &userRepository{
		db:              db,
		achievementRepo: achievementRepo,
	}
}

// Register implements entity.UsersRepositoryInterface.
func (ur *userRepository) Register(data entity.UsersCore) (entity.UsersCore, error) {
	request := entity.UsersCoreToUsersModel(data)

	tx := ur.db.Create(&request)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := entity.UsersModelToUsersCore(request)
	return dataResponse, nil
}

// GetById implements entity.UsersRepositoryInterface.
func (ur *userRepository) GetById(id string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	tx := ur.db.Preload("Communities").Where("id = ?", id).First(&dataUsers)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

func (ur *userRepository) FindByEmail(email string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("email = ?", email).First(&dataUsers)

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New(constanta.ERROR_DATA_EMAIL)
	}

	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

// UpdateById implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdateById(id string, data entity.UsersCore) error {

	request := entity.UsersCoreToUsersModel(data)

	tx := ur.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}

func (ur *userRepository) UpdateBadge(id string) error {
	dataUsers := model.Users{}

	tx := ur.db.Where("id = ?", id).First(&dataUsers)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	dataAchievement, errAchievement := ur.achievementRepo.GetAllAchievement()
	if errAchievement != nil {
		return errAchievement
	}

	for i := len(dataAchievement) - 1; i >= 0; i-- {
		v := dataAchievement[i]
		if dataUsers.Point >= v.TargetPoint {
			dataUsers.Badge = v.Name
		}
	}

	tx = ur.db.Save(&dataUsers)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// ForgetPassword implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdatePassword(id string, data entity.UsersCore) error {

	request := entity.UsersCoreToUsersModel(data)

	tx := ur.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}

// GetByVerificationToken implements entity.UsersRepositoryInterface.
func (ur *userRepository) GetByVerificationToken(token string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("verification_token = ?", token).First(&dataUsers)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New("token tidak ditemukan")
	}

	userToken := entity.UsersModelToUsersCore(dataUsers)
	return userToken, nil
}

// UpdateIsVerified implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdateIsVerified(id string, isVerified bool) error {
	dataUser := model.Users{}

	tx := ur.db.Where("id = ?", id).First(&dataUser)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	dataUser.IsVerified = isVerified

	errSave := ur.db.Save(&dataUser)
	if errSave.Error != nil {
		return errSave.Error
	}

	return nil
}

// SendOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) SendOTP(emailUser string, otp string, expiry int64) (data entity.UsersCore, err error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("email = ?", emailUser).First(&dataUsers)
	if tx.Error != nil {
		if tx.RowsAffected == 0 {
			return entity.UsersCore{}, errors.New(constanta.ERROR_DATA_EMAIL)
		}
		return entity.UsersCore{}, tx.Error
	}

	dataUsers.Otp = otp
	dataUsers.OtpExpiration = expiry

	errUpdate := ur.db.Save(&dataUsers).Error
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)

	return dataResponse, nil
}

// VerifyOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) VerifyOTP(email, otp string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("otp = ? AND email = ?", otp, email).First(&dataUsers)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New("email atau otp tidak ditemukan")
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

// ResetOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) ResetOTP(otp string) (data entity.UsersCore, err error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("otp = ?", otp).First(&dataUsers)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New("otp tidak ditemukan")
	}

	dataUsers.Otp = ""
	dataUsers.OtpExpiration = 0

	errUpdate := ur.db.Save(&dataUsers).Error
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

// ForgetPassword implements entity.UsersRepositoryInterface.
func (ur *userRepository) NewPassword(email string, data entity.UsersCore) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	tx := ur.db.Where("email = ?", email).First(&dataUsers)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New(constanta.ERROR_DATA_EMAIL)
	}

	errUpdate := ur.db.Model(&dataUsers).Updates(entity.UsersCoreToUsersModel(data))
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate.Error
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)

	return dataResponse, nil
}

func (ur *userRepository) UpdateUserPoint(id string, point int) error {

	tx := ur.db.Model(&model.Users{}).Where("id = ?", id).Update("point", point)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// JoinCommunity implements entity.UsersRepositoryInterface.
func (ur *userRepository) JoinCommunity(communityId string, userId string) error {
	dataCommunity := moco.Community{}
	saveData := model.UserCommunity{}

	// ambil data community
	txCommunity := ur.db.Where("id = ?", communityId).First(&dataCommunity)
	if txCommunity.Error != nil {
		return txCommunity.Error
	}

	saveData.CommunityID = dataCommunity.Id
	saveData.UsersID = userId

	txSave := ur.db.Create(&saveData).Error
	if txSave != nil {
		return txSave
	}

	return nil
}

// For History Point

func (ur *userRepository) FindById(userID string) (entity.UsersCore, error) {
	dataUser := model.Users{}

	tx := ur.db.Where("id = ?", userID).First(&dataUser)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.UsersCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := entity.UsersModelToUsersCore(dataUser)
	return dataResponse, nil
}
