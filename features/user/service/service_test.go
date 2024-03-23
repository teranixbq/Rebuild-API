package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"recything/features/user/entity"
	"recything/mocks"
	"recything/utils/constanta"
	"recything/utils/helper"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getCurrentWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func findEnvFile(dir string) string {
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}

		// Pindah ke direktori induk
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}

		dir = parentDir
	}

	return ""
}

func ReadEmailTemplate(templateName string) (string, error) {
	filePath := fmt.Sprintf("utils/email/templates/%s", templateName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// TestRegister tests the Register method of userService.
func TestRegister(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	mockEmail := new(mocks.MockUserRepository)
	userSvc := NewUserService(mockRepo)

	// readEmailTemplate := func(templateName string) (string, error) {
	// 	content, err := os.ReadFile(templateName)
	// 	if err != nil {
	// 		fmt.Printf("Error reading template: %v\n", err)
	// 		return "", fmt.Errorf("failed to read email template: %v", err)
	// 	}
	// 	return string(content), nil
	// }

	// t.Run("Registrasi Valid", func(t *testing.T) {
	// 	_, err := readEmailTemplate("../../../utils/email/templates/account_registration.html")
	// 	assert.Nil(t, err)

	// 	wd, err := os.Getwd()
	// 	if err != nil {
	// 		t.Fatalf("Failed to get working directory: %v", err)
	// 	}
	// 	t.Logf("Working Directory: %s", wd)

	// 	mockRepo.On("FindByEmail", "test@example.com").Return(entity.UsersCore{}, errors.New(constanta.ERROR_EMAIL_EXIST))
	// 	mockRepo.On("Register", mock.Anything).Return(entity.UsersCore{VerificationToken: "some_token"}, nil)
	// 	mockEmail.On("SendVerificationEmail", "test@example.com", "some_token").Return(nil).Once()

	// 	userData := entity.UsersCore{
	// 		Fullname:        "John Doe",
	// 		Email:           "test@example.com",
	// 		Password:        "password123",
	// 		ConfirmPassword: "password123",
	// 	}

	// 	result, err := userSvc.Register(userData)
	// 	assert.Nil(t, err)
	// 	assert.NotEmpty(t, result.VerificationToken)
	// 	mockRepo.AssertExpectations(t)
	// })

	t.Run("Data Empty", func(t *testing.T) {

		requestBody := entity.UsersCore{
			Fullname:        "",
			Email:           "",
			Password:        "",
			ConfirmPassword: "",
		}

		_, err := userSvc.Register(requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Invalid Fomat", func(t *testing.T) {

		requestBody := entity.UsersCore{
			Fullname:        "John Doe",
			Email:           "johnexecom",
			Password:        "password123",
			ConfirmPassword: "password123",
		}

		_, err := userSvc.Register(requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Invalid Length", func(t *testing.T) {

		requestBody := entity.UsersCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "pass",
			ConfirmPassword: "pass",
		}

		_, err := userSvc.Register(requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Confirm Password", func(t *testing.T) {
		mockRepo.On("FindByEmail", "john@example.com").Return(entity.UsersCore{}, errors.New("user not found"))

		requestBody := entity.UsersCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "ayamgoreng123",
			ConfirmPassword: "ayamgoreg12345",
		}
		result, err := userSvc.Register(requestBody)
		assert.Error(t, err)

		assert.Equal(t, entity.UsersCore{}, result)
		assert.Equal(t, constanta.ERROR_CONFIRM_PASSWORD, err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Duplicat", func(t *testing.T) {
		mockRepo.On("FindByEmail", "duplicate@example.com").Return(entity.UsersCore{}, nil)

		duplicateUserData := entity.UsersCore{
			Fullname:        "Jane Doe",
			Email:           "duplicate@example.com",
			Password:        "password",
			ConfirmPassword: "password",
		}

		mockEmail.On("SendVerificationEmail", mock.Anything, mock.Anything).Return(nil).Once()

		_, err := userSvc.Register(duplicateUserData)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constanta.ERROR_EMAIL_EXIST)
		mockRepo.AssertExpectations(t)
	})
}

// TestLogin tests the Login method of userService.
func TestLogin(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	mockUser := new(mocks.MockUserRepository)
	userSvc := NewUserService(mockRepo)

	currentDir := getCurrentWorkingDirectory()

	envFile := findEnvFile(currentDir)
	godotenv.Load(envFile)

	t.Run("Login Success", func(t *testing.T) {
		mockUserData := entity.UsersCore{
			Id:                "user_id",
			Fullname:          "John Doe",
			Email:             "test@example.com",
			Password:          "budigagah123",
			IsVerified:        true,
			VerificationToken: "",
		}

		password, err := helper.HashPassword("correct_password")
		assert.Nil(t, err)
		mockUserData.Password = password

		mockRepo.On("FindByEmail", "test@example.com").Return(mockUserData, nil)
		mockUser.On("CreateToken", mock.Anything, mock.Anything).Return("some_token", nil)

		email := "test@example.com"
		passwordInput := "correct_password"

		isPasswordCorrect := helper.CompareHash(mockUserData.Password, passwordInput)
		assert.True(t, isPasswordCorrect)

		dataUser, token, err := userSvc.Login(email, passwordInput)

		mockRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, "user_id", dataUser.Id)

		actualPayload, _ := helper.DecodeJWTToken(token)

		assert.Equal(t, "", actualPayload)
	})

	t.Run("Email Not Found", func(t *testing.T) {
		mockRepo.On("FindByEmail", "nonexistent@example.com").Return(entity.UsersCore{}, errors.New(constanta.EMAIL_NOT_REGISTER))

		email := "nonexistent@example.com"
		password := "password"

		_, _, err := userSvc.Login(email, password)
		mockRepo.AssertExpectations(t)

		// Periksa hasil fungsi Login
		assert.NotNil(t, err)
		assert.EqualError(t, err, constanta.EMAIL_NOT_REGISTER)
	})

	t.Run("Data Empty", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Email:    "",
			Password: "",
		}
		_, _, err := userSvc.Login(requestBody.Email, requestBody.Password)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Account Has Not Been Verified", func(t *testing.T) {
		mockUserData := entity.UsersCore{
			Id:                "user_id",
			Fullname:          "John Doe",
			Email:             "unverified@example.com",
			Password:          "correct_password_hash",
			IsVerified:        false,
			VerificationToken: "some_verification_token",
		}

		mockRepo.On("FindByEmail", "unverified@example.com").Return(mockUserData, nil)
		email := "unverified@example.com"
		password := "password"

		_, _, err := userSvc.Login(email, password)

		mockRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "akun belum terverifikasi")
	})

	t.Run("Wrong Password ", func(t *testing.T) {
		mockUserData := entity.UsersCore{
			Id:                "user_id",
			Fullname:          "John Doe",
			Email:             "test@example.com",
			Password:          "correct_password_hash",
			IsVerified:        true,
			VerificationToken: "",
		}

		mockRepo.On("FindByEmail", "test@example.com").Return(mockUserData, nil)
		email := "test@example.com"
		password := "wrong_password"

		_, _, err := userSvc.Login(email, password)

		mockRepo.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "email atau password salah")
	})

	t.Run("Email Invalid Fomat", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Email:    "johnexecom",
			Password: "password123",
		}

		_, _, err := userSvc.Login(requestBody.Email, requestBody.Password)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// TestGetById tests the GetById method of userService.
