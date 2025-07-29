package database

import (
	"fmt"

	"github.com/zulal/go-backend/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println("Migrating tables...")
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Transaction{},
		&domain.Balance{},
		&domain.AuditLog{},
	)
	if err != nil {
		return err
	}
	fmt.Println("✅ All tables migrated.")
	return nil
}
