package controllers

import (
	"main/auth/dto"
	"main/auth/services"
	"main/shared"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
	jwtService  *services.JwtService
}

func (this *AuthController) Auth(w http.ResponseWriter, r *http.Request) {
	authDto := &dto.AuthDto{}

	err := shared.HandleValidateRequest(w, r, authDto)

	if err != nil {
		return
	}

	user, passwordValidationErr := this.authService.Validate(authDto.Email, authDto.Password)

	if passwordValidationErr != nil {
		shared.ThrowHttpError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, createTokenErr := this.authService.CreateToken(user)

	if createTokenErr != nil {
		shared.ThrowHttpError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	refreshToken, createRefreshTokenErr := this.authService.CreateRefreshToken(user)

	if createRefreshTokenErr != nil {
		shared.ThrowHttpError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	shared.SendHttpResponse(w, http.StatusOK, dto.AuthResponseDto{Type: "bearer", Token: token, RefreshToken: refreshToken})
}

func (this *AuthController) Refresh(w http.ResponseWriter, r *http.Request) {
	requestDto := &dto.RefreshRequestDto{}

	err := shared.HandleValidateRequest(w, r, requestDto)

	if err != nil {
		return
	}

	token, refreshToken, refreshTokenErr := this.authService.RefreshToken(requestDto.RefreshToken)

	if refreshTokenErr != nil {
		shared.ThrowHttpError(w, http.StatusBadRequest, "Internal server Error")
		return
	}

	shared.SendHttpResponse(w, http.StatusOK, dto.AuthResponseDto{Type: "bearer", Token: token, RefreshToken: refreshToken})
}

func NewAuthController(authService *services.AuthService, jwtService *services.JwtService) *AuthController {
	return &AuthController{authService: authService, jwtService: jwtService}
}
