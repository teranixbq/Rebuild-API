package handler

import (
	"fmt"
	"net/http"
	"recything/features/user/dto/request"
	"recything/features/user/dto/response"
	"recything/features/user/entity"
	"recything/utils/constanta"
	"recything/utils/email"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUseCase entity.UsersUsecaseInterface
}

func NewUserHandlers(uc entity.UsersUsecaseInterface) *userHandler {
	return &userHandler{
		userUseCase: uc,
	}
}

func (uh *userHandler) Register(e echo.Context) error {
	input := request.UserRegister{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.UsersRequestRegisterToUsersCore(input)

	_, errCreate := uh.userUseCase.Register(request)
	if errCreate != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil membuat data"))
}

func (uh *userHandler) Login(e echo.Context) error {
	// Bind data
	login := request.UserLogin{}
	errBind := helper.DecodeJSON(e, &login)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	dataUser, token, errLogin := uh.userUseCase.Login(login.Email, login.Password)
	if errLogin != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errLogin.Error()))
	}

	response := response.UsersCoreToLoginResponse(dataUser, token)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_LOGIN, response))
}

func (uh *userHandler) GetUserById(e echo.Context) error {
	idUser, _, errExtract := jwt.ExtractToken(e)
	fmt.Println(idUser)
	if errExtract != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(errExtract.Error()))
	}

	result, errGet := uh.userUseCase.GetById(idUser)
	if errGet != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errGet.Error()))
	}

	response := response.UsersCoreToResponseProfile(result)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan profile", response))
}

func (uh *userHandler) UpdateById(e echo.Context) error {
	input := request.UserUpdate{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	idUser, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(errExtract.Error()))
	}

	request := request.UsersRequestUpdateToUsersCore(input)

	errUpdate := uh.userUseCase.UpdateById(idUser, request)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errUpdate.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil melakukan update data"))

}

func (uh *userHandler) UpdatePassword(e echo.Context) error {
	input := request.UserUpdatePassword{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	idUser, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(errExtract.Error()))
	}

	request := request.UsersRequestUpdatePasswordToUsersCore(input)

	err := uh.userUseCase.UpdatePassword(idUser, request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil update password"))
}

func (uh *userHandler) VerifyAccount(e echo.Context) error {
	token := e.QueryParam("token")

	alreadyVerified, err := uh.userUseCase.VerifyUser(token)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if alreadyVerified {
		emailDone, err := email.ParseTemplate("verification_active.html", nil)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ERROR_TEMPLATE))
		}
		return e.HTML(http.StatusOK, emailDone)
	}

	emailContent, err := email.ParseTemplate("success_verification.html", nil)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ERROR_TEMPLATE))
	}
	return e.HTML(http.StatusOK, emailContent)
}

func (uh *userHandler) ForgotPassword(e echo.Context) error {
	input := request.UserSendOTP{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	userCore := request.UsersRequestOTPToUsersCore(input)

	err := uh.userUseCase.SendOTP(userCore.Email)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_DATA_EMAIL) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("otp berhasil dikirim"))
}

func (uh *userHandler) VerifyOTP(e echo.Context) error {
	input := request.UserVerifyOTP{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.UsersRequestVerifyOTPToUsersCore(input)

	token, err := uh.userUseCase.VerifyOTP(request.Email, request.Otp)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal verifikasi "+err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("verifikasi otp berhasil", token))
}

func (uh *userHandler) NewPassword(e echo.Context) error {
	input := request.UserNewPassword{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	email, errExtract := jwt.ExtractTokenVerifikasi(e)
	if errExtract != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(errExtract.Error()))
	}

	request := request.UsersRequestNewPasswordToUsersCore(input)
	err := uh.userUseCase.NewPassword(email, request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil update password"))
}

func (uh *userHandler) JoinCommunity(e echo.Context) error {
	Id, _, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}

	communityId := e.Param("idKomunitas")

	err := uh.userUseCase.JoinCommunity(communityId, Id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}

		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil bergabung dengan komunitas"))
}
