package service

import (
	"errors"

	"github.com/zulal/go-backend/internal/domain"
	"github.com/zulal/go-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// userService: iş mantığını tutar
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService: Kullanmak için yeni bir service oluşturur
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Register: Yeni kullanıcı oluşturur (şifreyi hashler)
func (s *UserService) Register(user *domain.User) error {
	// Kullanıcı var mı?
	existing, _ := s.userRepo.GetByEmail(user.Email)
	if existing != nil {
		return errors.New("user already exists")
	}

	// Şifreyi hashle
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	return s.userRepo.Create(user)
}

// Login: E-posta ve şifre doğrulaması
func (s *UserService) Login(email, password string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// GetByID: Kullanıcıyı ID ile bul
func (s *UserService) GetByID(id uint) (*domain.User, error) {
	return s.userRepo.GetByID(id)
}

// GetAll: Tüm kullanıcıları döner
func (s *UserService) GetAll() ([]domain.User, error) {
	return s.userRepo.GetAll()
}

// Update: Kullanıcıyı günceller
func (s *UserService) Update(user *domain.User) error {
	return s.userRepo.Update(user)
}

// Delete: Kullanıcıyı siler
func (s *UserService) Delete(id uint) error {
	return s.userRepo.Delete(id)
}