func TestGetById(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	t.Run("Valid Get ById", func(t *testing.T) {
		userID := "user123"

		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, nil)
		user, err := userSvc.GetById(userID)
		mockRepo.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, user)
	})

	t.Run("Invalid Get ById", func(t *testing.T) {
		userID := ""
		_, err := userSvc.GetById(userID)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constanta.ERROR_ID_INVALID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Valid Update Badge", func(t *testing.T) {
		userID := "user123"
		mockRepo.On("UpdateBadge", userID).Return(nil).Once()

		err := mockRepo.UpdateBadge(userID)

		mockRepo.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Repository", func(t *testing.T) {
		userID := "12345asd"
		mockRepo := new(mocks.UsersRepositoryInterface)
		userSvc := NewUserService(mockRepo)

		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, errors.New("data user tidak ada")).Once()

		_, err := userSvc.GetById(userID)
		assert.Error(t, err)
		assert.EqualError(t, err, "data user tidak ada")
		mockRepo.AssertExpectations(t)
	})
}

// TestUpdateById tests the UpdateById method of userService.
func TestUpdateById(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	t.Run("Valid Update By Id", func(t *testing.T) {
		userID := "user123"
		userData := entity.UsersCore{
			Fullname:    "Updated Name",
			Phone:       "082189638011",
			Address:     "Updated Address",
			DateOfBirth: "1990-01-01",
			Purpose:     "Updated Purpose",
		}

		// Mock pemanggilan GetById dengan mengembalikan data yang sesuai
		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, nil)
		mockRepo.On("UpdateById", userID, userData).Return(nil)

		err := userSvc.UpdateById(userID, userData)
		assert.Nil(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Update ById", func(t *testing.T) {
		userID := ""
		userData := entity.UsersCore{}
		err := userSvc.UpdateById(userID, userData)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constanta.ERROR_ID_INVALID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Get By Id", func(t *testing.T) {
		userID := "user123"
		mockRepo := new(mocks.UsersRepositoryInterface)
		userSvc := NewUserService(mockRepo)

		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, errors.New("failed to get user by ID")).Once()
		err := userSvc.UpdateById(userID, entity.UsersCore{})

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to get user by ID")
		mockRepo.AssertExpectations(t)
		mockRepo.AssertCalled(t, "GetById", userID)
		mockRepo.AssertNotCalled(t, "UpdateBadge", userID)
	})

	t.Run("Data Empty", func(t *testing.T) {
		userID := "user123"
		requestBody := entity.UsersCore{
			Fullname:    "",
			Phone:       "",
			Address:     "",
			DateOfBirth: "",
			Purpose:     "",
		}
		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, nil)

		err := userSvc.UpdateById(userID, requestBody)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Phone Number", func(t *testing.T) {
		userID := "user123"
		requestBody := entity.UsersCore{
			Fullname:    "budiawan",
			Phone:       "1234567",
			Address:     "minasaupa",
			DateOfBirth: "2023-04-01",
			Purpose:     "pake nanya",
		}
		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, nil)

		err := userSvc.UpdateById(userID, requestBody)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Date", func(t *testing.T) {
		userID := "user123"
		requestBody := entity.UsersCore{
			Fullname:    "budiawan",
			Phone:       "082189638011",
			Address:     "minasaupa",
			DateOfBirth: "01-04-2023",
			Purpose:     "pake nanya",
		}
		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, nil)

		err := userSvc.UpdateById(userID, requestBody)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Fullname:    "budiawan",
			Phone:       "082189638011",
			Address:     "minasaupa",
			DateOfBirth: "01-04-2023",
			Purpose:     "pake nanya",
		}

		userID := "2"
		mockRepo.On("GetById", mock.AnythingOfType("string")).Return(requestBody, errors.New("data user tidak ada"))
		user, err := userSvc.GetById(userID)

		assert.Error(t, err)
		assert.NotEqual(t, userID, requestBody.Id)
		assert.EqualError(t, err, "data user tidak ada")
		assert.Empty(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateUserFailure(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	id := "some_user_id"
	requestBody := entity.UsersCore{
		Fullname:    "budiawan",
		Phone:       "082189638011",
		Address:     "minasaupa",
		DateOfBirth: "2023-01-04",
		Purpose:     "pake nanya",
	}
	mockRepo.On("GetById", id).Return(entity.UsersCore{}, nil)
	mockRepo.On("UpdateById", id, requestBody).Return(errors.New("update failed"))

	err := userSvc.UpdateById(id, requestBody)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "update failed")
	mockRepo.AssertExpectations(t)
}

