package request

import "recything/features/recybot/entity"

func ManageRequestRecybotToCoreRecybot(data RecybotManageRequest) entity.RecybotCore {
	return entity.RecybotCore{
		Category: data.Category,
		Question: data.Question,
	}
}

func ListRequestRecybotToCoreRecybot(data []RecybotManageRequest) []entity.RecybotCore {
	list := []entity.RecybotCore{}
	for _, v := range data {
		result := ManageRequestRecybotToCoreRecybot(v)
		list = append(list, result)
	}
	return list
}

func RequestRecybotToCoreRecybot(data RecybotRequest) entity.RecybotCore {
	return entity.RecybotCore{
		Question: data.Question,
	}
}
