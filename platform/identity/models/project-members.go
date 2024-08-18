package models

type Role string

const (
	ADMIN  Role = "ADMIN"
	MEMBER Role = "MEMBER"
)

type ProjectMember struct {
	Base

	ProjectID uint `gorm:"not null;index"`
	UserID    uint `gorm:"not null;index"`
	Role      Role `gorm:"not null;index;default:MEMBER"`

	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