// TestUpdatePassword tests the UpdatePassword method of userService.
func TestUpdatePassword(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	t.Run("Valid Update Password", func(t *testing.T) {
		userID := "user123"
		passwordData := entity.UsersCore{
			Password:        "old_password",
			NewPassword:     "new_password",
			ConfirmPassword: "new_password",
		}

		oldPasswordHash, _ := helper.HashPassword(passwordData.Password)

		mockRepo.On("GetById", userID).Return(entity.UsersCore{Password: oldPasswordHash}, nil)
		mockRepo.On("UpdatePassword", userID, mock.Anything).Return(nil)

		err := userSvc.UpdatePassword(userID, passwordData)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Get By Id", func(t *testing.T) {
		userID := "user123"
		mockRepo := new(mocks.UsersRepositoryInterface)
		userSvc := NewUserService(mockRepo)

		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, errors.New("data user tidak ada")).Once()
		err := userSvc.UpdatePassword(userID, entity.UsersCore{})

		assert.Error(t, err)
		assert.EqualError(t, err, "data user tidak ada")
		mockRepo.AssertExpectations(t)
		mockRepo.AssertCalled(t, "GetById", userID)
		mockRepo.AssertNotCalled(t, "UpdateBadge", userID)
	})

	t.Run("Data Empty", func(t *testing.T) {
		userID := "user123"
		requestBody := entity.UsersCore{
			Fullname:    "",
			Phone:       "",
			Address:     "",
			DateOfBirth: "",
			Purpose:     "",
		}
		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, nil)

		err := userSvc.UpdatePassword(userID, requestBody)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Invalid Length", func(t *testing.T) {
		userID := "user123"
		requestBody := entity.UsersCore{
			Password:        "123456asd",
			NewPassword:     "pass",
			ConfirmPassword: "pass",
		}

		mockRepo.On("GetById", userID).Return(entity.UsersCore{}, nil)
		err := userSvc.UpdatePassword(userID, requestBody)

		assert.Error(t, err)
		assert.EqualError(t, err, "minimal 8 karakter, ulangi kembali!")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Mismatch", func(t *testing.T) {
		userID := "user123"
		requestBody := entity.UsersCore{
			Password:        "currentPassword",
			NewPassword:     "wrongNewPassword",
			ConfirmPassword: "wrongNewPasswordConfirmation",
		}

		mockRepo.On("GetById", userID).Return(entity.UsersCore{Password: "currentPassword"}, nil)
		err := userSvc.UpdatePassword(userID, requestBody)
		assert.Error(t, err)
		assert.EqualError(t, err, constanta.ERROR_PASSWORD)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Confirm Password", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepositoryInterface)
		userSvc := NewUserService(mockRepo)

		userID := "user123"
		requestBody := entity.UsersCore{
			Password:        "currentPassword",
			NewPassword:     "newPassword",
			ConfirmPassword: "mismatchedPassword",
		}
		hashedPassword, _ := helper.HashPassword("currentPassword")
		mockRepo.On("GetById", userID).Return(entity.UsersCore{
			Password: hashedPassword,
		}, nil)

		err := userSvc.UpdatePassword(userID, requestBody)
		assert.Error(t, err)
		assert.EqualError(t, err, constanta.ERROR_CONFIRM_PASSWORD)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Invalid Update Password", func(t *testing.T) {
		userID := ""
		passwordData := entity.UsersCore{}
		err := userSvc.UpdatePassword(userID, passwordData)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constanta.ERROR_ID_INVALID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Update Password", func(t *testing.T) {
		userID := ""
		passwordData := entity.UsersCore{}
		err := userSvc.UpdatePassword(userID, passwordData)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constanta.ERROR_ID_INVALID)
		mockRepo.AssertExpectations(t)
	})
}

