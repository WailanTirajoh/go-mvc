package service

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/go-playground/validator/v10"
)

type (
	UserService interface {
		// To get all user
		GetUsers() []model.User

		// To get specific user by ID
		GetUser(userId string) (model.User, error)

		// To store user
		StoreUser(userRequest *model.StoreUserRequest) (model.User, error)

		// To update user by ID
		UpdateUser(userRequest *model.UpdateUserRequest, userId string) (model.User, error)

		// To delete user by ID
		DeleteUser(userId string) error
	}
)

func NewUserService(userRepository *repository.UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
		Validate:       validator,
	}
}
