package service

import (
	"time"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func (userService *UserServiceImpl) GetUsers() []model.User {
	return userService.UserRepository.GetUsers()
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
	var err error

	if err = userService.Validate.Struct(userRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return user, validationErrors
	}

	user = model.NewUser()
	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.Email = userRequest.Email
	user.Password = userService.UserRepository.GeneratePassword(userRequest.Email, userRequest.Password)
	user.CreatedAt = time.Now()

	if err = userService.UserRepository.StoreUser(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserServiceImpl) UpdateUser(userRequest *model.UpdateUserRequest, userId string) (model.User, error) {
	var user model.User
	var err error

	if err = userService.Validate.Struct(userRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return user, validationErrors
	}

	theuser, err := userService.UserRepository.GetUser(userId)

	if err != nil {
		return user, err
	}

	user = model.NewUser()
	user.ID = theuser.ID
	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.Email = theuser.Email
	user.Password = theuser.Password
	user.CreatedAt = theuser.CreatedAt
	user.UpdatedAt = time.Now()

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
