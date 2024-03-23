package service

import (
	//"errors"

	"errors"
	"mime/multipart"
	user "recything/features/user/entity"
	"recything/features/voucher/entity"
	"recything/mocks"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"

	// "time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

var dataVouchers = []entity.VoucherCore{
	{Id: "1", RewardName: "Dana 10k", Point: 10000, Description: "Voucher 10k", StartDate: "2023-12-16", EndDate: "2023-12-30"},
	{Id: "2", RewardName: "Dana 20k", Point: 20000, Description: "Voucher 20k", StartDate: "2023-12-17", EndDate: "2023-12-30"},
	{Id: "3", RewardName: "Dana 30k", Point: 30000, Description: "Voucher 30k", StartDate: "2023-12-18", EndDate: "2023-12-30"},
}

var dataVoucher = entity.VoucherCore{
	Id:          "1",
	RewardName:  "Dana 10k",
	Point:       10000,
	Description: "Voucher 10k",
	StartDate:   "2023-12-16",
	EndDate:     "2023-12-30",
}

var dataExchanges = []entity.ExchangeVoucherCore{
	{Id: "11", IdUser: "123", IdVoucher: "1", Phone: "082298673422", Status: "terbaru"},
	{Id: "22", IdUser: "222", IdVoucher: "2", Phone: "082298673455", Status: "diproses"},
	{Id: "33", IdUser: "333", IdVoucher: "3", Phone: "082298673477", Status: "selesai"},
}

var dataExchange = entity.ExchangeVoucherCore{
	Id: "11", IdUser: "123", IdVoucher: "1", Phone: "082298673422",
}

var dataUser = user.UsersCore{
	Id:    "123",
	Point: 500000,
}

func TestCreate(t *testing.T) {
	t.Run("Succes Create", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		adminService := NewVoucherService(mockVoucher, mockUser)

		image := &multipart.FileHeader{
			Filename: "testfile.png",
		}

		requestBody := entity.VoucherCore{
			RewardName:  "Dana 10k",
			Point:       10000,
			Description: "Voucher 10k",
			StartDate:   "2023-12-29",
			EndDate:     "2023-12-30",
		}

		mockVoucher.On("Create", mock.AnythingOfType("*multipart.FileHeader"), requestBody).Return(nil)

		err := adminService.Create(image, requestBody)

		assert.NoError(t, err)
		mockVoucher.AssertExpectations(t)

	})

	t.Run("Data Empty", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		adminService := NewVoucherService(mockVoucher, mockUser)

		requestBody := entity.VoucherCore{
			RewardName:  "Dana 10k",
			Point:       10000,
			Description: "",
			StartDate:   "",
			EndDate:     "",
		}

		err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("Invalid Date", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		adminService := NewVoucherService(mockVoucher, mockUser)

		requestBody := entity.VoucherCore{
			RewardName:  "Dana 10k",
			Point:       10000,
			Description: "Voucher 10k",
			StartDate:   "2023-12-16",
			EndDate:     "2021-12-30",
		}

		err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("Error Repository", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		image := &multipart.FileHeader{}
		requestBody := entity.VoucherCore{
			RewardName:  "TestReward",
			Point:       100,
			Description: "TestDescription",
			StartDate:   "2023-12-21",
			EndDate:     "2023-12-31",
		}

		mockVoucher.On("Create", mock.Anything, mock.Anything).Return(errors.New("mocked repository error"))

		err := voucherService.Create(image, requestBody)

		assert.Error(t, err)
		mockVoucher.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Success GetAll", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		adminService := NewVoucherService(mockVoucher, mockUser)

		mockVoucher.On("GetAll", 1, 10, "").Return(dataVouchers, pagination.PageInfo{}, len(dataVouchers), nil)

		result, pagination, count, err := adminService.GetAll("", "", "")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, pagination)
		assert.NotNil(t, count)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("Wrong Limit", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		adminService := NewVoucherService(mockVoucher, mockUser)

		result, pagination, count, err := adminService.GetAll("", "50", "")

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Empty(t, pagination)
		assert.Empty(t, count)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("Error Repository", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		adminService := NewVoucherService(mockVoucher, mockUser)

		mockVoucher.On("GetAll", 1, 10, "").Return(dataVouchers, pagination.PageInfo{}, len(dataVouchers), errors.New(constanta.ERROR))

		result, pagination, count, err := adminService.GetAll("", "", "")

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Empty(t, pagination)
		assert.Empty(t, count)
		mockVoucher.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Success Get", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		adminService := NewVoucherService(mockVoucher, mockUser)

		voucherID := "1"
		mockVoucher.On("GetById", voucherID).Return(dataVoucher, nil)

		result, err := adminService.GetById(voucherID)

		assert.NoError(t, err)
		assert.Equal(t, voucherID, dataVoucher.Id)
		assert.NotNil(t, result)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		voucherID := "3"
		mockVoucher.On("GetById", voucherID).Return(entity.VoucherCore{}, errors.New(constanta.ERROR))

		result, err := voucherService.GetById(voucherID)

		assert.Error(t, err)
		assert.NotEqual(t, voucherID, dataVoucher.Id)
		assert.Empty(t, result)
		mockVoucher.AssertExpectations(t)
	})

}

func TestUpdateData(t *testing.T) {
	t.Run("Succes Create", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		image := &multipart.FileHeader{
			Filename: "testfile.png",
		}

		requestBody := entity.VoucherCore{
			RewardName:  "Dana 10k",
			Point:       10000,
			Description: "Voucher 10k",
			StartDate:   "2023-12-29",
			EndDate:     "2023-12-30",
		}

		voucherID := "1"

		mockVoucher.On("Update", voucherID, mock.AnythingOfType("*multipart.FileHeader"), requestBody).Return(nil)

		err := voucherService.UpdateData(voucherID, image, requestBody)

		assert.NoError(t, err)
		assert.Equal(t, voucherID, dataVoucher.Id)
		mockVoucher.AssertExpectations(t)

	})

	t.Run("Data Empty", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		image := &multipart.FileHeader{
			Filename: "testfile.png",
		}

		requestBody := entity.VoucherCore{
			RewardName:  "Dana 10k",
			Point:       10000,
			Description: "",
			StartDate:   "",
			EndDate:     "",
		}

		err := voucherService.UpdateData("", image, requestBody)

		assert.Error(t, err)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("Invalid Date", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		image := &multipart.FileHeader{
			Filename: "testfile.png",
		}

		requestBody := entity.VoucherCore{
			RewardName:  "Dana 10k",
			Point:       10000,
			Description: "Voucher 10k",
			StartDate:   "2023-12-01",
			EndDate:     "2023-12-30",
		}

		err := voucherService.UpdateData("", image, requestBody)

		assert.Error(t, err)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("Error Repository", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		image := &multipart.FileHeader{}
		requestBody := entity.VoucherCore{
			RewardName:  "TestReward",
			Point:       100,
			Description: "TestDescription",
			StartDate:   "2023-12-21",
			EndDate:     "2023-12-31",
		}

		mockVoucher.On("Update", "", mock.AnythingOfType("*multipart.FileHeader"), requestBody).Return(errors.New(constanta.ERROR))

		err := voucherService.UpdateData("", image, requestBody)

		assert.Error(t, err)
		mockVoucher.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Succes Delete", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		voucherID := "1"

		mockVoucher.On("Delete", voucherID).Return(nil)

		err := voucherService.DeleteData(voucherID)

		assert.NoError(t, err)
		assert.Equal(t, voucherID, dataVoucher.Id)
		mockVoucher.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		voucherID := "2"

		mockVoucher.On("Delete", voucherID).Return(errors.New(constanta.ERROR))

		err := voucherService.DeleteData(voucherID)

		assert.Error(t, err)
		assert.NotEqual(t, voucherID, dataVoucher.Id)
		mockVoucher.AssertExpectations(t)
	})
}

func TestCreateExchangeVoucher(t *testing.T) {
	t.Run("Data Empty", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		requestBody := entity.ExchangeVoucherCore{
			IdVoucher: "1", Phone: "",
		}

		err := voucherService.CreateExchangeVoucher("", requestBody)

		assert.Error(t, err)
		mockVoucher.AssertExpectations(t)
	})
	t.Run("Wrong Phone", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		requestBody := entity.ExchangeVoucherCore{
			IdVoucher: "1", Phone: "111111",
		}
		err := voucherService.CreateExchangeVoucher("", requestBody)

		assert.Error(t, err)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		mockUser := new(mocks.UsersRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		requestBody := entity.ExchangeVoucherCore{
			IdVoucher: "1", Phone: "082298567540",
		}
		userID := "1235"
		mockUser.On("GetById", userID).Return(dataUser, errors.New(constanta.ERROR_RECORD_NOT_FOUND))
		resultUser, errUser := mockUser.GetById(userID)
		err := voucherService.CreateExchangeVoucher(userID, requestBody)

		assert.Error(t, err)

		assert.NotEqual(t, userID, dataUser.Id)
		assert.Error(t, errUser)
		assert.NotNil(t, resultUser)
		mockVoucher.AssertExpectations(t)

	})

	t.Run("Voucher Not Found", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		requestBody := entity.ExchangeVoucherCore{
			IdVoucher: "22", Phone: "082298567540",
		}

		mockUser.On("UpdateById", mock.AnythingOfType("string"), mock.Anything).Return(errors.New(constanta.ERROR))
		mockUser.On("GetById", mock.AnythingOfType("string")).Return(dataUser, nil)
		mockVoucher.On("GetById", mock.AnythingOfType("string")).Return(dataVoucher, nil)

		err := voucherService.CreateExchangeVoucher("", requestBody)

		assert.Error(t, err)
		assert.NotEqual(t, requestBody.IdVoucher, dataVoucher.Id)
		mockUser.AssertExpectations(t)
		mockVoucher.AssertExpectations(t)
	})

}

