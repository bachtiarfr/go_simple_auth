package services

import (
	"fmt"
	"log"
	"simple-auth/internal/config"
	"simple-auth/internal/entity"
	"simple-auth/internal/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(userRepository repositories.UserRepository) AuthenticationService {
	return &authenticationService{
		userRepository: userRepository,
	}
}

func (s *authenticationService) Authenticate(email, password string) (string, error) {
	config, err := config.ReadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return "", err
	}

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil || !s.verifyPassword(user, password) {
		return "", fmt.Errorf("invalid email or password")
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.SecretKey))

	return tokenString, err
}

func (s *authenticationService) Register(fullname string, email string, age int, phoneNumber string, hashedPassword string) (entity.User, error) {
	println("email : ", email)
	user, err := s.userRepository.FindByEmail(email)
	if user != nil {
		return entity.User{}, fmt.Errorf("Email exist")
	}

	dataUser := entity.User{
		Fullname:    fullname,
		Email:       email,
		Age:         age,
		PhoneNumber: phoneNumber,
		Password:    hashedPassword,
	}

	newUser, err := s.userRepository.CreateNewUser(&dataUser)
	if user != nil {
		return entity.User{}, fmt.Errorf("Email exist")
	}
	fmt.Printf("newUser :", newUser)
	return entity.User{}, err
}

func (s *authenticationService) verifyPassword(user *entity.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
