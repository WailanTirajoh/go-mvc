package repository

import "github.com/wailantirajoh/gorilla/app/model"

type UserRepository interface {
	GetUsers() []model.User
	GetUser(userId string) model.User
	StoreUser(user *model.User)
	UpdateUser(user *model.User)
	DeleteUser(user *model.User)
}
