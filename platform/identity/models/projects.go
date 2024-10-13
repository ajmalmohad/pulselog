package models

type Project struct {
	Base

	Name    string `gorm:"not null" json:"name"`
	OwnerID uint   `gorm:"not null;index" json:"owner_id"`

	Owner          User            `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE" json:"owner"`
	APIKeys        []APIKey        `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"api_keys"`
	ProjectMembers []ProjectMember `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"project_members"`
}
