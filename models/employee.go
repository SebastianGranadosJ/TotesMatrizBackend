package models

type Employee struct {
	ID               int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Names            string         `gorm:"size:100;not null" json:"names"`
	LastNames        string         `gorm:"size:100;not null" json:"last_names"`
	PersonalID       string         `gorm:"size:50;not null" json:"personal_id"`
	Address          string         `gorm:"size:200" json:"address,omitempty"`
	PhoneNumbers     string         `gorm:"size:50" json:"phone_numbers,omitempty"`
	UserID           int            `gorm:"not null" json:"user_id"`
	User             User           `gorm:"foreignKey:UserID;references:ID" json:"user"`
	IdentifierTypeID int            `gorm:"not null" json:"identifier_type_id"`
	IdentifierType   IdentifierType `gorm:"foreignKey:IdentifierTypeID;references:ID" json:"identifier_type"`
}