func TestGetAllExchange(t *testing.T) {
	t.Run("Succes Get All", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		mockVoucher.On("GetAllExchange", 1, 10, "", "").Return(dataExchanges, pagination.PageInfo{}, helper.CountExchangeVoucher{}, nil)

		result, pagination, count, err := voucherService.GetAllExchange("", "", "", "")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, pagination)
		assert.NotNil(t, count)

		mockVoucher.AssertExpectations(t)

	})

	t.Run("Wrong Limit", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		result, pagination, count, err := voucherService.GetAllExchange("", "50", "", "")

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Empty(t, pagination)
		assert.Empty(t, count)

		mockVoucher.AssertExpectations(t)

	})

	t.Run("Error Repository", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		mockVoucher.On("GetAllExchange", 1, 10, "", "").Return(dataExchanges, pagination.PageInfo{}, helper.CountExchangeVoucher{}, errors.New(constanta.ERROR))

		result, pagination, count, err := voucherService.GetAllExchange("", "", "", "")

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Empty(t, pagination)
		assert.Empty(t, count)

		mockVoucher.AssertExpectations(t)

	})
}

func TestGetByIdExchange(t *testing.T) {
	t.Run("Succes Get By Id", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		exchangeID := "11"
		mockVoucher.On("GetByIdExchange", exchangeID).Return(dataExchange, nil)

		result, err := voucherService.GetByIdExchange(exchangeID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, exchangeID, dataExchange.Id)
		mockVoucher.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		exchangeID := "112"
		mockVoucher.On("GetByIdExchange", exchangeID).Return(entity.ExchangeVoucherCore{}, errors.New(constanta.ERROR))

		result, err := voucherService.GetByIdExchange(exchangeID)

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.NotEqual(t, exchangeID, dataExchange.Id)
		mockVoucher.AssertExpectations(t)
	})
}

