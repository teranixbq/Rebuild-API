package repository

import (
	"errors"
	"recything/features/recybot/entity"
	"recything/features/recybot/model"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type recybotRepository struct {
	db *gorm.DB
}

// GetAllHistory implements entity.RecybotRepositoryInterface.
func (rb *recybotRepository) GetAllHistory(userId string) ([]entity.RecybbotHistories, error) {
	var hst []model.RecybotHistory

	err := rb.db.Where("user_id = ?", userId).Find(&hst).Error
	if err != nil {
		return []entity.RecybbotHistories{}, err
	}
	result:= entity.ListModelRecyHistoryToEntityRecyHistory(hst)
	return result, nil
}

func NewRecybotRepository(db *gorm.DB) entity.RecybotRepositoryInterface {
	return &recybotRepository{
		db: db,
	}
}

func (rb *recybotRepository) Create(recybot entity.RecybotCore) (entity.RecybotCore, error) {
	input := entity.CoreRecybotToModelRecybot(recybot)

	tx := rb.db.Create(&input)
	if tx.Error != nil {
		return entity.RecybotCore{}, tx.Error
	}

	result := entity.ModelRecybotToCoreRecybot(input)
	return result, nil
}

func (rb *recybotRepository) FindAll(page, limit int, filter, search string) ([]entity.RecybotCore, pagination.PageInfo, helper.CountPrompt, error) {
	dataRecybots := []model.Recybot{}

	offsetInt := (page - 1) * limit
	var totalCount int64
	counts, err := rb.GetCountAllData(search, filter)
	if err != nil {
		return nil, pagination.PageInfo{}, helper.CountPrompt{}, err
	}

	paginationQuery := rb.db.Limit(limit).Offset(offsetInt)
	if filter == "" || search == "" {
		totalCount = counts.TotalCount
		tx := paginationQuery.Find(&dataRecybots)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, helper.CountPrompt{}, err
		}
	}

	if filter != "" {
		if strings.Contains(filter, constanta.ANORGANIC) {
			totalCount = counts.CountAnorganic
		}

		if strings.Contains(filter, constanta.ORGANIC) {
			totalCount = counts.CountOrganic
		}

		tx := paginationQuery.Where("category = ?", filter).Find(&dataRecybots)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, helper.CountPrompt{}, err
		}
	}

	if search != "" {
		totalCount = counts.TotalCount
		tx := paginationQuery.Where("question LIKE ? ", "%"+search+"%").Find(&dataRecybots)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, helper.CountPrompt{}, err
		}
	}

	result := entity.ListModelRecybotToCoreRecybot(dataRecybots)
	paginationInfo := pagination.CalculateData(int(totalCount), limit, page)
	return result, paginationInfo, counts, nil

}

