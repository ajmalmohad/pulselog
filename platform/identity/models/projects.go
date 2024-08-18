package models

type Project struct {
	Base

	Name    string `gorm:"not null"`
	OwnerID uint   `gorm:"not null;index"`

	Owner   User     `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
	APIKeys []APIKey `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}
