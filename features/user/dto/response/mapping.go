package response

import "recything/features/user/entity"

func UsersCoreToUsersCreateResponse(data entity.UsersCore) UserCreateResponse {
	return UserCreateResponse{
		Id:       data.Id,
		Fullname: data.Fullname,
		Email:    data.Email,
	}
}

func UsersCoreToLoginResponse(data entity.UsersCore, token string) UserLoginResponse {
	return UserLoginResponse{
		Id:       data.Id,
		Fullname: data.Fullname,
		Email:    data.Email,
		Token:    token,
	}
}

func UsersCoreToResponseProfile(data entity.UsersCore) UserResponseProfile {
	userResp := UserResponseProfile{
		Id:           data.Id,
		Fullname:     data.Fullname,
		Email:        data.Email,
		DateOfBirth:  data.DateOfBirth,
		Phone:        data.Phone,
		Address:      data.Address,
		Purpose:      data.Purpose,
		Point:        data.Point,
		Badge:        data.Badge,
		Community_id: data.Community_id,
	}
	community := ListCommunityCoreToCommunityResponse(data.Communities)
	userResp.Communities = community
	return userResp
}

func UsersCoreToResponseManageUsers(data entity.UsersCore) UserResponseManageUsers {
	return UserResponseManageUsers{
		Id:       data.Id,
		Fullname: data.Fullname,
		Email:    data.Email,
		Point:    data.Point,
	}
}

func UsersCoreToResponseManageUsersList(dataCore []entity.UsersCore) []UserResponseManageUsers {
	var result []UserResponseManageUsers
	for _, v := range dataCore {
		result = append(result, UsersCoreToResponseManageUsers(v))
	}
	return result
}

func UsersCoreToResponseDetailManageUsers(data entity.UsersCore) UserResponseDetailManageUsers {
	return UserResponseDetailManageUsers{
		Id:          data.Id,
		Fullname:    data.Fullname,
		Email:       data.Email,
		Point:       data.Point,
		Address:     data.Address,
		DateOfBirth: data.DateOfBirth,
		Purpose:     data.Purpose,
		CreatedAt:   data.CreatedAt,
	}
}

func CommunityCoreToCommunityResponse(community entity.UserCommunityCore) UserCommunityResponse {
	return UserCommunityResponse{
		Id:       community.Id,
		Name:     community.Name,
		Image:    community.Image,
		Location: community.Location,
	}
}

func ListCommunityCoreToCommunityResponse(communities []entity.UserCommunityCore) []UserCommunityResponse {
	ResponseCommunity := []UserCommunityResponse{}
	for _, v := range communities {
		community := CommunityCoreToCommunityResponse(v)
		ResponseCommunity = append(ResponseCommunity, community)
	}
	return ResponseCommunity
}

func UserDailyPointsCoreToUserDailyPointsResponse(data entity.UserDailyPointsCore) UserDailyPointsResponse {
	return UserDailyPointsResponse{
		Claim:        data.Claim,
		DailyPointID: data.DailyPointID,
		CreatedAt:    data.CreatedAt,
	}
}

func ListUserDailyPointsCoreToUserDailyPointsResponse(data []entity.UserDailyPointsCore) []UserDailyPointsResponse {
	dataDaily := []UserDailyPointsResponse{}
	for _, v := range data {
		daily := UserDailyPointsCoreToUserDailyPointsResponse(v)
		dataDaily = append(dataDaily, daily)
	}
	return dataDaily
}
