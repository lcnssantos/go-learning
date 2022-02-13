package user

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"main/auth/middlewares"
	"main/auth/provider"
	services2 "main/auth/services"
	"main/user/controllers"
	"main/user/repository"
	"main/user/services"
)

func BuildUserModule(db *sql.DB, router *mux.Router, validate *validator.Validate) {
	userRepository := repository.NewUserRepository(db)
	hashService := services.NewHashService()
	userService := services.NewUserService(userRepository, hashService)
	userController := controllers.NewUserController(userService, validate)
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
