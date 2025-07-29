package domain

import (
	"errors"
	"time"
)

// Transaction: Kullanıcılar arası para transferi, kredi, ödeme gibi işlemleri temsil eder.
type Transaction struct {
	ID         uint      `gorm:"primaryKey" json:"id"` // Otomatik artan anahtar
	FromUserID uint      `json:"from_user_id"`         // Gönderen kullanıcı (nullable olabilir)
	ToUserID   uint      `json:"to_user_id"`           // Alıcı kullanıcı
	Amount     float64   `json:"amount"`               // Tutar
	Type       string    `json:"type"`                 // "credit", "debit", "transfer"
	Status     string    `json:"status"`               // "pending", "completed", "failed"
	CreatedAt  time.Time `json:"created_at"`           // Oluşturulma zamanı
}

// Validate: Geçerli bir işlem verisi olup olmadığını kontrol eder.
// API'den gelen verilerde çağrılarak hatalı işlemler engellenir.
func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.Type == "" {
		return errors.New("type is required")
	}
	if t.Status == "" {
		return errors.New("status is required")
	}
	if t.ToUserID == 0 {
		return errors.New("to_user_id is required")
	}
	return nil
}
