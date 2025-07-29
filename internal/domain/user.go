package domain

import (
	"errors"
	"time"
)

// User: Sistemdeki kullanıcıları temsil eder.
type User struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	Name         string
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"` // JSON'e dahil etme
	Role         string    `json:"role"`              // Örn: "admin", "user"
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Validate: Kullanıcı verisinin geçerli olup olmadığını kontrol eder.
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.PasswordHash == "" {
		return errors.New("password is required")
	}
	return nil
}