func (rb *recybotRepository) GetCountAllData(search string, filter string) (helper.CountPrompt, error) {

	counts := helper.CountPrompt{}

	tx := rb.db.Model(&model.Recybot{}).Select(
		"COUNT(CASE WHEN category = ? THEN 1 END) AS CountAnorganic, "+
			"COUNT(CASE WHEN category = ? THEN 1 END) AS CountOrganic, "+
			"COUNT(CASE WHEN category = ? THEN 1 END) AS CountLimitation, "+
			"COUNT(CASE WHEN category = ? THEN 1 END) AS CountInformation",
		constanta.ANORGANIC, constanta.ORGANIC, constanta.LIMITATION, constanta.INFORMATION).
		Scan(&counts)
	if tx.Error != nil {
		return counts, tx.Error
	}

	if search != "" {
		if filter != "" {
			tx = rb.db.Model(&model.Recybot{}).Where("category = ? AND question LIKE ? ", filter, "%"+search+"%").Count(&counts.TotalCount)
			if tx.Error != nil {
				return counts, tx.Error
			}

			// if filter != constanta.ANORGANIC {

			// }

			tx := rb.db.Model(&model.Recybot{}).
				Select("COUNT(CASE WHEN category = ? AND question LIKE ? THEN 1 ELSE NULL END) AS CountAnorganic, "+
					"COUNT(CASE WHEN category = ? AND question LIKE ? THEN 1 ELSE NULL END) AS CountOrganic, "+
					"COUNT(CASE WHEN category = ? AND question LIKE ? THEN 1 ELSE NULL END) AS CountLimitation, "+
					"COUNT(CASE WHEN category = ? AND question LIKE ? THEN 1 ELSE NULL END) AS CountInformation",
					constanta.ANORGANIC, "%"+search+"%", constanta.ORGANIC, "%"+search+"%",
					constanta.LIMITATION, "%"+search+"%", constanta.INFORMATION, "%"+search+"%").
				Scan(&counts)
			if tx.Error != nil {
				return counts, tx.Error
			}

			return counts, nil

		}

		if filter == "" {
			tx := rb.db.Model(&model.Recybot{}).Select(
				"COUNT(CASE WHEN category = ? THEN 1 END) AS CountAnorganic, "+
					"COUNT(CASE WHEN category = ? THEN 1 END) AS CountOrganic, "+
					"COUNT(CASE WHEN category = ? THEN 1 END) AS CountLimitation, "+
					"COUNT(CASE WHEN category = ? THEN 1 END) AS CountInformation",
				constanta.ANORGANIC, constanta.ORGANIC, constanta.LIMITATION, constanta.INFORMATION).
				Where("question LIKE ?", "%"+search+"%").Scan(&counts)

			tx = rb.db.Model(&model.Recybot{}).Where("question LIKE ? ", "%"+search+"%").Count(&counts.TotalCount)
			if tx.Error != nil {
				return counts, tx.Error
			}
		}

	}

	if search == "" {
		tx := rb.db.Model(&model.Recybot{}).Count(&counts.TotalCount)
		if tx.Error != nil {
			return counts, tx.Error
		}
	}

	return counts, nil
}

func (rb *recybotRepository) GetAll() ([]entity.RecybotCore, error) {
	dataRecybots := []model.Recybot{}

	tx := rb.db.Find(&dataRecybots)
	if tx.Error != nil {
		return []entity.RecybotCore{}, tx.Error
	}

	result := entity.ListModelRecybotToCoreRecybot(dataRecybots)
	return result, nil
}

func (rb *recybotRepository) GetById(idData string) (entity.RecybotCore, error) {
	dataRecybots := model.Recybot{}

	tx := rb.db.Where("id = ?", idData).First(&dataRecybots)
	if tx.Error != nil {
		return entity.RecybotCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.RecybotCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.ModelRecybotToCoreRecybot(dataRecybots)
	return result, nil
}

func (rb *recybotRepository) Update(idData string, recybot entity.RecybotCore) (entity.RecybotCore, error) {
	data := entity.CoreRecybotToModelRecybot(recybot)

	tx := rb.db.Where("id = ?", idData).Updates(&data)
	if tx.Error != nil {
		return entity.RecybotCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.RecybotCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.ModelRecybotToCoreRecybot(data)
	return result, nil
}

func (rb *recybotRepository) Delete(idData string) error {
	data := model.Recybot{}

	tx := rb.db.Where("id = ?", idData).Delete(&data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}

// InsertHistory implements entity.RecybotRepositoryInterface.
func (rb *recybotRepository) InsertHistory(userId, answer, question string) error {
	var history model.RecybotHistory
	history.ID = uuid.New().String()
	history.Question = question
	history.UserId = userId
	history.CreatedAt = time.Time{}
	history.DeletedAt = gorm.DeletedAt{}
	// history.teks = question + answer
	if strings.Contains(answer, "Maaf") {
		history.Answer = ""
	} else {
		history.Answer = answer
	}

	err := rb.db.Create(&history).Error
	if err != nil {
		return err
	}
	return nil
}
