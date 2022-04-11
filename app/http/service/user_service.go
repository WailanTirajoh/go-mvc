package service

import (
	"io"

	"github.com/WailanTirajoh/go-simple-clean-architecture/go-simple-clean-architecture/app/model"
)

type UserService interface {
	GetUsers() []model.User
	GetUser(userId string) model.User
	StoreUser(rbody io.ReadCloser) model.User
	UpdateUser(userId string, rbody io.ReadCloser) model.User
	DeleteUser(userId string) string
}
