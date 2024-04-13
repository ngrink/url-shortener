package users

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/ngrink/url-shortener/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type IUsersService interface {
	CreateUser(data CreateUserDto) (User, error)
	GetAllUsers() ([]User, error)
	GetUser(id uint64) (User, error)
	UpdateUser(id uint64, data UpdateUserDto) (User, error)
	DeleteUser(id uint64) error
}

type UsersService struct {
	repository IUsersRepository
}

func NewUsersService(repository IUsersRepository) *UsersService {
	return &UsersService{repository: repository}
}

func (s *UsersService) CreateUser(data CreateUserDto) (User, error) {
	// validate payload
	if err := utils.Validate.Struct(data); err != nil {
		errors := err.(validator.ValidationErrors)
		return User{}, fmt.Errorf("invalid payload %v", errors)
	}

	// check if user not already exists
	_, err := s.repository.GetUserByEmail(data.Email)
	if err == nil {
		return User{}, fmt.Errorf("User with email \"%s\" already exists", data.Email)
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	// save user to database
	user := User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hash),
	}
	user, err = s.repository.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	// log user creation
	log.Println("User created")

	return user, nil
}

func (s *UsersService) GetAllUsers() ([]User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UsersService) GetUser(id uint64) (User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UsersService) GetUserByEmail(email string) (User, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UsersService) UpdateUser(id uint64, data UpdateUserDto) (User, error) {
	user, err := s.repository.UpdateUser(id, data)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UsersService) DeleteUser(id uint64) error {
	err := s.repository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
