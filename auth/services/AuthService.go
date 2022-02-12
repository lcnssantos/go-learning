package services

import (
	"errors"
	"main/user/entities"
	"main/user/services"
)

type AuthService struct {
	userService *services.UserService
	hashService *services.HashService
	jwtService  *JwtService
}

func NewAuthService(userService *services.UserService, hashService *services.HashService, jwtService *JwtService) *AuthService {
	return &AuthService{userService: userService, hashService: hashService, jwtService: jwtService}
}

func (this *AuthService) Validate(email string, password string) (*entities.User, error) {
	user, err := this.userService.FindOneByEmail(email)

	if err == nil {
		isPasswordValid := this.hashService.Compare(user.Password, password)

		if isPasswordValid && user.IsActive && user.EmailConfirmed {
			return user, nil
		} else {
			return nil, errors.New("Login not authorized")
		}
	}

	return nil, err
}

func (this *AuthService) CreateToken(user *entities.User) (string, error) {
	return this.jwtService.encode(user.Id, "token", 15*60)
}

func (this *AuthService) CreateRefreshToken(user *entities.User) (string, error) {
	return this.jwtService.encode(user.Id, "refresh", 24*60*60)
}

func (this *AuthService) GetByToken(token string) (*entities.User, error) {
	uid, kind, err := this.jwtService.decode(token)

	if err != nil {
		return nil, err
	}

	if kind != "token" {
		return nil, errors.New("invalid token type")
	}

	user, err := this.userService.FindOneById(uid)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this *AuthService) RefreshToken(bearerToken string) (string, string, error) {
	uid, kind, err := this.jwtService.decode(bearerToken)

	if err != nil {
		return "", "", err
	}

	if kind != "refresh" {
		return "", "", errors.New("invalid bearerToken type")
	}

	user, err1 := this.userService.FindOneById(uid)

	if err1 != nil {
		return "", "", err
	}

	token, err2 := this.CreateToken(user)

	if err2 != nil {
		return "", "", err
	}

	refreshToken, err3 := this.CreateRefreshToken(user)

	if err3 != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}
