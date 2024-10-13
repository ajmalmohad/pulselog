package models

import (
	"time"
)

type RefreshToken struct {
	Base

	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"type:text;not null;unique" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}
