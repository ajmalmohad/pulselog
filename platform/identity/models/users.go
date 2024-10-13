package models

type User struct {
	Base

	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}
