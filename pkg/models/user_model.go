package models

import (
	"time"

	"github.com/google/uuid"
)


// User struct to describe the User object
type User struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	AccountStatus bool      `json:"account_status"`
	IsOnline      bool      `json:"is_online"`
	Name          string    `json:"name" `
	Email         string    `json:"email" gorm:"uniqueIndex"`
	Phone         string    `json:"phone" gorm:"uniqueIndex"`
	Password      string    `json:"password"`
	UserRole      string    `json:"user_role"`
	PhotoUrl      string    `json:"photo_url"`
}
