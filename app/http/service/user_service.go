package service

import (
	"io"

	"github.com/wailantirajoh/gorilla/app/model"
)

type UserService interface {
	GetUsers() []model.User
	GetUser(userId string) model.User
	StoreUser(rbody io.ReadCloser) model.User
	UpdateUser(userId string, rbody io.ReadCloser) model.User
	DeleteUser(userId string) string
}
