package response

import "recything/features/community/entity"

func CoreCommunityToResponCommunity(data entity.CommunityCore) CommunityResponse {
	return CommunityResponse{
		Id:        data.Id,
		Name:      data.Name,
		Image:     data.Image,
		Location:  data.Location,
		CreatedAt: data.CreatedAt,
	}
}

func CoreCommunityToResponCommunityForDetails(data entity.CommunityCore) CommunityResponseForDetails {
	return CommunityResponseForDetails{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
		MaxMembers:  data.MaxMembers,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt,
	}
}

func ListCoreCommunityToResponseCommunity(data []entity.CommunityCore) []CommunityResponse {
	list := []CommunityResponse{}
	for _, v := range data {
		result := CoreCommunityToResponCommunity(v)
		list = append(list, result)
	}
	return list
}

func EventCoreToEventResponse(event entity.CommunityEventCore) EventResponse {
	return EventResponse{
		Id:          event.Id,
		CommunityId: event.CommunityId,
		Title:       event.Title,
		Quota:       event.Quota,
		Date:        event.Date,
		Status:      event.Status,
		Image:       event.Image,
	}
}

func ListEventCoreToListEventRessponse(event []entity.CommunityEventCore) []EventResponse {
	list := []EventResponse{}
	for _, v := range event {
		eventData := EventCoreToEventResponse(v)
		list = append(list, eventData)
	}
	return list
}

func EventCoreToEventResponseDetail(event entity.CommunityEventCore) EventResponseDetail {
	return EventResponseDetail{
		Id:          event.Id,
		CommunityId: event.CommunityId,
		Title:       event.Title,
		Image:       event.Image,
		Description: event.Description,
		Location:    event.Location,
		MapLink:     event.MapLink,
		FormLink:    event.FormLink,
		Quota:       event.Quota,
		Date:        event.Date,
		Status:      event.Status,
	}
}
