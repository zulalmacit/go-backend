package database

import (
	"github.com/zulal/go-backend/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.User{})
}
