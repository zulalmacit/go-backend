package domain

import (
	"errors"
	"time"
)

// AuditLog: Sistemdeki önemli olayların izlenmesini sağlar.
type AuditLog struct {
	ID         uint      `gorm:"primaryKey" json:"id"` // Otomatik artan ID
	EntityType string    `json:"entity_type"`          // Örneğin: "user", "transaction"
	EntityID   uint      `json:"entity_id"`            // İlgili varlık (örneğin: user ID)
	Action     string    `json:"action"`               // Örn: "create", "update", "delete"
	Details    string    `json:"details"`              // Ek açıklama (örn: eski değer -> yeni değer)
	CreatedAt  time.Time `json:"created_at"`           // İşlem zamanı
}

// Validate: Zorunlu alanların dolu olduğunu kontrol eder.
func (a *AuditLog) Validate() error {
	if a.EntityType == "" {
		return errors.New("entity_type is required")
	}
	if a.EntityID == 0 {
		return errors.New("entity_id is required")
	}
	if a.Action == "" {
		return errors.New("action is required")
	}
	return nil
}
