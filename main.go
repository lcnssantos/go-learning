package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/shared"
	"main/user/controllers"
	"main/user/repository"
	"main/user/services"
	"net/http"
)

func main() {
	db, err := shared.Connect()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	hashService := services.NewHashService()
	userService := services.NewUserService(userRepository, hashService)
	userController := controllers.NewUserController(userService)

	router := mux.NewRouter()
	router.HandleFunc("/user", userController.Create)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