// TestVerifyUser tests the VerifyUser method of userService.
func TestVerifyUser(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	t.Run("Valid Verify User", func(t *testing.T) {
		token := "valid_token"
		mockRepo.On("GetByVerificationToken", token).Return(entity.UsersCore{IsVerified: false}, nil)
		mockRepo.On("UpdateIsVerified", mock.Anything, true).Return(nil)

		isVerified, err := userSvc.VerifyUser(token)

		assert.False(t, isVerified)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Verify User", func(t *testing.T) {
		token := "invalid_token"
		mockRepo.On("GetByVerificationToken", token).Return(entity.UsersCore{IsVerified: true}, nil)

		isVerified, err := userSvc.VerifyUser(token)
		assert.True(t, isVerified)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		isVerified, err := userSvc.VerifyUser("")

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid token")
		assert.False(t, isVerified)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		token := "test_token"

		mockRepo.On("GetByVerificationToken", token).Return(entity.UsersCore{}, errors.New("database error")).Once()
		isVerified, err := userSvc.VerifyUser(token)

		assert.Error(t, err)
		assert.EqualError(t, err, "gagal mendapatkan data")
		assert.False(t, isVerified)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Verified", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepositoryInterface)
		userSvc := NewUserService(mockRepo)

		token := "test_token"
		mockRepo.On("GetByVerificationToken", token).Return(entity.UsersCore{Id: "user123"}, nil).Once()
		mockRepo.On("UpdateIsVerified", "user123", true).Return(errors.New("database error")).Once()

		isVerified, err := userSvc.VerifyUser(token)
		assert.Error(t, err)
		assert.EqualError(t, err, "gagal untuk mengaktifkan user")
		assert.False(t, isVerified)

		mockRepo.AssertExpectations(t)
	})

}

