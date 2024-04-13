package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ngrink/url-shortener/internal/modules/users"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	usersService *users.UsersService
}

func NewAuthService(
	usersService *users.UsersService,
) *AuthService {
	return &AuthService{
		usersService: usersService,
	}
}

func (s *AuthService) Login(user users.User) LoginResponse {
	token, err := generateJwtToken(user.ID)
	if err != nil {
		log.Fatal(err)
	}

	return LoginResponse{
		Token: token,
		User:  user,
	}
}

func (s *AuthService) LoginByCredentials(login, password string) (LoginResponse, error) {
	user, err := s.usersService.GetUserByEmail(login)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("user not found or password is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return LoginResponse{}, fmt.Errorf("user not found or password is incorrect")
	}

	data := s.Login(user)

	return data, nil
}

func generateJwtToken(userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJwtToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
