package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model

    Name      string         `gorm:"size:255;not null"`
    Email     string         `gorm:"size:255;unique;not null"`
	Password  string         `gorm:"size:255;not null"`
	IsActive  bool           `gorm:"default:true"`
}