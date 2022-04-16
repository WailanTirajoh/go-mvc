package model

import (
	"time"
)

type (
	User struct {
		ID        uint      `json:"id" gorm:"primarykey"`
		FirstName string    `json:"first_name" gorm:"not null;type:varchar(191)"`
		LastName  string    `json:"last_name" gorm:"not null;type:varchar(191)"`
		Email     string    `json:"email" gorm:"uniqueIndex:email;index;not null"`
		Token     string    `json:"token" gorm:"uniqueIndex:token;index;not null"`
		Password  string    `json:"password" gorm:"not null;type:varchar(191)"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	StoreUserRequest struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
	}

	UpdateUserRequest struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
	}

	UserResponse struct {
		ID        uint   `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func (u *User) Response() UserResponse {
	return UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
