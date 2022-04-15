package service

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
)

type (
	UserService interface {
		GetUsers() []model.User
		GetUser(userId string) (model.User, error)
		StoreUser(user *model.User) error
		UpdateUser(userId string, user *model.User) error
		DeleteUser(userId string) error
	}
)