// TestUpdateIsVerified tests the UpdateIsVerified method of userService.
func TestUpdateIsVerified(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	t.Run("Valid Update Verified", func(t *testing.T) {
		userID := "user123"
		isVerified := true
		mockRepo.On("UpdateIsVerified", userID, isVerified).Return(nil)

		err := userSvc.UpdateIsVerified(userID, isVerified)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Update Verified", func(t *testing.T) {
		userID := ""
		isVerified := false
		err := userSvc.UpdateIsVerified(userID, isVerified)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constanta.ERROR_ID_INVALID)
		mockRepo.AssertExpectations(t)
	})
}

// TestSendOTP tests the SendOTP method of userService.
func TestSendOTP(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	// t.Run("Valid Send OTP", func(t *testing.T) {
	// 	email := "test@example.com"
	// 	mockRepo.On("SendOTP", email, mock.Anything, mock.Anything).Return(entity.UsersCore{}, nil)
	// 	err := userSvc.SendOTP(email)
	// 	assert.Nil(t, err)
	// 	mockRepo.AssertExpectations(t)
	// })

	t.Run("Data Empty", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Email: "",
		}
		err := userSvc.SendOTP(requestBody.Email)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Format", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Email: "ayamgoreng",
		}
		err := userSvc.SendOTP(requestBody.Email)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Generate OTP", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)

		mockRepo.On("GenerateOTP", 4).Return("", errors.New("generate otp gagal")).Once()
		result, err := mockRepo.GenerateOTP(4)

		assert.Error(t, err)
		assert.EqualError(t, err, "generate otp gagal")
		assert.Empty(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Send OTP", func(t *testing.T) {
		email := "invalid@example.com"
		// Mock user not found scenario
		mockRepo.On("SendOTP", email, mock.Anything, mock.Anything).Return(entity.UsersCore{}, errors.New("user not found"))
		err := userSvc.SendOTP(email)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "user not found")
		mockRepo.AssertExpectations(t)
	})

}

