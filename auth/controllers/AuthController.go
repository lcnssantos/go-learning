package controllers

import "main/auth/services"

type AuthController struct {
	authService *services.AuthService
	jwtService  *services.JwtService
}

func NewAuthController(authService *services.AuthService, jwtService *services.JwtService) *AuthController {
	return &AuthController{authService: authService, jwtService: jwtService}
}
