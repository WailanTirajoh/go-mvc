package repository

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
	"github.com/WailanTirajoh/go-simple-clean-architecture/config"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
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
	}

	UserRepositoryImpl struct {
		Collection *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Collection: db,
	}
}

func (ur *UserRepositoryImpl) GetUsers() []model.User {
	var users []model.User

	ur.Collection.Find(&users)

	return users
}

func (ur *UserRepositoryImpl) GetUser(userId string) (model.User, error) {
	var user model.User

	if err := ur.Collection.First(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) StoreUser(user *model.User) error {
	return ur.Collection.Create(&user).Error
}

func (ur *UserRepositoryImpl) UpdateUser(user *model.User) error {
	return ur.Collection.Save(&user).Error
}

func (ur *UserRepositoryImpl) DeleteUser(user *model.User) error {
	return ur.Collection.Delete(user).Error
}

func (ur *UserRepositoryImpl) GeneratePassword(email string, password string) string {
	strPassword := email + password + config.GetEnv("APP_KEY", "mysecretpassword")
	var sha = sha1.New()
	sha.Write([]byte(strPassword))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted)
}

func (ur *UserRepositoryImpl) LoginUser(loginRequest *model.LoginRequest) (model.User, error) {
	var user model.User

	err := ur.Collection.Where(&model.User{
		Email:    loginRequest.Email,
		Password: ur.GeneratePassword(loginRequest.Email, loginRequest.Password),
	}).First(&user).Error

	if err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}

func (ur *UserRepositoryImpl) UpdateUserKey(user *model.User, key string) error {
	user.Key = key

	return ur.Collection.Save(&user).Error
}
