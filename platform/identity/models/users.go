package models

type User struct {
	Base

	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;unique;not null"`
	Password string `gorm:"size:255;not null"`
	IsActive bool   `gorm:"default:true"`
}
