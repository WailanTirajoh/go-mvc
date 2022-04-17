package service

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/helper"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/go-playground/validator/v10"
)

type (
	AuthService interface {
		// To login, return either error or logged in user
		Login(loginRequest *model.LoginRequest) (string, error)

		// To logout, return error
		Logout(token string) error

		// To validate user token from string
		ValidateUserToken(token string) (model.User, error)

		// To register user
		RegisterUser(registerRequest *model.RegisterRequest) (model.User, error)
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

	jwt.
		SetPayload(Payload{
			"id":    user.ID,
			"email": user.Email,
		}).
		SetSub(user.Key).
		GenerateToken()

	token := jwt.GetToken()

	return token, nil
}

func (authService *AuthServiceImpl) Logout(token string) error {
	var user model.User
	var payload BasePayload

	split := strings.Split(token, ".")
	bytePayload, err := helper.Base64UrlDecoding(split[1])

	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytePayload, &payload); err != nil {
		return err
	}

	if err := authService.UserRepository.DeleteUserKey(&user, payload.SUB); err != nil {
		return err
	}

	return nil
}

func (authService *AuthServiceImpl) RegisterUser(registerRequest *model.RegisterRequest) (model.User, error) {
	var user model.User
	var err error

	validate := validator.New()
	if err = validate.Struct(registerRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return user, validationErrors
	}

	user = model.User{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Email:     registerRequest.Email,
		Password:  authService.UserRepository.GeneratePassword(registerRequest.Email, registerRequest.Password),
		CreatedAt: time.Now(),
	}

	if err = authService.UserRepository.StoreUser(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (authService *AuthServiceImpl) ValidateUserToken(token string) (model.User, error) {
	var user model.User
	var payload Payload

	split := strings.Split(token, ".")
	bytePayload, err := helper.Base64UrlDecoding(split[1])

	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(bytePayload, &payload); err != nil {
		return user, err
	}

	sub, err := helper.GetStrKey(payload, "sub")
	if err != nil {
		return user, err
	}

	email, err := helper.GetStrKey(payload, "email")
	if err != nil {
		return user, err
	}

	if err := authService.UserRepository.FindUserByEmailKey(&user, sub, email); err != nil {
		return user, errors.New("token is invalid")
	}

	return user, nil
}
