package response

import "recything/features/faq/entity"

func FaqsCoreToFaqsResponse(data entity.FaqCore ) FaqResponse {
	return FaqResponse{
		Id: data.Id,
		Title: data.Title,
		Description : data.Description,
	}
}

func FaqsCoreToResponseFaqsList(dataCore []entity.FaqCore) []FaqResponse {
	var result []FaqResponse
	for _, v := range dataCore {
		result = append(result, FaqsCoreToFaqsResponse(v))
	}
	return result
}