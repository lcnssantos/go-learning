package controllers

import (
	"main/shared"
	"main/user/services"
	"net/http"
)

type MeController struct {
	userService *services.UserService
}

func (this *MeController) Name(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	shared.SendHttpResponse(w, http.StatusOK, user)
}

func NewMeController(userService *services.UserService) *MeController {
	return &MeController{userService: userService}
}
