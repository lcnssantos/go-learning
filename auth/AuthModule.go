package auth

import (
	"database/sql"
	"github.com/gorilla/mux"
	"main/auth/controllers"
	"main/auth/provider"
	"main/auth/services"
	"main/user/repository"
	services2 "main/user/services"
)

func BuildAuthModule(db *sql.DB, router *mux.Router) {
	userRepository := repository.NewUserRepository(db)
	hashService := services2.NewHashService()
	userService := services2.NewUserService(userRepository, hashService)
	jwtProviderImpl := provider.NewJwtProviderImpl()
	jwtService := services.NewJwtService(jwtProviderImpl)
	authService := services.NewAuthService(userService, hashService, jwtService)
	authController := controllers.NewAuthController(authService, jwtService)

	router.HandleFunc("/refresh", authController.Refresh).Methods("POST")
	router.Methods("POST").HandlerFunc(authController.Auth)
}
