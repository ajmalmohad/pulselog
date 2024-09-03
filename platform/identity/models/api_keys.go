package models

type APIKey struct {
	Base

	Key       string `gorm:"not null;uniqueIndex"`
	ProjectID uint   `gorm:"not null;index"`
	CreatedBy uint   `gorm:"index"`

	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Creator User    `gorm:"foreignKey:CreatedBy;constraint:OnDelete:SET NULL"`
}