func TestUpdateTStatusExchange(t *testing.T) {
	t.Run("Succes Update", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		exchangeID := "11"

		status := "diproses"

		mockVoucher.On("GetByIdExchange", exchangeID).Return(dataExchanges[0], nil)
		mockVoucher.On("UpdateStatusExchange", exchangeID, status).Return(nil)

		_, errVoucher := voucherService.GetByIdExchange(exchangeID)
		err := voucherService.UpdateStatusExchange(exchangeID, status)

		assert.NoError(t, errVoucher)
		assert.NoError(t, err)

		mockVoucher.AssertExpectations(t)
	})

	t.Run("Data Empty", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		status := ""

		err := voucherService.UpdateStatusExchange("", status)

		assert.Error(t, err)

		mockVoucher.AssertExpectations(t)
	})

	t.Run("Status Input Invalid", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		exchangeID := "11"

		status := "terverifikasi"

		err := voucherService.UpdateStatusExchange(exchangeID, status)

		assert.Error(t, err)
		assert.Equal(t, exchangeID, dataExchange.Id)
		assert.NotEqualValues(t, []string{"diproses", "selesai"}, status)

		mockVoucher.AssertExpectations(t)
	})

	t.Run("Voucher Not Found", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		exchangeID := "121"

		status := "diproses"

		mockVoucher.On("GetByIdExchange", exchangeID).Return(dataExchange, errors.New(constanta.ERROR))

		_, errVoucher := voucherService.GetByIdExchange(exchangeID)
		err := voucherService.UpdateStatusExchange(exchangeID, status)

		assert.Error(t, errVoucher)
		assert.Error(t, err)
		assert.NotEqual(t, exchangeID, dataExchange.Id)

		mockVoucher.AssertExpectations(t)
	})

	t.Run("Already Action", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		exchangeID := "22"
		status := "diproses"

		mockVoucher.On("GetByIdExchange", exchangeID).Return(dataExchanges[1], nil)

		err := voucherService.UpdateStatusExchange(exchangeID, status)

		assert.Error(t, err)
		assert.Equal(t,exchangeID,dataExchanges[1].Id)
		mockVoucher.AssertExpectations(t)

	})

	t.Run("Already Done", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		exchangeID := "22"
		status := "selesai"

		mockVoucher.On("GetByIdExchange", exchangeID).Return(dataExchanges[2], nil)
		// 

		err := voucherService.UpdateStatusExchange(exchangeID, status)

		assert.Error(t, err)
		assert.Equal(t,exchangeID,dataExchanges[1].Id)
		mockVoucher.AssertExpectations(t)

	})

	t.Run("Error Repository", func(t *testing.T) {
		mockUser := new(mocks.UsersRepositoryInterface)
		mockVoucher := new(mocks.VoucherRepositoryInterface)
		voucherService := NewVoucherService(mockVoucher, mockUser)

		exchangeID := "11"

		status := "diproses"

		mockVoucher.On("GetByIdExchange", exchangeID).Return(dataExchanges[0], nil)
		mockVoucher.On("UpdateStatusExchange", exchangeID, status).Return(errors.New(constanta.ERROR))

		_, errVoucher := voucherService.GetByIdExchange(exchangeID)
		err := voucherService.UpdateStatusExchange(exchangeID, status)

		assert.NoError(t, errVoucher)
		assert.Error(t, err)

		mockVoucher.AssertExpectations(t)
	})

}
