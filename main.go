package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"main/auth"
	"main/shared"
	"main/shared/middleware"
	"main/user"
	"net/http"
)

func main() {
	db, err := shared.Connect()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	validate := validator.New()

	r := mux.NewRouter()
	r.Use(middleware.NewJsonMiddleware().Handler)

	user.BuildUserModule(db, r.PathPrefix("/v1").Subrouter(), validate)
	auth.BuildAuthModule(db, r.PathPrefix("/v1/auth").Subrouter())

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
