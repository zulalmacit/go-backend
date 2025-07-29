package domain

import (
	"errors"
	"time"
)

// Balance: Her kullanıcının güncel bakiyesini temsil eder.
type Balance struct {
	UserID        uint      `gorm:"primaryKey" json:"user_id"` // Kullanıcı ID (aynı zamanda primary key)
	Amount        float64   `json:"amount"`                    // Kullanıcının bakiyesi
	LastUpdatedAt time.Time `json:"last_updated_at"`           // Son güncellenme zamanı
}

// Validate: Geçerli bir bakiye olup olmadığını kontrol eder.
func (b *Balance) Validate() error {
	if b.UserID == 0 {
		return errors.New("user_id is required")
	}
	if b.Amount < 0 {
		return errors.New("amount cannot be negative")
	}
	return nil
}
