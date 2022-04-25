package service

import (
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
)

func NewAuthService(userRepository *repository.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: *userRepository,
		Validate:       *validate,
	}
}
