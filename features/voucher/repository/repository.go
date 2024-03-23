package repository

import (
	"errors"
	"log"
	"mime/multipart"
	"recything/features/voucher/entity"
	"recything/features/voucher/model"
	"recything/utils/constanta"
	"recything/utils/helper"

	//"recything/utils/helper"
	"recything/utils/pagination"
	"recything/utils/storage"

	"gorm.io/gorm"
)

type voucherRepository struct {
	db *gorm.DB
	sb storage.StorageInterface
}

func NewVoucherRepository(db *gorm.DB,sb storage.StorageInterface) entity.VoucherRepositoryInterface {
	return &voucherRepository{
		db: db,
		sb: sb,
	}
}

func (vr *voucherRepository) Create(image *multipart.FileHeader, data entity.VoucherCore) error {
	input := entity.CoreVoucherToModelVoucher(data)

	imageURL, errUpload := vr.sb.Upload(image)
	if errUpload != nil {
		return errUpload
	}

	input.Image = imageURL
	log.Println(input)
	tx := vr.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (vr *voucherRepository) CreateExchangeVoucher(idUser string, data entity.ExchangeVoucherCore) error {
	input := entity.CoreExchangeVoucherToModelExchangeVoucher(data)

	input.IdUser = idUser
	tx := vr.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (vr *voucherRepository) GetAll(page, limit int, search string) ([]entity.VoucherCore, pagination.PageInfo, int, error) {
	dataVouchers := []model.Voucher{}
	offsetInt := (page - 1) * limit

	totalCount, err := vr.GetCount(search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	if search == "" {
		tx := vr.db.Limit(limit).Offset(offsetInt).Order("created_at DESC").Find(&dataVouchers)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if search != "" {
		tx := vr.db.Where("reward_name LIKE ? or point LIKE ? ", "%"+search+"%", "%"+search+"%").Limit(limit).Offset(offsetInt).Order("created_at DESC").Find(&dataVouchers)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	dataResponse := entity.ListModelVoucherToCoreVoucher(dataVouchers)
	paginationInfo := pagination.CalculateData(totalCount, limit, page)

	return dataResponse, paginationInfo, totalCount, nil
}

func (vr *voucherRepository) GetCount(search string) (int, error) {
	var totalCount int64

	if search == "" {
		tx := vr.db.Model(&model.Voucher{}).Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	if search != "" {
		tx := vr.db.Model(&model.Voucher{}).Where("reward_name LIKE ? or point LIKE ? ", "%"+search+"%", "%"+search+"%").Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}

	}
	return int(totalCount), nil
}

func (vr *voucherRepository) GetById(idVoucher string) (entity.VoucherCore, error) {
	dataVouchers := model.Voucher{}

	tx := vr.db.Where("id = ?", idVoucher).First(&dataVouchers)
	if tx.Error != nil {
		return entity.VoucherCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.VoucherCore{}, tx.Error
	}

	result := entity.ModelVoucherToCoreVoucher(dataVouchers)
	return result, nil
}

func (vr *voucherRepository) Update(idVoucher string, image *multipart.FileHeader, data entity.VoucherCore) error {
	input := entity.CoreVoucherToModelVoucher(data)
	dataVoucher := model.Voucher{}

	tx := vr.db.Where("id = ?", idVoucher).First(&dataVoucher)
	if tx.Error != nil {
		return tx.Error
	}

	if image != nil {
		imageURL, errUpload := vr.sb.Upload(image)
		if errUpload != nil {
			return errUpload
		}
		input.Image = imageURL
	} else {
		input.Image = dataVoucher.Image
	}

	tx = vr.db.Where("id = ?", idVoucher).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (vr *voucherRepository) Delete(idVoucher string) error {
	request := model.Voucher{}

	tx := vr.db.Where("id = ?", idVoucher).Delete(&request)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// Exchange Point

func (vr *voucherRepository) GetCountVouchers(filter, search string) (int, error) {

	var totalCount int64
	models := vr.db.Model(&model.ExchangeVoucher{})

	if filter == "" || search == "" {
		tx := models.Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	if search != "" {
		if err := models.
			Joins("JOIN vouchers ON exchange_vouchers.id_voucher = vouchers.id").
			Joins("JOIN users ON exchange_vouchers.id_user = users.id").
			Where("vouchers.reward_name LIKE ? OR users.fullname LIKE ?", "%"+search+"%", "%"+search+"%").
			Count(&totalCount).Error; err != nil {
			return 0, err
		}

	}

	if filter != "" {
		tx := models.Where("status LIKE ?", "%"+filter+"%").Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}

	}
	return int(totalCount), nil
}

func (vr *voucherRepository) GetCountDataVouchers() (helper.CountExchangeVoucher, error) {

	counts := helper.CountExchangeVoucher{}

	tx := vr.db.Model(&model.ExchangeVoucher{}).Count(&counts.TotalCount)
	if tx.Error != nil {
		return counts, tx.Error
	}

	err := vr.db.Model(&model.ExchangeVoucher{}).Where("status LIKE ?", "%"+constanta.TERBARU+"%").Count(&counts.CountNewest).Error
	if err != nil {
		return counts, err
	}
	err = vr.db.Model(&model.ExchangeVoucher{}).Where("status LIKE ?", "%"+constanta.DIPROSES+"%").Count(&counts.CountProcess).Error
	if err != nil {
		return counts, err
	}
	err = vr.db.Model(&model.ExchangeVoucher{}).Where("status LIKE ?", "%"+constanta.SELESAI+"%").Count(&counts.CountDone).Error
	if err != nil {
		return counts, err
	}

	return counts, nil
}

func (vr *voucherRepository) GetAllExchange(page, limit int, search, filter string) ([]entity.ExchangeVoucherCore, pagination.PageInfo, helper.CountExchangeVoucher, error) {
	dataExchange := []model.ExchangeVoucher{}
	offsetInt := (page - 1) * limit
	paginationQuery := vr.db.Limit(limit).Offset(offsetInt)
	counts, _ := vr.GetCountDataVouchers()
	totalCount, err := vr.GetCountVouchers(filter, search)

	if err != nil {
		return nil, pagination.PageInfo{}, counts, err
	}

	if filter != "" {
		tx := paginationQuery.Where("status LIKE ?", "%"+filter+"%").Order("created_at DESC").Find(&dataExchange)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, tx.Error
		}
	}

	if search != "" {
		newCounts := helper.CountExchangeVoucher{}
		tx := vr.db.Model(&model.ExchangeVoucher{}).
			Joins("JOIN vouchers ON exchange_vouchers.id_voucher = vouchers.id").
			Joins("JOIN users ON exchange_vouchers.id_user = users.id").
			Where("(vouchers.reward_name LIKE ? OR users.fullname LIKE ?) AND status LIKE ?", "%"+search+"%", "%"+search+"%", "%"+constanta.TERBARU+"%").
			Count(&newCounts.CountNewest)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, errors.New("error disini")

		}

		tx = vr.db.Model(&model.ExchangeVoucher{}).
			Joins("JOIN vouchers ON exchange_vouchers.id_voucher = vouchers.id").
			Joins("JOIN users ON exchange_vouchers.id_user = users.id").
			Where("(vouchers.reward_name LIKE ? OR users.fullname LIKE ?) AND status LIKE ?", "%"+search+"%", "%"+search+"%", "%"+constanta.DIPROSES+"%").
			Count(&newCounts.CountProcess)

		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, errors.New("error disini")

		}

		tx = vr.db.Model(&model.ExchangeVoucher{}).
			Joins("JOIN vouchers ON exchange_vouchers.id_voucher = vouchers.id").
			Joins("JOIN users ON exchange_vouchers.id_user = users.id").
			Where("(vouchers.reward_name LIKE ? OR users.fullname LIKE ?) AND status LIKE ?", "%"+search+"%", "%"+search+"%", "%"+constanta.SELESAI+"%").
			Count(&newCounts.CountDone)

		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, errors.New("error disini")

		}

		counts.TotalCount = int64(totalCount)
		counts.CountNewest = newCounts.CountNewest
		counts.CountProcess = newCounts.CountProcess
		counts.CountDone = newCounts.CountDone
		tx = paginationQuery.Model(&model.ExchangeVoucher{}).
			Joins("JOIN vouchers ON exchange_vouchers.id_voucher = vouchers.id").
			Joins("JOIN users ON exchange_vouchers.id_user = users.id").
			Where("vouchers.reward_name LIKE ? OR users.fullname LIKE ?", "%"+search+"%", "%"+search+"%").
			Order("created_at DESC").
			Find(&dataExchange)

		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, errors.New("error disini")
		}
	}

	if search == "" && filter == "" {
		tx := paginationQuery.Order("created_at DESC").Find(&dataExchange)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, counts, tx.Error
		}
	}
	dataResponse := []entity.ExchangeVoucherCore{}

	for _, exchange := range dataExchange {
		vr.db.Model(&exchange).Association("Users").Find(&exchange.Users)
		vr.db.Model(&exchange).Association("Vouchers").Find(&exchange.Vouchers)

		exchange.IdUser = exchange.Users.Fullname
		exchange.IdVoucher = exchange.Vouchers.RewardName
		data := entity.ModelExchangeVoucherToCoreExchangeVoucher(exchange)

		dataResponse = append(dataResponse, data)
	}

	paginationInfo := pagination.CalculateData(totalCount, limit, page)

	return dataResponse, paginationInfo, counts, nil

}

func (vr *voucherRepository) GetAllExchangeHistory(userID string) ([]map[string]interface{}, error) {
	dataExchange := []model.ExchangeVoucher{}

	tx := vr.db.Where("id_user", userID).Find(&dataExchange)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataResponse []map[string]interface{}

	for _, exchange := range dataExchange {
		vr.db.Model(&exchange).Association("Vouchers").Find(&exchange.Vouchers)

		exchange.IdVoucher = exchange.Vouchers.RewardName
		point := exchange.Vouchers.Point

		data := entity.ModelExchangeVoucherToMap(exchange,point)

		dataResponse = append(dataResponse, data)
	}

	return dataResponse, nil
}

func (vr *voucherRepository) GetByIdExchange(idExchange string) (entity.ExchangeVoucherCore, error) {
	dataExchange := model.ExchangeVoucher{}

	tx := vr.db.Where("id = ?", idExchange).First(&dataExchange)
	if tx.Error != nil {
		return entity.ExchangeVoucherCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.ExchangeVoucherCore{}, tx.Error
	}

	vr.db.Model(&dataExchange).Association("Users").Find(&dataExchange.Users)
	vr.db.Model(&dataExchange).Association("Vouchers").Find(&dataExchange.Vouchers)

	dataExchange.IdUser = dataExchange.Users.Fullname
	dataExchange.IdVoucher = dataExchange.Vouchers.RewardName

	dataResponse := entity.ModelExchangeVoucherToCoreExchangeVoucher(dataExchange)

	return dataResponse, nil
}

func (vr *voucherRepository) UpdateStatusExchange(id, status string) error {
	dataExchange := model.ExchangeVoucher{}

	errData := vr.db.Where("id = ?", id).First(&dataExchange)
	if errData.Error != nil {
		return errData.Error
	}

	dataExchange.Status = status

	tx := vr.db.Save(&dataExchange)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// Point History 

func (vr *voucherRepository) GetByIdExchangeTransactions(userID,idTransaction string) (map[string]interface{}, error) {
	dataExchange := model.ExchangeVoucher{}

	tx := vr.db.Where("id_user = ? AND id = ?", userID,idTransaction).First(&dataExchange)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, tx.Error
	}

	vr.db.Model(&dataExchange).Association("Vouchers").Find(&dataExchange.Vouchers)

	point := dataExchange.Vouchers.Point
	dataExchange.IdVoucher = dataExchange.Vouchers.RewardName

	dataResponse := entity.ModelExchangeVoucherToMapDetail(dataExchange,point)

	return dataResponse, nil
}
