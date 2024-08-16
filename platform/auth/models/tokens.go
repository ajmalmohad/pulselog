package models

import (
    "time"
)

type RefreshToken struct {
    ID        uint           `gorm:"primaryKey"`
    UserID    uint           `gorm:"not null;index"`
    Token     string         `gorm:"type:text;not null;unique"`
    ExpiresAt time.Time      `gorm:"not null"`
    CreatedAt time.Time      `gorm:"autoCreateTime"`
	
    User      User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}