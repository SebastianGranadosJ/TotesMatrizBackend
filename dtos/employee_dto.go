package dtos

type GetEmployeeDTO struct {
	ID               int    `json:"id"`
	Names            string `json:"names"`
	LastNames        string `json:"last_names"`
	PersonalID       string `json:"personal_id"`
	Address          string `json:"address,omitempty"`
	PhoneNumbers     string `json:"phone_numbers,omitempty"`
	UserID           int    `json:"user_id"`
	IdentifierTypeID int    `json:"identifier_type_id"`
}

type UpdateEmployeeDTO struct {
	Names            string `json:"names"`
	LastNames        string `json:"last_names"`
	PersonalID       string `json:"personal_id"`
	Address          string `json:"address,omitempty"`
	PhoneNumbers     string `json:"phone_numbers,omitempty"`
	UserID           int    `json:"user_id"`
	IdentifierTypeID int    `json:"identifier_type_id"`
}
