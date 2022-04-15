package repository

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
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
