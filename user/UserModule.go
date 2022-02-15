package user

import (
	"database/sql"
	"main/auth/middlewares"
	"main/auth/provider"
	services2 "main/auth/services"
	"main/user/controllers"
	"main/user/repository"
	"main/user/services"

	"github.com/gorilla/mux"
)

func BuildUserModule(db *sql.DB, router *mux.Router) {
	userRepository := repository.NewUserRepository(db)
	hashService := services.NewHashService()
	userService := services.NewUserService(userRepository, hashService)
	userController := controllers.NewUserController(userService)
	meController := controllers.NewMeController(userService)
	jwtProviderImpl := provider.NewJwtProviderImpl()
	jwtService := services2.NewJwtService(jwtProviderImpl)
	authService := services2.NewAuthService(userService, hashService, jwtService)

	meRouter := router.PathPrefix("/me").Subrouter()
	meRouter.Use(middlewares.NewAuthenticationMiddleware(authService).Handler)
	meRouter.Methods("GET").HandlerFunc(meController.Name)

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Methods("POST").HandlerFunc(userController.Create)
}
