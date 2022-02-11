package services

import (
	"errors"
	"main/user/entities"
	"main/user/services"
)

type AuthService struct {
	userService *services.UserService
	hashService *services.HashService
}

func NewAuthService(userService *services.UserService, hashService *services.HashService) *AuthService {
	return &AuthService{userService: userService, hashService: hashService}
}

func (this *AuthService) validate(email string, password string) (*entities.User, error) {
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
