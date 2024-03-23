package entity

import "recything/features/community/model"

func CoreCommunityToModelCommunity(data CommunityCore) model.Community {
	return model.Community{
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
		Members:     data.Members,
		MaxMembers:  data.MaxMembers,
		Image:       data.Image,
	}
}

func ListCoreCommunityToModelCommunity(data []CommunityCore) []model.Community {
	list := []model.Community{}
	for _, v := range data {
		result := CoreCommunityToModelCommunity(v)
		list = append(list, result)
	}
	return list
}

func ModelCommunityToCoreCommunity(data model.Community) CommunityCore {
	return CommunityCore{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
		Members:     data.Members,
		MaxMembers:  data.MaxMembers,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ListModelCommunityToCoreCommunity(data []model.Community) []CommunityCore {
	list := []CommunityCore{}
	for _, v := range data {
		result := ModelCommunityToCoreCommunity(v)
		list = append(list, result)
	}
	return list
}

func EventCoreToEventModel(event CommunityEventCore) model.CommunityEvent {
	eventModel := model.CommunityEvent{
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
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
	return eventModel
}

func EventModelToEventCore(event model.CommunityEvent) CommunityEventCore {
	eventCore := CommunityEventCore{
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
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
	return eventCore
}

func ListEventModelToEventCore(event []model.CommunityEvent) []CommunityEventCore {
	coreEvent := []CommunityEventCore{}
	for _, v := range event {
		event := EventModelToEventCore(v)
		coreEvent = append(coreEvent, event)
	}
	return coreEvent
}
