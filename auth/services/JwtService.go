package services

import (
	"main/auth/provider"
)

type JwtService struct {
	jwtProvider *provider.JwtProviderImpl
}

func NewJwtService(jwtProvider *provider.JwtProviderImpl) *JwtService {
	return &JwtService{jwtProvider: jwtProvider}
}

func (this *JwtService) encode(id string, kind string, expirationTime int) (string, error) {
	return this.jwtProvider.Encode(id, kind, expirationTime)
}

func (this *JwtService) decode(token string) (string, string, error) {
	return this.jwtProvider.Decode(token)
}
