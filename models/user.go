package models

type User struct {
	ID              int           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email           string        `gorm:"size:80;not null" json:"email"`
	Password        string        `gorm:"size:100;not null" json:"password"`
	Token           string        `gorm:"size:50" json:"token"`
	UserStateTypeID int           `gorm:"not null" json:"-"`
	UserTypeID      int           `gorm:"not null" json:"-"`
	UserType        UserType      `gorm:"foreignKey:UserTypeID;references:ID" json:"user_type"`
	UserStateType   UserStateType `gorm:"foreignKey:UserStateTypeID;references:ID" json:"user_state_type"`
}
