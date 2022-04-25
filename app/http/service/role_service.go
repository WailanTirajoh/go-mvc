package service

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/go-playground/validator/v10"
)

type RoleService interface {
}

func NewRoleService(roleRepository *repository.RoleRepository, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: *roleRepository,
		Validate:       *validate,
	}
}
