package models

type OrderStateType struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Description string `gorm:"not null;size:300" json:"description"`
}
