package entity

import "recything/features/recybot/model"

func CoreRecybotToModelRecybot(recybot RecybotCore) model.Recybot {
	return model.Recybot{
		Category: recybot.Category,
		Question: recybot.Question,
	}
}

func ListCoreRecybotToModelRecybot(recybot []RecybotCore) []model.Recybot {
	list := []model.Recybot{}
	for _, v := range recybot {
		result := CoreRecybotToModelRecybot(v)
		list = append(list, result)
	}
	return list
}

func ModelRecybotToCoreRecybot(recybot model.Recybot) RecybotCore {
	return RecybotCore{
		ID:        recybot.ID,
		Category:  recybot.Category,
		Question:  recybot.Question,
		CreatedAt: recybot.CreatedAt,
		UpdatedAt: recybot.UpdatedAt,
	}
}

func ListModelRecybotToCoreRecybot(recybot []model.Recybot) []RecybotCore {
	list := []RecybotCore{}
	for _, v := range recybot {
		result := ModelRecybotToCoreRecybot(v)
		list = append(list, result)
	}
	return list
}

func RecybotHistoryCoreToModelRecyHistory(recybot RecybotHistories) model.RecybotHistory{
	return model.RecybotHistory{
		Question: recybot.Question,
		Answer: recybot.Answer,
		UserId: recybot.UserId,
	}
}

func ModelRecyHistoryToEntityRecyHistory(recybot model.RecybotHistory)RecybotHistories{
	return RecybotHistories{
		ID: recybot.ID,
		Question: recybot.Question,
		Answer: recybot.Answer,
		UserId: recybot.UserId,
		CreatedAt: recybot.CreatedAt,
	}
}
func ListModelRecyHistoryToEntityRecyHistory(r []model.RecybotHistory)[]RecybotHistories{
	list := []RecybotHistories{}
	for _, v := range r {
		result := ModelRecyHistoryToEntityRecyHistory(v)
		list = append(list, result)
	}
	return list
}
