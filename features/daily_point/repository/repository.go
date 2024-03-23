package repository

import (
	"errors"
	"fmt"
	"recything/features/daily_point/entity"
	"recything/features/daily_point/model"
	mission "recything/features/mission/entity"
	trashEx "recything/features/trash_exchange/entity"
	user_entity "recything/features/user/entity"
	user "recything/features/user/model"
	voucher "recything/features/voucher/entity"
	"recything/utils/constanta"

	//"recything/utils/constanta"
	"time"

	"gorm.io/gorm"
)

type dailyPointRepository struct {
	db                *gorm.DB
	missionRepository mission.MissionRepositoryInterface
	trashExRepository trashEx.TrashExchangeRepositoryInterface
	userRepository    user_entity.UsersRepositoryInterface
	voucherRepository voucher.VoucherRepositoryInterface
}

func NewDailyPointRepository(db *gorm.DB, missionRepository mission.MissionRepositoryInterface, trashExRepository trashEx.TrashExchangeRepositoryInterface, userRepository user_entity.UsersRepositoryInterface, voucherRepository voucher.VoucherRepositoryInterface) entity.DailyPointRepositoryInterface {
	return &dailyPointRepository{
		db:                db,
		missionRepository: missionRepository,
		trashExRepository: trashExRepository,
		userRepository:    userRepository,
		voucherRepository: voucherRepository,
	}
}

// DailyClaim implements entity.DailyPointRepositoryInterface.
func (dailyPoint *dailyPointRepository) DailyClaim(userId string) error {
	userDaily := new(user.UserDailyPoints)
	userProf := new(user.Users)
	pointData := new(model.DailyPoint)
	tx := dailyPoint.db.Begin()

	userDaily.UsersID = userId

	//pengecekan user
	check := tx.Where("id = ?", userId).First(&userProf)
	if check.Error != nil {
		return check.Error
	}

	//melakukan pengecekan untuk menghitung hari claim
	var countData int64
	err := tx.Model(&user.UserDailyPoints{}).Where("users_id = ?", userId).Count(&countData).Error
	if err != nil {
		fmt.Println("masuk error count data : ", err)
		tx.Rollback()
		return err
	}
	dpId := countData + 1

	if dpId == 8 {
		errDeleteAll := tx.Where("users_id = ?", userId).Delete(&user.UserDailyPoints{}).Error
		if errDeleteAll != nil {
			tx.Rollback()
			return errDeleteAll
		}
		dpId = 1
	}
	//melakukakn pengecekan agar user tidak claim 2 kali sehari
	today := time.Now().Truncate(24 * time.Hour)
	fmt.Println("tanggal hari ini : ", today)
	var existingCount int64
	errClaim := tx.Model(&user.UserDailyPoints{}).Where("users_id = ? AND created_at = ?", userId, today).Count(&existingCount).Error
	if errClaim != nil {
		tx.Rollback()
		return errClaim
	}
	fmt.Println("check : ", existingCount)

	//apabila lebih dari 1 maka user telah melakukan claim
	if existingCount > 0 {
		tx.Rollback()
		return errors.New("user telah melakukan claim hari ini")
	}

	//melakukan pengecekan id pada dailypoint untuk mengambil point
	errDPId := tx.Where("id = ?", dpId).First(&pointData)
	if errDPId.Error != nil {
		fmt.Println("err id point : ", errDPId.Error)
		tx.Rollback()
		return errDPId.Error
	}

	userDaily.DailyPointID = int(dpId)
	userDaily.CreatedAt = time.Now().Truncate(24 * time.Hour)
	userProf.Point += pointData.Point
	fmt.Println("daily point Id : ", userDaily.DailyPointID)
	fmt.Println("user point : ", userProf.Point)

	//save daily record
	saveDaily := tx.Create(&userDaily)
	if saveDaily.Error != nil {
		tx.Rollback()
		return saveDaily.Error
	}

	//update user point
	savePoint := tx.Updates(&userProf).Error
	if savePoint != nil {
		tx.Rollback()
		return savePoint
	}

	tx.Commit()

	return nil
}

