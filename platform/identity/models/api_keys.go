package models

type APIKey struct {
	Base

	Key       string `gorm:"not null;uniqueIndex" json:"key"`
	ProjectID uint   `gorm:"not null;index" json:"project_id"`
	CreatedBy uint   `gorm:"index" json:"created_by"`

	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"project"`
	Creator User    `gorm:"foreignKey:CreatedBy;constraint:OnDelete:SET NULL" json:"creator"`
}
