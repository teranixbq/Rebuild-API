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

func ModelRecyHistoryToEntityRecyHistory(r model.RecybotHistory)RecybbotHistories{
	return RecybbotHistories{
		Question: r.Question,
		Answer: r.Answer,
	}
}
func ListModelRecyHistoryToEntityRecyHistory(r []model.RecybotHistory)[]RecybbotHistories{
	list := []RecybbotHistories{}
	for _, v := range r {
		result := ModelRecyHistoryToEntityRecyHistory(v)
		list = append(list, result)
	}
	return list
}
