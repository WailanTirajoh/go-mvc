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

		// To find user by key
		FindUserByKey(user *model.User, key string) error

		// To delete user key
		DeleteUserKey(user *model.User, key string) error
	}

	UserRepositoryImpl struct {
		Db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (userRepository *UserRepositoryImpl) GetUsers() []model.User {
	var users []model.User
	userRepository.Db.Find(&users)
	return users
}

func (userRepository *UserRepositoryImpl) GetUser(userId string) (model.User, error) {
	var user model.User
	if err := userRepository.Db.First(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (userRepository *UserRepositoryImpl) StoreUser(user *model.User) error {
	tx := userRepository.Db.Begin()
	if err := userRepository.Db.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (userRepository *UserRepositoryImpl) UpdateUser(user *model.User) error {
	tx := userRepository.Db.Begin()
	if err := userRepository.Db.Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (userRepository *UserRepositoryImpl) DeleteUser(user *model.User) error {
	tx := userRepository.Db.Begin()
	if err := userRepository.Db.Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (userRepository *UserRepositoryImpl) GeneratePassword(email string, password string) string {
	strPassword := email + password + config.GetEnv("APP_KEY", "mysecretpassword")
	var sha = sha1.New()
	sha.Write([]byte(strPassword))
	var encrypted = sha.Sum(nil)
	return fmt.Sprintf("%x", encrypted)
}

func (userRepository *UserRepositoryImpl) LoginUser(loginRequest *model.LoginRequest) (model.User, error) {
	var user model.User
	err := userRepository.Db.Where(&model.User{
		Email:    loginRequest.Email,
		Password: userRepository.GeneratePassword(loginRequest.Email, loginRequest.Password),
	}).First(&user).Error
	if err != nil {
		return user, errors.New("invalid credentials")
	}
	return user, nil
}

func (userRepository *UserRepositoryImpl) UpdateUserKey(user *model.User, key string) error {
	tx := userRepository.Db.Begin()
	user.Key = key
	if err := userRepository.Db.Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (userRepository *UserRepositoryImpl) FindUserByKey(user *model.User, key string) error {
	return userRepository.Db.Where(&model.User{
		Key: key,
	}).First(&user).Error
}

func (userRepository *UserRepositoryImpl) DeleteUserKey(user *model.User, key string) error {
	err := userRepository.Db.Where(&model.User{
		Key: key,
	}).First(&user).Error

	if err != nil {
		return err
	}

	user.Key = ""
	if err := userRepository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
