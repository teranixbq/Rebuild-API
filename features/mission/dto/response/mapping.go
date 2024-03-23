package response

import "recything/features/mission/entity"

func MissionCoreToMissionResponse(missi entity.Mission) Mission {
	missionResp := Mission{
		ID:               missi.ID,
		Title:            missi.Title,
		Status:           missi.Status,
		Creator:          missi.Creator,
		MissionImage:     missi.MissionImage,
		Point:            missi.Point,
		Description:      missi.Description,
		StartDate:        missi.StartDate,
		EndDate:          missi.EndDate,
		DescriptionStage: missi.DescriptionStage,
		TitleStage:       missi.TitleStage,
		// MissionStages: []MissionStage{},
		CreatedAt: missi.CreatedAt,
		UpdatedAt: missi.UpdatedAt,
	}
	// missionStagesResp := ListMissionStagesCoreToMissionStagesResponse(missi.MissionStages)
	// missionResp.MissionStages = missionStagesResp
	return missionResp
}

func MissionStagesCoreToMissionStagesResponse(missionStages entity.MissionStage) MissionStage {
	missionStagesResp := MissionStage{
		ID:          missionStages.ID,
		Title:       missionStages.Title,
		Description: missionStages.Description,
		CreatedAt:   missionStages.CreatedAt,
		UpdatedAt:   missionStages.UpdatedAt,
	}
	return missionStagesResp
}

func ListMissionStagesCoreToMissionStagesResponse(missionStages []entity.MissionStage) []MissionStage {
	missionStagesResp := []MissionStage{}
	for _, misiStages := range missionStages {
		missi := MissionStagesCoreToMissionStagesResponse(misiStages)
		missionStagesResp = append(missionStagesResp, missi)
	}
	return missionStagesResp
}

func ListMissionCoreToMissionResponse(mission []entity.Mission) []Mission {
	missions := []Mission{}
	for _, mission := range mission {
		missionResp := MissionCoreToMissionResponse(mission)
		missions = append(missions, missionResp)
	}
	return missions
}

func ImgUpMissionCoreToImgUpMissionResponse(data entity.ImageUploadMissionCore) ImageUploadMission {
	return ImageUploadMission{
		ID:                  data.ID,
		UploadMissionTaskID: data.UploadMissionTaskID,
		Image:               data.Image,
		CreatedAt:           data.CreatedAt,
	}
}

func ListImgUpMissionCoreToImgUpMissionResponse(data []entity.ImageUploadMissionCore) []ImageUploadMission {
	list := []ImageUploadMission{}
	for _, v := range data {
		result := ImgUpMissionCoreToImgUpMissionResponse(v)
		list = append(list, result)
	}
	return list
}

func UpMissionTaskCoreToUpMissionTaskResp(data entity.UploadMissionTaskCore) UploadMissionTask {
	return UploadMissionTask{
		ID:          data.ID,
		UserID:      data.UserID,
		User:        data.User,
		MissionID:   data.MissionID,
		MissionName: data.MissionName,
		Description: data.Description,
		Reason:      data.Reason,
		Images:      ListImgUpMissionCoreToImgUpMissionResponse(data.Images),
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ListUpMissionTaskCoreToUpMissionTaskResp(data []entity.UploadMissionTaskCore) []UploadMissionTask {
	list := []UploadMissionTask{}
	for _, v := range data {
		result := UpMissionTaskCoreToUpMissionTaskResp(v)
		list = append(list, result)
	}
	return list
}

func HistoriesCoreToHistoriesResponse(data entity.MissionHistories) MissionHistories {
	return MissionHistories{
		MissionID:      data.MissionID,
		ClaimedID:      data.ClaimedID,
		TransactionID:  data.TransactionID,
		Title:          data.Title,
		StatusApproval: data.StatusApproval,
		StatusMission:  data.StatusMission,
		MissionImage:   data.MissionImage,
		// Reason:           data.Reason,
		Point:            data.Point,
		Description:      data.Description,
		StartDate:        data.StartDate,
		EndDate:          data.EndDate,
		DescriptionStage: data.DescriptionStage,
		TitleStage:       data.TitleStage,
		// MissionStages:  []MissionStage{},
		CreatedAt: data.CreatedAt,
	}
}

func ListHistoriesCoreToHistoriesResponse(data []entity.MissionHistories) []MissionHistories {
	list := []MissionHistories{}
	for _, v := range data {
		result := HistoriesCoreToHistoriesResponse(v)
		list = append(list, result)
	}
	return list
}

func UploadMissionTaskResponses(data entity.UploadMissionTaskCore) UploadMissionTaskResponse {
	return UploadMissionTaskResponse{
		ID:        data.ID,
	}
}
