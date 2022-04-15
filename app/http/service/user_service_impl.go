package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
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

func (userService *UserServiceImpl) StoreUser(userRequest *model.StoreUserRequest) (model.User, error) {
	var user model.User

	var text = userRequest.Email + userRequest.Password + config.GetEnv("APP_KEY", "mysecretpassword")
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	user = model.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  encryptedString,
		CreatedAt: time.Now(),
	}

	err := userService.UserRepository.StoreUser(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserServiceImpl) UpdateUser(userRequest *model.UpdateUserRequest, userId string) (model.User, error) {
	var user model.User

	theuser, err := userService.UserRepository.GetUser(userId)

	if err != nil {
		return user, err
	}

	user = model.User{
		ID:        theuser.ID,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     theuser.Email,
		Password:  theuser.Password,
		CreatedAt: theuser.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if err = userService.UserRepository.UpdateUser(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserServiceImpl) DeleteUser(userId string) error {
	user, err := userService.UserRepository.GetUser(userId)

	if err != nil {
		return err
	}

	return userService.UserRepository.DeleteUser(&user)
}