// TestVerifyOTP tests the VerifyOTP method of userService.
func TestVerifyOTP(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	t.Run("Valid Verify OTP", func(t *testing.T) {
		email := "test@example.com"
		// Mock user data for verification
		userData := entity.UsersCore{Email: email, OtpExpiration: time.Now().Add(time.Minute).Unix()}
		mockRepo.On("VerifyOTP", email, "123456").Return(userData, nil)

		token, err := userSvc.VerifyOTP(email, "123456")
		assert.NotNil(t, err)
		assert.EqualError(t, err, "otp tidak valid")
		assert.Empty(t, token)

		// Ensure that ResetOTP is called once
		mockRepo.AssertExpectations(t)
	})

	t.Run("Data Empty", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Email: "",
			Otp:   "",
		}
		_, err := userSvc.VerifyOTP(requestBody.Email, requestBody.Otp)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email or OTP Invalid", func(t *testing.T) {
		mockRepo.On("VerifyOTP", mock.Anything, mock.Anything).Return(entity.UsersCore{}, errors.New("email atau otp salah")).Once()
		requestBody := entity.UsersCore{
			Email: "asdasd",
			Otp:   "aVC5",
		}

		_, err := userSvc.VerifyOTP(requestBody.Email, requestBody.Otp)
		assert.Error(t, err)
		assert.EqualError(t, err, "email atau otp salah")

		mockRepo.AssertExpectations(t)
	})

	t.Run("Otp Expired", func(t *testing.T) {
		expiredOTP := "123456"
		expirationTime := time.Now().Add(-time.Hour)
		mockRepo.On("VerifyOTP", mock.Anything, mock.Anything).Return(entity.UsersCore{Otp: expiredOTP, OtpExpiration: expirationTime.Unix()}, nil).Once()

		userSvc := NewUserService(mockRepo)
		_, err := userSvc.VerifyOTP("anyemail@example.com", "anyOTP")
		assert.Error(t, err)
		assert.EqualError(t, err, "otp sudah kadaluwarsa")
		mockRepo.AssertExpectations(t)
	})

}

func TestUpdateNewPassword(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	t.Run("Valid New Password", func(t *testing.T) {
		mockRepo.On("NewPassword", mock.Anything, mock.Anything).Return(entity.UsersCore{}, nil).Once()
		email := "test@example.com"
		password := "NewSecurePass"
		confirmPassword := "NewSecurePass"

		err := userSvc.NewPassword(email, entity.UsersCore{
			Password:        password,
			ConfirmPassword: confirmPassword,
		})

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Mismatch", func(t *testing.T) {
		us := &userService{userRepo: mockRepo}

		err := us.NewPassword("mismatch@email.com", entity.UsersCore{
			Password:        "password1",
			ConfirmPassword: "password2",
		})

		assert.EqualError(t, err, constanta.ERROR_PASSWORD)
	})

	t.Run("Data Empty", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Email:           "",
			Password:        "",
			ConfirmPassword: "",
		}

		err := userSvc.NewPassword(requestBody.Email, requestBody)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Invalid Length", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Email:           "bidadari@gmail.com",
			Password:        "pass",
			ConfirmPassword: "pass",
		}

		err := userSvc.NewPassword(requestBody.Email, requestBody)

		assert.Error(t, err)
		assert.EqualError(t, err, "minimal 8 karakter, ulangi kembali!")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Format", func(t *testing.T) {
		requestBody := entity.UsersCore{
			Email:           "ayamgoreng",
			Password:        "123456asd",
			ConfirmPassword: "123456asd",
		}
		err := userSvc.NewPassword(requestBody.Email, requestBody)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// TestJoinCommunity tests the JoinCommunity method of userService.
func TestJoinCommunity(t *testing.T) {
	mockRepo := new(mocks.UsersRepositoryInterface)
	userSvc := NewUserService(mockRepo)

	t.Run("Valid Join Community", func(t *testing.T) {
		userID := "user123"
		communityID := "community456"
		mockRepo.On("JoinCommunity", userID, communityID).Return(nil)

		err := userSvc.JoinCommunity(userID, communityID)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Join Community", func(t *testing.T) {
		userID := ""
		communityID := "community456"
		err := userSvc.JoinCommunity(userID, communityID)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "id tidak boleh kosong")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Join Community", func(t *testing.T) {
		mockRepo.On("JoinCommunity", mock.Anything, mock.Anything).Return(errors.New("failed to join community")).Once()

		communityID := "community123"
		userID := "user123"

		err := userSvc.JoinCommunity(communityID, userID)

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to join community")
		mockRepo.AssertExpectations(t)
	})
}
