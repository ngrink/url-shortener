package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserDto struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UpdateUserDto struct {
	Name string `json:"name" db:"name"`
}
