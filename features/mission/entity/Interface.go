package entity

import (
	"mime/multipart"
	"recything/features/mission/model"
	"recything/utils/helper"
	"recything/utils/pagination"
)

type MissionRepositoryInterface interface {
	CreateMission(input Mission) error
	FindAllMission(page, limit int, search, filter string) ([]Mission, pagination.PageInfo, helper.CountMission, error)
	FindById(missionID string) (Mission, error)
	GetCountDataMission(filter, search string) (helper.CountMission, error)
	GetCountDataMissionApproval(search string) (helper.CountMissionApproval, error)
	GetImageURL(missionID string) (string, error)
	UpdateMission(missionID string, data Mission) error
	DeleteMission(missionID string) error
	UpdateMissionStatus(data model.Mission) error

	FindAllMissionUser(userID string, filter string) ([]MissionHistories, error)
	FindHistoryById(userID, transactionID string) (UploadMissionTaskCore, error)
	ClaimMission(userID string, data ClaimedMission) error
	FindClaimed(userID, missionID string) error

	FindUploadMissionStatus(id, missionID, userID, status string) error
	FindUploadById(id string) error
	CreateUploadMissionTask(userID string, data UploadMissionTaskCore, images []*multipart.FileHeader) (UploadMissionTaskCore, error)
	UpdateUploadMissionTask(id string, images []*multipart.FileHeader, data UploadMissionTaskCore) error

	FindAllMissionApproval(page, limit int, search, filter string) ([]UploadMissionTaskCore, pagination.PageInfo, helper.CountMissionApproval, error)
	FindMissionApprovalById(UploadMissionTaskID string) (UploadMissionTaskCore, error)
	UpdateStatusMissionApproval(uploadMissionTaskID, status, reason string) error

	FindHistoryByIdTransaction(userID, transactionID string) (map[string]interface{}, error)
}

type MissionServiceInterface interface {
	CreateMission(image *multipart.FileHeader, data Mission) error
	// CreateMission(data Mission) error
	FindAllMission(page, limit, search, filter string) ([]Mission, pagination.PageInfo, helper.CountMission, error)
	FindAllMissionUser(userID string, filter string) ([]MissionHistories, error)
	FindHistoryById(userID, transactionID string) (UploadMissionTaskCore, error)
	FindById(missionID string) (Mission, error)
	UpdateMission(image *multipart.FileHeader, missionID string, data Mission) error

	// UpdateMissionStage(missionStageID string, data []MissionStage) error
	ClaimMission(userID string, data ClaimedMission) error
	DeleteMission(missionID string) error

	CreateUploadMissionTask(userID string, data UploadMissionTaskCore, images []*multipart.FileHeader) (UploadMissionTaskCore, error)
	UpdateUploadMissionTask(userID, id string, images []*multipart.FileHeader, data UploadMissionTaskCore) error

	FindAllMissionApproval(page, limit, search, filter string) ([]UploadMissionTaskCore, pagination.PageInfo, helper.CountMissionApproval, error)
	FindMissionApprovalById(UploadMissionTaskID string) (UploadMissionTaskCore, error)
	UpdateStatusMissionApproval(uploadMissionTaskID, status, reason string) error
}
