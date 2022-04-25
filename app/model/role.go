package model

import "time"

type (
	Role struct {
		ID        uint      `json:"id" gorm:"primarykey"`
		Name      string    `json:"name" gorm:"unique;not null"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func NewRole() *Role {
	return &Role{}
}
