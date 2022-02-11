package services

import (
	"main/auth/provider"
	"main/user/entities"
)

type JwtService struct {
	jwtProvider *provider.JwtProviderImpl
}

func NewJwtService(jwtProvider *provider.JwtProviderImpl) *JwtService {
	return &JwtService{jwtProvider: jwtProvider}
}

func (this *JwtService) encode(user *entities.User) string {
	return ""
}

func (this *JwtService) decode(token string) any {
	return nil
}
