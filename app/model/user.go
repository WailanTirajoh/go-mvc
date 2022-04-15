package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	FirstName string         `json:"first_name" gorm:"not null" sqlType:"text"`
	LastName  string         `json:"last_name" gorm:"not null" sqlType:"text"`
	Email     string         `json:"email" gorm:"uniqueIndex:email;index;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type StoreUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
