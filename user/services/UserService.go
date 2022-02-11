package services

import (
	"errors"
	"main/user/dto"
	"main/user/entities"
	"main/user/repository"
)

type UserService struct {
	repository  *repository.UserRepository
	hashService *HashService
}

func (this *UserService) Create(data dto.CreateUserDto) (entities.User, error) {
	_, err := this.repository.FindOneByEmail(data.Email)

	if err == nil {
		return entities.User{}, errors.New("Email already exist")
	}

	hash, err := this.hashService.Hash(data.Password)

	if err != nil {
		return entities.User{}, err
	}

	data.Password = hash

	err = this.repository.Create(data)

	if err != nil {
		return entities.User{}, err
	}

	return this.repository.FindOneByEmail(data.Email)
}

func NewUserService(userRepository *repository.UserRepository, hashService *HashService) *UserService {
	return &UserService{repository: userRepository, hashService: hashService}
}
