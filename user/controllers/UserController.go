package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"main/shared"
	"main/user/dto"
	"main/user/services"
	"net/http"
)

type UserController struct {
	userService *services.UserService
	validate    *validator.Validate
}

func (this *UserController) Create(w http.ResponseWriter, r *http.Request) {
	userDto := &dto.CreateUserDto{}

	validationErr := shared.HandleValidateRequest(w, r, userDto)

	if validationErr != nil {
		return
	}

	user, creationErr := this.userService.Create(userDto)

	if creationErr != nil {
		shared.ThrowHttpError(w, http.StatusUnprocessableEntity, creationErr.Error())
		return
	}

	response, _ := json.Marshal(user)
	w.Write(response)
	w.WriteHeader(http.StatusCreated)
}

func NewUserController(userService *services.UserService, validate *validator.Validate) *UserController {
	return &UserController{userService: userService, validate: validate}
}
