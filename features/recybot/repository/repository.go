package repository

import (
	"errors"
	"log"
	"recything/app/database"
	"recything/features/recybot/entity"
	"recything/features/recybot/model"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"strings"

	"gorm.io/gorm"
)

type recybotRepository struct {
	db  *gorm.DB
	rdb *database.Redis
}


func NewRecybotRepository(db *gorm.DB, rdb *database.Redis) entity.RecybotRepositoryInterface {
	return &recybotRepository{
		db:  db,
		rdb: rdb,
	}
}

func (rb *recybotRepository) Create(recybot entity.RecybotCore) (entity.RecybotCore, error) {
	input := entity.CoreRecybotToModelRecybot(recybot)

	tx := rb.db.Create(&input)
	if tx.Error != nil {
		return entity.RecybotCore{}, tx.Error
	}

	result := entity.ModelRecybotToCoreRecybot(input)

	go func() {
		errDel := rb.rdb.DelString("prompt")
		if errDel != nil {
			log.Println(errDel)
		}

		resultGet, errGet := rb.GetAll()
		if errGet != nil {
			log.Println(errGet)
		}

		errSet := rb.rdb.SetString("prompt", resultGet)
		if errSet != nil {
			log.Println(errSet)
		}
	}()

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

	errRDB := rb.rdb.GetString("prompt", &dataRecybots)
	if errRDB != nil {
		tx := rb.db.Find(&dataRecybots)
		if tx.Error != nil {
			return []entity.RecybotCore{}, tx.Error
		}

		errSet := rb.rdb.SetString("prompt", dataRecybots)
		if errSet != nil {
			return nil, errSet
		}

		log.Println("data from db")
		result := entity.ListModelRecybotToCoreRecybot(dataRecybots)

		return result, nil
	}

	log.Println("data from redis")
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


func (rb *recybotRepository) GetAllHistory(userId string) ([]entity.RecybotHistories, error) {
	dataHistory := []model.RecybotHistory{}

	// Implement redis history
	errGet := rb.rdb.GetString("history."+userId, &dataHistory)
	if errGet != nil {
		err := rb.db.Where("user_id = ?", userId).Find(&dataHistory).Error
		if err != nil {
			return []entity.RecybotHistories{}, err
		}

		errSet := rb.rdb.SetString("history."+userId,dataHistory)
		if errSet != nil {
			return nil,errSet
		}
		
		result := entity.ListModelRecyHistoryToEntityRecyHistory(dataHistory)
		return result, nil
	}

	log.Println("Data from redis history")
	result := entity.ListModelRecyHistoryToEntityRecyHistory(dataHistory)
	return result, nil
}

func (rb *recybotRepository) InsertHistory(history entity.RecybotHistories) error {
	input := entity.RecybotHistoryCoreToModelRecyHistory(history)

	if strings.Contains(input.Answer, "Maaf") {
		input.Answer = ""
	}

	err := rb.db.Create(&input).Error
	if err != nil {
		return err
	}

	go func() {
		errDel := rb.rdb.DelString("history."+history.UserId)
		if errDel != nil {
			log.Println(errDel)
		}

		resultGet, errGet := rb.GetAllHistory(history.UserId)
		if errGet != nil {
			log.Println(errGet)
		}

		errSet := rb.rdb.SetString("history."+history.UserId, resultGet)
		if errSet != nil {
			log.Println(errSet)
		}
	}()

	return nil
}
