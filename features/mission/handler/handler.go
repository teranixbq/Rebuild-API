package handler

import (
	"log"
	"net/http"
	"recything/features/mission/dto/request"
	"recything/features/mission/dto/response"
	"recything/features/mission/entity"

	"strings"

	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
)

type missionHandler struct {
	missionService entity.MissionServiceInterface
}

func NewMissionHandler(missionService entity.MissionServiceInterface) *missionHandler {
	return &missionHandler{missionService: missionService}
}

func (mh *missionHandler) CreateMission(e echo.Context) error {
	id, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	requestMission := request.Mission{}
	err = e.Bind(&requestMission)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// log.Println("mission : ", requestMission)
	// log.Println("mission stages", requestMission.MissionStages)

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_EMPTY_FILE))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal upload file"))
	}

	input := request.MissionRequestToMissionCore(requestMission)
	input.AdminID = id
	// err = mh.missionService.CreateMission(input)
	err = mh.missionService.CreateMission(image, input)

	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("Berhasil menambahkan misi"))
}

func (mh *missionHandler) GetAllMission(e echo.Context) error {

	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	search := e.QueryParam("search")
	filter := e.QueryParam("filter")

	result, pagnation, count, err := mh.missionService.FindAllMission(page, limit, search, filter)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_INVALID_TYPE) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("Belum ada misi"))
	}

	response := response.ListMissionCoreToMissionResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCountAll("Berhasil mendapatkan seluruh misi", response, pagnation, count))
}

func (mh *missionHandler) FindById(e echo.Context) error {
	missionID := e.Param("id")
	id, _, err := jwt.ExtractToken(e)
	if id == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	result, err := mh.missionService.FindById(missionID)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	response := response.MissionCoreToMissionResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mengambil data mission", response))
}

func (mh *missionHandler) DeleteMission(e echo.Context) error {
	missionID := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err = mh.missionService.DeleteMission(missionID)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus data misi"))

}

func (mh *missionHandler) UpdateMission(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	id := e.Param("id")
	requestMission := request.Mission{}
	err = e.Bind(&requestMission)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	input := request.MissionRequestToMissionCore(requestMission)
	image, _ := e.FormFile("image")

	err = mh.missionService.UpdateMission(image, id, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_NOT_FOUND))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil mengupdate misi"))
}

func (mh *missionHandler) ClaimMission(e echo.Context) error {
	userID, role, err := jwt.ExtractToken(e)

	if role != "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.Claim{}
	err = helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.ClaimRequestToClaimCore(input)
	log.Println(userID)
	err = mh.missionService.ClaimMission(userID, request)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil melakukan klaim"))

}

// Upload User
func (mh *missionHandler) CreateUploadMission(e echo.Context) error {
	userID, role, err := jwt.ExtractToken(e)
	if role != "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	input := request.UploadMissionTask{}
	if err := e.Bind(&input); err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	form, err := e.MultipartForm()
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan form multipart"))
	}
	images := form.File["image"]

	request := request.UploadMissionTaskRequestToUploadMissionTaskCore(input)

	result, err := mh.missionService.CreateUploadMissionTask(userID, request, images)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	response := response.UploadMissionTaskResponses(result)

	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("berhasil mengupload bukti",response))
}

func (mh *missionHandler) UpdateUploadMission(e echo.Context) error {
	UploadMissionID := e.Param("id")
	input := request.UpdateUploadMissionTask{}

	userID, role, err := jwt.ExtractToken(e)
	if role != "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	if err := e.Bind(&input); err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	form, err := e.MultipartForm()
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan form multipart"))
	}

	images := form.File["images"]

	request := request.UpdateUploadMissionTaskRequestToUpdateUploadMissionTaskCore(input)

	err = mh.missionService.UpdateUploadMissionTask(userID, UploadMissionID, images, request)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil memperbarui bukti"))
}

// Mission Approval
func (mh *missionHandler) GetAllMissionApproval(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	search := e.QueryParam("search")
	filter := e.QueryParam("filter")

	data, pagination, counts, err := mh.missionService.FindAllMissionApproval(page, limit, search, filter)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))

	}

	response := response.ListUpMissionTaskCoreToUpMissionTaskResp(data)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCountAll("Berhasil mendapatkan data", response, pagination, counts))
}

func (mh *missionHandler) GetMissionApprovalById(e echo.Context) error {
	id := e.Param("id")
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	result, err := mh.missionService.FindMissionApprovalById(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	response := response.UpMissionTaskCoreToUpMissionTaskResp(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mengambil data mission", response))
}

func (mh *missionHandler) UpdateStatusApprovalMission(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	id := e.Param("id")
	newStatus := request.StatusApproval{}
	if err := e.Bind(&newStatus); err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	err = mh.missionService.UpdateStatusMissionApproval(id, newStatus.Status, newStatus.Reason)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))

		}

		if helper.HttpResponseCondition(err, constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil mengupdate status approval"))

}


//histories user
func (mh *missionHandler) FindHistoryById(e echo.Context) error {
	trasactionID := e.Param("idTransaksi")
	id, _, err := jwt.ExtractToken(e)
	if id == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	result, err := mh.missionService.FindHistoryById(id, trasactionID)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	resp := response.UpMissionTaskCoreToUpMissionTaskResp(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mengambil data misi", resp))
}

func (mh *missionHandler) GetAllMissionUser(e echo.Context) error {

	id, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	filter := e.QueryParam("filter")
	result, err := mh.missionService.FindAllMissionUser(id, filter)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_NOT_FOUND))
		}

		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("Belum ada misi"))
	}
	resp := response.ListHistoriesCoreToHistoriesResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan seluruh misi", resp))
}