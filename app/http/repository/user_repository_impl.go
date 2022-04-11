package repository

import (
	"github.com/wailantirajoh/gorilla/app/model"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Collection: db,
	}
}

type UserRepositoryImpl struct {
	Collection *gorm.DB
}

func (userRepository *UserRepositoryImpl) GetUsers() []model.User {
	var users []model.User

	userRepository.Collection.Find(&users)

	return users
}

func (userRepository *UserRepositoryImpl) GetUser(userId string) model.User {
	var user model.User

	userRepository.Collection.First(&user, userId)

	return user
}

func (userRepository *UserRepositoryImpl) StoreUser(user *model.User) {
	userRepository.Collection.Create(&user)
}

func (userRepository *UserRepositoryImpl) UpdateUser(user *model.User) {
	userRepository.Collection.Save(&user)
}

func (userRepository *UserRepositoryImpl) DeleteUser(user *model.User) {
	userRepository.Collection.Delete(user)
}
