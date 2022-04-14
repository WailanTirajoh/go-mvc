package repository

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/go-simple-clean-architecture/app/model"
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

func (userRepository *UserRepositoryImpl) GetUser(userId string) (model.User, error) {
	var user model.User

	if err := userRepository.Collection.First(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
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
