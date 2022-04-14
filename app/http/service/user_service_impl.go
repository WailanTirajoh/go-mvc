package service

import (
	"encoding/json"
	"io"

	"github.com/WailanTirajoh/go-simple-clean-architecture/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/go-simple-clean-architecture/app/model"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
	}
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (userService *UserServiceImpl) GetUsers() []model.User {
	responses := userService.UserRepository.GetUsers()
	return responses
}

func (userService *UserServiceImpl) GetUser(userId string) (model.User, error) {
	user, err := userService.UserRepository.GetUser(userId)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserServiceImpl) StoreUser(rbody io.ReadCloser) model.User {
	var user model.User

	json.NewDecoder(rbody).Decode(&user)

	userService.UserRepository.StoreUser(&user)

	return user
}

func (userService *UserServiceImpl) UpdateUser(userId string, rbody io.ReadCloser) (model.User, error) {
	user, err := userService.UserRepository.GetUser(userId)

	if err != nil {
		return user, err
	}

	json.NewDecoder(rbody).Decode(&user)
	userService.UserRepository.UpdateUser(&user)

	return user, nil
}

func (userService *UserServiceImpl) DeleteUser(userId string) (string, error) {
	user, err := userService.UserRepository.GetUser(userId)

	if err != nil {
		return "", err
	}

	userService.UserRepository.DeleteUser(&user)

	return "User deleted successfully!", nil
}
