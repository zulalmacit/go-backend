package service

import (
	"errors"
	"time"

	"go-backend/internal/domain"
	"go-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{userRepo: r}
}

func (s *UserService) Register(user *domain.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashed)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return s.userRepo.Create(user)
}

func (s *UserService) Login(email, password string) (*domain.User, string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("kullanıcı bulunamadı")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", errors.New("şifre hatalı")
	}

	// Basit token örneği — gerçek projede JWT kullanılmalı
	token := "example-token-" + user.Email
	return user, token, nil
}
