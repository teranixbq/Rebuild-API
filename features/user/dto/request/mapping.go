package request

import "recything/features/user/entity"

func UsersRequestRegisterToUsersCore(data UserRegister) entity.UsersCore {
	return entity.UsersCore{
		Fullname:        data.Fullname,
		Email:           data.Email,
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
	}
}

func UsersRequestUpdateToUsersCore(data UserUpdate) entity.UsersCore {
	return entity.UsersCore{
		Fullname:    data.Fullname,
		Phone:       data.Phone,
		DateOfBirth: data.DateOfBirth,
		Address:     data.Address,
		Purpose:     data.Purpose,
	}
}

func UsersRequestUpdatePasswordToUsersCore(data UserUpdatePassword) entity.UsersCore {
	return entity.UsersCore{
		Password:        data.Password,
		NewPassword:     data.NewPassword,
		ConfirmPassword: data.ConfirmPassword,
	}
}

func UsersRequestNewPasswordToUsersCore(data UserNewPassword) entity.UsersCore {
	return entity.UsersCore{
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
	}
}

func UsersRequestLoginToUsersCore(data UserLogin) entity.UsersCore {
	return entity.UsersCore{
		Fullname: data.Email,
		Password: data.Password,
	}
}

func UsersRequestOTPToUsersCore(data UserSendOTP) entity.UsersCore {
	return entity.UsersCore{
		Email: data.Email,
	}
}

func UsersRequestVerifyOTPToUsersCore(data UserVerifyOTP) entity.UsersCore {
	return entity.UsersCore{
		Email: data.Email,
		Otp:   data.Otp,
	}
}
