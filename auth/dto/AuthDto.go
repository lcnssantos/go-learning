package dto

type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshRequestDto struct {
	RefreshToken string `json:"refreshToken"`
}
