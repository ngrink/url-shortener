package users

import (
	"fmt"
	"log"

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
	usersRepository IUsersRepository
}

func NewUsersService(usersRepository IUsersRepository) *UsersService {
	return &UsersService{usersRepository: usersRepository}
}

func (s *UsersService) CreateUser(data CreateUserDto) (User, error) {
	_, err := s.usersRepository.GetUserByEmail(data.Email)
	if err == nil {
		return User{}, fmt.Errorf("User with email \"%s\" already exists", data.Email)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hash),
	}

	user, err = s.usersRepository.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	log.Println("User created")
	return user, nil
}

func (s *UsersService) GetAllUsers() ([]User, error) {
	users, err := s.usersRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UsersService) GetUser(id uint64) (User, error) {
	user, err := s.usersRepository.GetUser(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UsersService) UpdateUser(id uint64, data UpdateUserDto) (User, error) {
	user, err := s.usersRepository.UpdateUser(id, data)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UsersService) DeleteUser(id uint64) error {
	err := s.usersRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
