package service

import (
	"errors"
	dropPoint "recything/features/drop-point/entity"
	trashCategory "recything/features/trash_category/entity"
	trashExchange "recything/features/trash_exchange/entity"
	user "recything/features/user/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"recything/utils/validation"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type trashExchangeService struct {
	trashExchangeRepository trashExchange.TrashExchangeRepositoryInterface
	dropPointRepository     dropPoint.DropPointRepositoryInterface
	userRepository          user.UsersRepositoryInterface
	trashCategoryRepository trashCategory.TrashCategoryRepositoryInterface
}

func NewTrashExchangeService(trashExchange trashExchange.TrashExchangeRepositoryInterface, dropPoint dropPoint.DropPointRepositoryInterface, user user.UsersRepositoryInterface, trashCategory trashCategory.TrashCategoryRepositoryInterface) trashExchange.TrashExchangeServiceInterface {
	return &trashExchangeService{
		trashExchangeRepository: trashExchange,
		dropPointRepository:     dropPoint,
		userRepository:          user,
		trashCategoryRepository: trashCategory,
	}
}

// CreateTrashExchange implements entity.TrashExchangeServiceInterface.
func (tes *trashExchangeService) CreateTrashExchange(data trashExchange.TrashExchangeCore) (trashExchange.TrashExchangeCore, error) {
	
	data.Id = helper.GenerateRandomID("PS", 5)
	errEmpty := validation.CheckDataEmpty(data.Name, data.EmailUser, data.DropPointName)
	if errEmpty != nil {
		return trashExchange.TrashExchangeCore{}, errEmpty
	}

	user, err := tes.userRepository.FindByEmail(data.EmailUser)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, errors.New("pengguna dengan email tersebut tidak ditemukan")
	}

	dropPoint, err := tes.dropPointRepository.GetDropPointByName(data.DropPointName)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, errors.New("nama drop point tidak ditemukan")
	}
	data.DropPointId = dropPoint.Id

	totalPoints := 0
	totalUnits := 0.0
	var details []trashExchange.TrashExchangeDetailCore
	for _, detail := range data.TrashExchangeDetails {

		errEmptyDetail := validation.CheckDataEmpty(detail.TrashType, detail.Amount)
		if errEmptyDetail != nil {
			return trashExchange.TrashExchangeCore{}, errEmptyDetail
		}

		titleCase := cases.Title(language.Indonesian)
		detail.TrashType = titleCase.String(detail.TrashType)
		trashCategory, err := tes.trashCategoryRepository.GetByType(detail.TrashType)
		if err != nil {
			return trashExchange.TrashExchangeCore{}, errors.New("kategori sampah tidak ditemukan")
		}

		detail.Unit = trashCategory.Unit

		detail.TotalPoints = int(float64(detail.Amount) * float64(trashCategory.Point))
		totalPoints += detail.TotalPoints

		totalUnits += detail.Amount
		details = append(details, detail)
	}

	reductionPercentage := 0.10
	reductionAmount := int(float64(totalPoints) * reductionPercentage)
	reducedPoints := totalPoints - reductionAmount

	data.TotalIncome = reductionAmount
	data.TotalPoint = reducedPoints
	data.TotalUnit = totalUnits

	user.Point += data.TotalPoint
	// Update user
	err = tes.userRepository.UpdateById(user.Id, user)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, errors.New("gagal memperbarui nilai point pengguna")
	}
	loc, err := time.LoadLocation(constanta.ASIABANGKOK)
	if err != nil {
		return trashExchange.TrashExchangeCore{},err
	}

	data.CreatedAt = time.Now().In(loc)
	result, err := tes.trashExchangeRepository.CreateTrashExchange(data)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, errors.New("gagal menyimpan data trash exchange")
	}

	for _, detail := range details {
		detail.TrashExchangeId = result.Id
		_, err := tes.trashExchangeRepository.CreateTrashExchangeDetails(detail)
		if err != nil {
			return trashExchange.TrashExchangeCore{}, errors.New("gagal menyimpan data trash exchange detail")
		}
	}

	return result, nil
}

// DeleteTrashExchangeById implements entity.TrashExchangeServiceInterface.
func (tes *trashExchangeService) DeleteTrashExchangeById(id string) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	err := tes.trashExchangeRepository.DeleteTrashExchangeById(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllTrashExchange implements entity.TrashExchangeServiceInterface.
func (tes *trashExchangeService) GetAllTrashExchange(page, limit, search string) ([]trashExchange.TrashExchangeCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateTypePaginationParameter(limit, page)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	pageValid, limitValid := validation.ValidateCountLimitAndPage(pageInt, limitInt)

	dropPointCores, pageInfo, count, err := tes.trashExchangeRepository.GetAllTrashExchange(pageValid, limitValid, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return dropPointCores, pageInfo, count, nil
}

// GetTrashExchangeById implements entity.TrashExchangeServiceInterface.
func (tes *trashExchangeService) GetTrashExchangeById(id string) (trashExchange.TrashExchangeCore, error) {
	if id == "" {
		return trashExchange.TrashExchangeCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	idtrashExchange, err := tes.trashExchangeRepository.GetTrashExchangeById(id)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, err
	}

	return idtrashExchange, err
}
