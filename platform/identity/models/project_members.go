package models

type Role string

const (
	ADMIN  Role = "ADMIN"
	MEMBER Role = "MEMBER"
)

type ProjectMember struct {
	Base

	ProjectID uint `gorm:"not null;index" json:"project_id"`
	UserID    uint `gorm:"not null;index" json:"user_id"`
	Role      Role `gorm:"not null;index;default:MEMBER" json:"role"`

	Project Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"project"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}
