package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" db:"name" validate:"required,max=255"`
	Email    string `json:"email" db:"email" validate:"required,email,max=255"`
	Password string `json:"-" db:"password" validate:"required,min=8,max=255"`
}

type CreateUserDto struct {
	Name     string `json:"name" db:"name" validate:"required,max=255"`
	Email    string `json:"email" db:"email" validate:"required,email,max=255"`
	Password string `json:"password" db:"password" validate:"required,min=8,max=255"`
}

type UpdateUserDto struct {
	Name string `json:"name" db:"name" validate:"required,max=255"`
}
