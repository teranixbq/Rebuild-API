package request

import "recything/features/mission/entity"

func MissionRequestToMissionCore(missi Mission) entity.Mission {
	missionCore := entity.Mission{
		Title:            missi.Title,
		MissionImage:     missi.MissionImage,
		Point:            missi.Point,
		Description:      missi.Description,
		StartDate:        missi.Start_Date,
		EndDate:          missi.End_Date,
		TitleStage:       missi.TitleStage,
		DescriptionStage: missi.DescriptionStage,
	}
	return missionCore
}

// func AddMissionStageToMissionStageCore(addMissionStage AddMissionStage) []entity.MissionStage {
//     var missionStages []entity.MissionStage
//     for _, stage := range addMissionStage.Stages {
//         newStage := entity.MissionStage{
//             MissionID:   addMissionStage.MissionID,
//             Title:       stage.Name,
//             Description: stage.DescriptionStage,

//         }
//         missionStages = append(missionStages, newStage)
//     }
//     return missionStages
// }

func RequestMissionStageToMissionStageCore(missionID string, data RequestMissionStage) []entity.MissionStage {
	missionStagesCore := []entity.MissionStage{}
	for _, stage := range data.MissionStage {
		newStage := entity.MissionStage{
			MissionID:   missionID,
			ID:          stage.ID,
			Title:       stage.Title,
			Description: stage.Description,
		}
		missionStagesCore = append(missionStagesCore, newStage)
	}
	return missionStagesCore
}

// func MissionStagesRequestToMissionStagesCore(missionStages MissionStage) entity.MissionStage {
// 	missionStagesCore := entity.MissionStage{
// 		Title:       missionStages.Name,
// 		Description: missionStages.DescriptionStage,
// 	}

// 	return missionStagesCore
// }

// func ListMissionStagesRequestToMissionStagesCore(missionStages []MissionStage) []entity.MissionStage {
// 	missionStagesCore := []entity.MissionStage{}
// 	for _, misiStages := range missionStages {
// 		missi := MissionStagesRequestToMissionStagesCore(misiStages)
// 		missionStagesCore = append(missionStagesCore, missi)
// 	}
// 	return missionStagesCore
// }

func ClaimRequestToClaimCore(claim Claim) entity.ClaimedMission {
	return entity.ClaimedMission{
		MissionID: claim.MissionID,
	}

}

func UploadMissionTaskRequestToUploadMissionTaskCore(data UploadMissionTask) entity.UploadMissionTaskCore {
	return entity.UploadMissionTaskCore{
		UserID:      data.UserID,
		MissionID:   data.MissionID,
		Description: data.Description,
		//Images:   ListImageUploadMissionRequestToImageUploadMissionCore(data.Images),
		// Stage_two:   ListImageUploadMissionRequestToImageUploadMissionCore(data.Stage_two),
		// Stage_three: ListImageUploadMissionRequestToImageUploadMissionCore(data.Stage_three),
	}
}

func UpdateUploadMissionTaskRequestToUpdateUploadMissionTaskCore(data UpdateUploadMissionTask) entity.UploadMissionTaskCore {
	return entity.UploadMissionTaskCore{
		Description: data.Description,
	}
}

// func ImageUploadMissionRequestToImageUploadMissionCore(data ImageUploadMission) entity.ImageUploadMissionCore {
// 	return entity.ImageUploadMissionCore{
// 		Stage_one:   data.Stage_one,
// 		Stage_two:   data.Stage_two,
// 		Stage_three: data.Stage_three,
// 	}
// }

// func ListImageUploadMissionRequestToImageUploadMissionCore(data []ImageUploadMission) []entity.ImageUploadMissionCore {
// 	dataStage := []entity.ImageUploadMissionCore{}
// 	for _, v := range data {
// 		result := ImageUploadMissionRequestToImageUploadMissionCore(v)
// 		dataStage = append(dataStage, result)
// 	}
// 	return dataStage
// }
