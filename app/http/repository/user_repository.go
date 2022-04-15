package repository

import "github.com/WailanTirajoh/go-simple-clean-architecture/app/model"

type UserRepository interface {
	GetUsers() []model.User
	GetUser(userId string) (model.User, error)
	StoreUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error
}
