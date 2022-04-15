package repository

import "github.com/WailanTirajoh/go-simple-clean-architecture/app/model"

type UserRepository interface {
	// To get all user
	GetUsers() []model.User

	// To get specific user by ID
	GetUser(userId string) (model.User, error)

	// To store user
	StoreUser(user *model.User) error

	// To update user
	UpdateUser(user *model.User) error

	// To delete user
	DeleteUser(user *model.User) error
}
