package models

type Role struct {
	ID          uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string       `gorm:"size:100;not null" json:"name"`
	Description string       `gorm:"size:300" json:"description,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permission;" json:"permissions"`
}
