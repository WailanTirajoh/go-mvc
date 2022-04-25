package service

import (
	"github.com/WailanTirajoh/go-simple-clean-architecture/app/http/repository"
	"github.com/go-playground/validator/v10"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	Validate       validator.Validate
}
