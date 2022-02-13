package dto

type AuthDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type AuthResponseDto struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshRequestDto struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
