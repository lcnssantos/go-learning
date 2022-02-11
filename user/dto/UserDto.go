package dto

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}
