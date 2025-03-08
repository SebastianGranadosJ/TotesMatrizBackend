package dtos

type GetUserDTO struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	UserType  int    `json:"user_type"`
	UserState int    `json:"user_state"`
}
