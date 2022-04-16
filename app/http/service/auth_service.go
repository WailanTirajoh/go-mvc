package service

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/go-playground/validator/v10"
)

type (
	AuthService interface {
		// To login, return either error or logged in user
		Login(loginRequest *model.LoginRequest) (model.User, error)
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

func (authService *AuthServiceImpl) Login(loginRequest *model.LoginRequest) (model.User, error) {
	var user model.User
	var err error

	validate := validator.New()
	if err := validate.Struct(loginRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return user, validationErrors
	}

	user, err = authService.UserRepository.LoginUser(loginRequest)
	if err != nil {
		return user, err
	}

	return user, nil
}
