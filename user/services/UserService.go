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

func (this *UserService) Create(data *dto.CreateUserDto) (*entities.User, error) {
	_, err := this.repository.FindOneByEmail(data.Email)

	if err == nil {
		return nil, errors.New("Email already exist")
	}

	hash, err := this.hashService.Hash(data.Password)

	if err != nil {
		return nil, err
	}

	data.Password = hash

	err = this.repository.Create(data)

	if err != nil {
		return nil, err
	}

	return this.repository.FindOneByEmail(data.Email)
}

func (this *UserService) FindOneByEmail(email string) (*entities.User, error) {
	return this.repository.FindOneByEmail(email)
}

func (this *UserService) FindOneById(uid string) (*entities.User, error) {
	return this.repository.FindOneById(uid)
}

func NewUserService(userRepository *repository.UserRepository, hashService *HashService) *UserService {
	return &UserService{repository: userRepository, hashService: hashService}
}
