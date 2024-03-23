package request

import "recything/features/community/entity"

func RequestCommunityToCoreCommunity(data CommunityRequest) entity.CommunityCore {
	return entity.CommunityCore{
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
		MaxMembers:  data.Max_Members,
		Image:       data.Image,
	}
}

func EventRequestToEventCore(event EventRequest) entity.CommunityEventCore {
	return entity.CommunityEventCore{
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
