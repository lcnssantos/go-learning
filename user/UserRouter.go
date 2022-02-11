package user

import (
	"database/sql"
	"github.com/gorilla/mux"
	"main/user/controllers"
	"main/user/repository"
	"main/user/services"
)

func BuildUserRouter(db *sql.DB) *mux.Router {
	userRepository := repository.NewUserRepository(db)
	hashService := services.NewHashService()
	userService := services.NewUserService(userRepository, hashService)
	userController := controllers.NewUserController(userService)

	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/", userController.Create).Methods("POST")

	return userRouter
}
