package repository

import "gorm.io/gorm"

type RoleRepositoryImpl struct {
	Db *gorm.DB
}

func (roleRepository *RoleRepositoryImpl) GetRoles() {

}
