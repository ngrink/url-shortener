package users

import "testing"

func TestCreateUser(t *testing.T) {
	mockUsersRepo := NewUsersMockRepository()
	usersService := NewUsersService(mockUsersRepo)

	t.Run("Success", func(t *testing.T) {
		data := CreateUserDto{
			Name:     "John Doe",
			Email:    "johndoe@gmail.com",
			Password: "strongsecret",
		}

		_, err := usersService.CreateUser(data)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Invalid payload", func(t *testing.T) {
		data := CreateUserDto{
			Name:     "",
			Email:    "NOT_A_EMAIL",
			Password: "weak",
		}

		_, err := usersService.CreateUser(data)
		if err == nil {
			t.Error("CreateUser() should have failed")
		}
	})
}
