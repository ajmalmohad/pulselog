package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model

	UserID    uint      `gorm:"not null;index"`
	Token     string    `gorm:"type:text;not null;unique"`
	ExpiresAt time.Time `gorm:"not null"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
