package repository

import "github.com/zulal/go-backend/internal/domain"

// UserRepository: Veritabanı işlemleri için arayüz
type UserRepository interface {
	Create(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
	GetByID(id uint) (*domain.User, error)
	GetAll() ([]domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
}
