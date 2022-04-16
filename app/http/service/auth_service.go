package service

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
	"github.com/go-playground/validator/v10"
)

type (
	AuthService interface {
		// To login, return either error or logged in user
		Login(loginRequest *model.LoginRequest) (string, error)

		// To logout, return error
		Logout(logoutRequest *model.LogoutRequest) error
	}

	AuthServiceImpl struct {
		UserRepository repository.UserRepository
	}
)

func NewAuthService(userRepository *repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		UserRepository: *userRepository,
	}
}

func (authService *AuthServiceImpl) Login(loginRequest *model.LoginRequest) (string, error) {
	var user model.User
	var err error

	validate := validator.New()
	if err := validate.Struct(loginRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return "", validationErrors
	}

	user, err = authService.UserRepository.LoginUser(loginRequest)
	if err != nil {
		return "", err
	}

	if err = authService.UserRepository.UpdateUserKey(&user, helper.RandStringBytesMaskImprSrcUnsafe(40)); err != nil {
		return "", err
	}

	jwt := NewJWT()

	jwt.SetSecret(config.GetEnv("APP_KEY", "mysecretpassword")).
		SetSub(user.Key).
		GenerateToken()

	token := jwt.GetToken()

	return token, nil
}

func (authService *AuthServiceImpl) Logout(logoutRequest *model.LogoutRequest) error {
	validate := validator.New()
	if err := validate.Struct(logoutRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	jwt := NewJWT()

	jwt.SetSecret(config.GetEnv("APP_KEY", "mysecretpassword"))

	// Check if token is valid
	if err := jwt.ValidateToken(logoutRequest.Token); err != nil {
		return err
	}
	return nil
}
