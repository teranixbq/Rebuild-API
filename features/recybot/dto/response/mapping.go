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

func CoreRecybotHistoryToResponse(recybot entity.RecybotHistories) RecybotHistoryResponse {
	return RecybotHistoryResponse{
		UserId:   recybot.UserId,
		Question: recybot.Question,
		Answer:   recybot.Answer,
	}
}

func ListCoreRecybotHistoryToResponse(recybot []entity.RecybotHistories) []RecybotHistoryResponse{
	list := []RecybotHistoryResponse{}
	for _,v := range recybot{
		result := CoreRecybotHistoryToResponse(v)
		list = append(list, result)
	}
	return list
}