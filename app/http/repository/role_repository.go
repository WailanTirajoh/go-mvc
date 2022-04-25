package repository

import "gorm.io/gorm"

type RoleRepository interface {
	GetRoles()
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		Db: db,
	}
}
