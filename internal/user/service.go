package user

import (
	"errors"

	"github.com/joaodemutti/user-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *Service) Register(name, email, password string) (User, error) {
	// Check if user already exists
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return User{}, errors.New("email already exists")
	}

	// If error is NOT "record not found", something else failed
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, err
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	return s.repo.Create(user)
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
