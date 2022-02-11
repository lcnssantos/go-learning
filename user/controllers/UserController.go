package controllers

import (
	"encoding/json"
	"main/shared"
	"main/user/dto"
	"main/user/services"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func (this *UserController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userDto = dto.CreateUserDto{}

	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		shared.ThrowHttpError(w, http.StatusBadRequest, "Invalid body request")
		return
	}

	user, err := this.userService.Create(userDto)

	if err != nil {
		shared.ThrowHttpError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response, _ := json.Marshal(user)
	w.Write(response)
	w.WriteHeader(http.StatusCreated)
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}
