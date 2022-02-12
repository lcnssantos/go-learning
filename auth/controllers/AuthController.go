package controllers

import (
	"encoding/json"
	"main/auth/dto"
	"main/auth/services"
	"main/shared"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
	jwtService  *services.JwtService
}

func (this *AuthController) Auth(writer http.ResponseWriter, request *http.Request) {
	authDto := dto.AuthDto{}
	err := json.NewDecoder(request.Body).Decode(&authDto)

	if err != nil {
		shared.ThrowHttpError(writer, http.StatusBadRequest, "Failed to parse request")
		return
	}

	user, err1 := this.authService.Validate(authDto.Email, authDto.Password)

	if err1 != nil {
		shared.ThrowHttpError(writer, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err2 := this.authService.CreateToken(user)

	if err2 != nil {
		shared.ThrowHttpError(writer, http.StatusInternalServerError, "Internal server error")
		return
	}

	refreshToken, err3 := this.authService.CreateRefreshToken(user)

	if err3 != nil {
		shared.ThrowHttpError(writer, http.StatusInternalServerError, "Internal server error")
		return
	}

	authResponseDto := dto.AuthResponseDto{Type: "bearer", Token: token, RefreshToken: refreshToken}

	shared.SendHttpResponse(writer, http.StatusOK, authResponseDto)
}

func (this *AuthController) Refresh(w http.ResponseWriter, r *http.Request) {
	requestDto := dto.RefreshRequestDto{}

	err := json.NewDecoder(r.Body).Decode(&requestDto)

	if err != nil {
		shared.ThrowHttpError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	token, refreshToken, err2 := this.authService.RefreshToken(requestDto.RefreshToken)

	if err2 != nil {
		shared.ThrowHttpError(w, http.StatusBadRequest, "Internal server Error")
		return
	}

	authResponseDto := dto.AuthResponseDto{Type: "bearer", Token: token, RefreshToken: refreshToken}

	shared.SendHttpResponse(w, http.StatusOK, authResponseDto)
}

func NewAuthController(authService *services.AuthService, jwtService *services.JwtService) *AuthController {
	return &AuthController{authService: authService, jwtService: jwtService}
}
