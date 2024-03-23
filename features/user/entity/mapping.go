package entity

import (
	como "recything/features/community/model"
	"recything/features/user/model"
)

func UsersCoreToUsersModel(mainData UsersCore) model.Users {
	userModel := model.Users{
		Email:             mainData.Email,
		Password:          mainData.Password,
		Fullname:          mainData.Fullname,
		Phone:             mainData.Phone,
		Address:           mainData.Address,
		DateOfBirth:       mainData.DateOfBirth,
		Purpose:           mainData.Purpose,
		Point:             mainData.Point,
		IsVerified:        mainData.IsVerified,
		VerificationToken: mainData.VerificationToken,
		Otp:               mainData.Otp,
		OtpExpiration:     mainData.OtpExpiration,
	}
	community := ListCommunityCoreToCommunityModel(mainData.Communities)
	userModel.Communities = community
	return userModel
}

func ListUserCoreToUserModel(mainData []UsersCore) []model.Users {
	listUser := []model.Users{}
	for _, user := range mainData {
		userModel := UsersCoreToUsersModel(user)
		listUser = append(listUser, userModel)
	}
	return listUser
}

func UsersModelToUsersCore(mainData model.Users) UsersCore {
	userCore := UsersCore{
		Id:                mainData.Id,
		Email:             mainData.Email,
		Password:          mainData.Password,
		Fullname:          mainData.Fullname,
		Badge:             mainData.Badge,
		Phone:             mainData.Phone,
		Address:           mainData.Address,
		DateOfBirth:       mainData.DateOfBirth,
		Purpose:           mainData.Purpose,
		Point:             mainData.Point,
		IsVerified:        mainData.IsVerified,
		VerificationToken: mainData.VerificationToken,
		CreatedAt:         mainData.CreatedAt,
		UpdatedAt:         mainData.UpdatedAt,
		Otp:               mainData.Otp,
		OtpExpiration:     mainData.OtpExpiration,
	}
	community := ListCommunityModelToCommunityCore(mainData.Communities)
	userCore.Communities = community
	return userCore
}

func ListUserModelToUserCore(mainData []model.Users) []UsersCore {
	listUser := []UsersCore{}
	for _, user := range mainData {
		userModel := UsersModelToUsersCore(user)
		listUser = append(listUser, userModel)
	}
	return listUser
}

func CommunityModelToCommunityCore(community como.Community) UserCommunityCore {
	return UserCommunityCore{
		Id:       community.Id,
		Name:     community.Name,
		Image:    community.Image,
		Location: community.Location,
	}
}

func ListCommunityModelToCommunityCore(community []como.Community) []UserCommunityCore {
	communityCore := []UserCommunityCore{}
	for _, v := range community {
		community := CommunityModelToCommunityCore(v)
		communityCore = append(communityCore, community)
	}
	return communityCore
}

func CommunityCoreToCommunityModel(community UserCommunityCore) como.Community {
	return como.Community{
		Id:       community.Id,
		Name:     community.Name,
		Image:    community.Image,
		Location: community.Location,
	}
}

func ListCommunityCoreToCommunityModel(community []UserCommunityCore) []como.Community {
	communityCore := []como.Community{}
	for _, v := range community {
		communitys := CommunityCoreToCommunityModel(v)
		communityCore = append(communityCore, communitys)
	}
	return communityCore
}

func UserDailyPointsModelToUserDailyPointsCore(data model.UserDailyPoints) UserDailyPointsCore {
	return UserDailyPointsCore{

		UsersID:      data.UsersID,
		DailyPointID: data.DailyPointID,
		Claim:        data.Claim,
		CreatedAt:    data.CreatedAt,
	}

}

func ListUserDailyPointsModelToUserDailyPointsCore(data []model.UserDailyPoints) []UserDailyPointsCore {
	dataDaily := []UserDailyPointsCore{}
	for _, v := range data {
		daily := UserDailyPointsModelToUserDailyPointsCore(v)
		dataDaily = append(dataDaily, daily)
	}
	return dataDaily
}
