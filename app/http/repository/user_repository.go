package repository

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"gorm.io/gorm"
)

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

	// To generate user password (email+password+saltsecret) using sha1
	GeneratePassword(email string, password string) string

	// To login
	LoginUser(loginRequest *model.LoginRequest) (model.User, error)

	// To update user key after login
	UpdateUserKey(user *model.User, key string) error

	// To find user by key
	FindUserByEmailKey(user *model.User, key string, email string) error

	// To delete user key
	DeleteUserKey(user *model.User, key string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}
