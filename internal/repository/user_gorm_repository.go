package repository

import (
	"github.com/zulal/go-backend/internal/domain"
	"gorm.io/gorm"
)

type userGormRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userGormRepo{db: db}
}

func (r *userGormRepo) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userGormRepo) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userGormRepo) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userGormRepo) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userGormRepo) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userGormRepo) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
