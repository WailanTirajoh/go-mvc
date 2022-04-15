package service

import (
	"time"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
	}
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (ur *UserServiceImpl) GetUsers() []model.User {
	responses := ur.UserRepository.GetUsers()
	return responses
}

func (ur *UserServiceImpl) GetUser(userId string) (model.User, error) {
	user, err := ur.UserRepository.GetUser(userId)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserServiceImpl) StoreUser(user *model.User) error {
	return ur.UserRepository.StoreUser(user)
}

func (ur *UserServiceImpl) UpdateUser(userId string, user *model.User) error {
	tuser, err := ur.UserRepository.GetUser(userId)

	user.ID = tuser.ID
	user.CreatedAt = tuser.CreatedAt
	user.UpdatedAt = time.Now()

	if err != nil {
		return err
	}

	return ur.UserRepository.UpdateUser(user)
}

func (ur *UserServiceImpl) DeleteUser(userId string) error {
	user, err := ur.UserRepository.GetUser(userId)

	if err != nil {
		return err
	}

	return ur.UserRepository.DeleteUser(&user)
}
