package service

import (
	"errors"
	"mime/multipart"
	user "recything/features/user/entity"
	"recything/features/voucher/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"recything/utils/validation"
	"time"
)

type voucherService struct {
	voucherRepository entity.VoucherRepositoryInterface
	userRepository    user.UsersRepositoryInterface
}

func NewVoucherService(voucher entity.VoucherRepositoryInterface, user user.UsersRepositoryInterface) entity.VoucherServiceInterface {
	return &voucherService{
		voucherRepository: voucher,
		userRepository:    user,
	}
}

func (vs *voucherService) Create(image *multipart.FileHeader, data entity.VoucherCore) error {

	errEmpty := validation.CheckDataEmpty(data.RewardName, data.Point, data.Description, data.StartDate, data.EndDate)
	if errEmpty != nil {
		return errEmpty
	}

	errDate := validation.ValidateDate(data.StartDate, data.EndDate)
	if errDate != nil {
		return errDate
	}

	errCreate := vs.voucherRepository.Create(image, data)
	if errCreate != nil {
		return errCreate
	}
	return nil
}

func (vs *voucherService) GetAll(page, limit, search string) ([]entity.VoucherCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}
	data, pagnationInfo, count, err := vs.voucherRepository.GetAll(pageInt, limitInt, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return data, pagnationInfo, count, nil
}

func (vs *voucherService) GetById(idVoucher string) (entity.VoucherCore, error) {
	result, err := vs.voucherRepository.GetById(idVoucher)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (vs *voucherService) UpdateData(idVoucher string, image *multipart.FileHeader, data entity.VoucherCore) error {

	errEmpty := validation.CheckDataEmpty(data.RewardName, data.Point, data.Description, data.StartDate, data.EndDate)
	if errEmpty != nil {
		return errEmpty
	}

	errDate := validation.ValidateDate(data.StartDate, data.EndDate)
	if errDate != nil {
		return errDate
	}

	err := vs.voucherRepository.Update(idVoucher, image, data)
	if err != nil {
		return err
	}

	return nil
}

func (vs *voucherService) DeleteData(idVoucher string) error {

	err := vs.voucherRepository.Delete(idVoucher)
	if err != nil {
		return err
	}
	return nil
}

// CreateExchangeVoucher implements entity.VoucherServiceInterface.
func (vs *voucherService) CreateExchangeVoucher(idUser string, data entity.ExchangeVoucherCore) error {
	errEmpty := validation.CheckDataEmpty(data.IdVoucher, data.Phone)
	if errEmpty != nil {
		return errEmpty
	}

	errPhone := validation.PhoneNumber(data.Phone)
	if errPhone != nil {
		return errors.New("error : nomor telepon tidak valid")
	}

	userData, err := vs.userRepository.GetById(idUser)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	voucherData, err := vs.voucherRepository.GetById(data.IdVoucher)
	if err != nil {
		return err
	}

	if userData.Point <= voucherData.Point {
		return errors.New("error : point tidak cukup")
	}

	userData.Point -= voucherData.Point

	// Update user
	err = vs.userRepository.UpdateById(userData.Id, userData)
	if err != nil {
		return errors.New("gagal memperbarui nilai point pengguna")
	}

	loc, err := time.LoadLocation(constanta.ASIABANGKOK)
	if err != nil {
		return err
	}

	data.TimeTransaction = time.Now().In(loc).Format("15:04:05.000")
	err = vs.voucherRepository.CreateExchangeVoucher(idUser, data)
	if err != nil {
		return err
	}
	return nil
}
func (vs *voucherService) GetAllExchange(page, limit, search, filter string) ([]entity.ExchangeVoucherCore, pagination.PageInfo, helper.CountExchangeVoucher, error) {

	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, helper.CountExchangeVoucher{}, err
	}
	dataExchange, pagination, count, err := vs.voucherRepository.GetAllExchange(pageInt, limitInt, search, filter)
	if err != nil {
		return nil, pagination, count, err

	}

	return dataExchange, pagination, count, nil
}

func (vs *voucherService) GetByIdExchange(idExchange string) (entity.ExchangeVoucherCore, error) {

	dataExchange, errGet := vs.voucherRepository.GetByIdExchange(idExchange)
	if errGet != nil {
		return entity.ExchangeVoucherCore{}, errGet
	}

	return dataExchange, nil
}

func (vs *voucherService) UpdateStatusExchange(id string, status string) error {

	errEmpty := validation.CheckDataEmpty(status)
	if errEmpty != nil {
		return errors.New("error : status harus diisi")
	}

	statusEqual, errEqual := validation.CheckEqualData(status, constanta.Status_Exchange)
	if errEqual != nil {
		return errors.New("error : input status tidak valid")
	}

	dataStatus, err := vs.voucherRepository.GetByIdExchange(id)
	if err != nil {
		return err
	}

	if dataStatus.Status == "selesai" {
		return errors.New("error : status sudah selesai")
	}

	if dataStatus.Status == "diproses" && statusEqual == "perlu ditinjau" {
		return errors.New("error : status sudah diproses ")
	}

	if dataStatus.Status == statusEqual {
		return errors.New("error : status sudah diperbarui tidak bisa diubah ")
	}

	err = vs.voucherRepository.UpdateStatusExchange(id, statusEqual)
	if err != nil {
		return err
	}

	return nil
}
