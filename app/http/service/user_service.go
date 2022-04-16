package service

import (
	"time"

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

	UserServiceImpl struct {
		UserRepository repository.UserRepository
	}
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
	}
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
	var err error

	validate := validator.New()
	if err = validate.Struct(userRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return user, validationErrors
	}

	user = model.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Email:     userRequest.Email,
		Password:  userService.UserRepository.GeneratePassword(userRequest.Email, userRequest.Password),
		CreatedAt: time.Now(),
	}

	if err = userService.UserRepository.StoreUser(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserServiceImpl) UpdateUser(userRequest *model.UpdateUserRequest, userId string) (model.User, error) {
	var user model.User
	var err error

	validate := validator.New()
	if err = validate.Struct(userRequest); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return user, validationErrors
	}

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
