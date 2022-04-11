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

func (userService *UserServiceImpl) GetUser(userId string) model.User {
	user := userService.UserRepository.GetUser(userId)
	return user
}

func (userService *UserServiceImpl) StoreUser(rbody io.ReadCloser) model.User {
	var user model.User

	json.NewDecoder(rbody).Decode(&user)

	userService.UserRepository.StoreUser(&user)

	return user
}

func (userService *UserServiceImpl) UpdateUser(userId string, rbody io.ReadCloser) model.User {
	user := userService.UserRepository.GetUser(userId)

	json.NewDecoder(rbody).Decode(&user)
	userService.UserRepository.UpdateUser(&user)

	return user
}

func (userService *UserServiceImpl) DeleteUser(userId string) string {
	user := userService.UserRepository.GetUser(userId)

	userService.UserRepository.DeleteUser(&user)

	return "User deleted successfully!"
}