// PostWeekly implements entity.DailyPointRepositoryInterface.
func (daily *dailyPointRepository) PostWeekly() error {
	dailyPointData := []entity.DailyPointCore{}

	hari1 := entity.DailyPointCore{Point: 100, Description: "hari 1"}
	dailyPointData = append(dailyPointData, hari1)

	hari2 := entity.DailyPointCore{Point: 150, Description: "hari 2"}
	dailyPointData = append(dailyPointData, hari2)

	hari3 := entity.DailyPointCore{Point: 200, Description: "hari 3"}
	dailyPointData = append(dailyPointData, hari3)

	hari4 := entity.DailyPointCore{Point: 250, Description: "hari 4"}
	dailyPointData = append(dailyPointData, hari4)

	hari5 := entity.DailyPointCore{Point: 300, Description: "hari 5"}
	dailyPointData = append(dailyPointData, hari5)

	hari6 := entity.DailyPointCore{Point: 350, Description: "hari 6"}
	dailyPointData = append(dailyPointData, hari6)

	hari7 := entity.DailyPointCore{Point: 400, Description: "hari 7"}
	dailyPointData = append(dailyPointData, hari7)

	post := entity.ListDailyPointCoreToDailyPointModel(dailyPointData)

	tx := daily.db.Create(post)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// History Point

func (daily *dailyPointRepository) GetAllHistoryPoint(userID string) ([]map[string]interface{}, error) {
	dataMission, errFind := daily.missionRepository.FindAllMissionUser(userID, constanta.SELESAI)
	if errFind != nil {
		return nil, errFind
	}

	dataUser, errUser := daily.userRepository.FindById(userID)
	if errUser != nil {
		return nil, errUser
	}

	dataTrashEx, errTrash := daily.trashExRepository.GetByEmail(dataUser.Email)
	if errTrash != nil {
		return nil, errTrash
	}

	dataVoucherEx, errVoucher := daily.voucherRepository.GetAllExchangeHistory(userID)
	if errVoucher != nil {
		return nil, errVoucher
	}

	var combinedData []map[string]interface{}
	for _, missionData := range dataMission {

		data := mission.MissionHistoriesCoreToMap(missionData)
		combinedData = append(combinedData, data)
	}

	combinedData = append(combinedData, dataTrashEx...)
	combinedData = append(combinedData, dataVoucherEx...)
	return combinedData, nil
}

func (daily *dailyPointRepository) GetByIdHistoryPoint(userID, idTransaction string) (map[string]interface{}, error) {
	dataMission, errMission := daily.missionRepository.FindHistoryByIdTransaction(userID, idTransaction)
	if errMission == nil {
		return dataMission, nil
	}

	dataUser, errUser := daily.userRepository.FindById(userID)
	if errUser != nil {
		return nil, errUser
	}

	dataTrashEx, errTrash := daily.trashExRepository.GetTrashExchangeByIdTransaction(dataUser.Email, idTransaction)
	if errTrash == nil {
		return dataTrashEx, nil
	}

	dataVoucherEx, errVoucher := daily.voucherRepository.GetByIdExchangeTransactions(userID, idTransaction)
	if errVoucher == nil {
		return dataVoucherEx, nil
	}

	return nil, errors.New("record not found")
}

func (daily *dailyPointRepository) GetAllClaimedDaily(userID string) ([]user_entity.UserDailyPointsCore, error) {
	dataDaily := []user.UserDailyPoints{}

	tx := daily.db.Where("users_id = ? ", userID).Order("created_at DESC").Find(&dataDaily)
	if tx.Error != nil {
		return []user_entity.UserDailyPointsCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return []user_entity.UserDailyPointsCore{}, tx.Error

	}

	dataResponse := user_entity.ListUserDailyPointsModelToUserDailyPointsCore(dataDaily)

	return dataResponse,nil
}
