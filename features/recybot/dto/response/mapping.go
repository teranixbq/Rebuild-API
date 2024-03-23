package response

import "recything/features/recybot/entity"

func CoreRecybotToResponRecybot(recybot entity.RecybotCore) RecybotResponse {
	return RecybotResponse{
		ID:        recybot.ID,
		Category:  recybot.Category,
		Question:  recybot.Question,
		CreatedAt: recybot.CreatedAt,
	}
}

func ListCoreRecybotToCoreRecybot(recybot []entity.RecybotCore) []RecybotResponse {
	list := []RecybotResponse{}
	for _, v := range recybot {
		result := CoreRecybotToResponRecybot(v)
		list = append(list, result)
	}
	return list
}
