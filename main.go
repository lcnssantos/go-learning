package main

import (
	"log"
	"main/auth"
	"main/shared"
	"main/shared/middleware"
	"main/user"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db, err := shared.Connect()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	r := mux.NewRouter()
	r.Use(middleware.NewJsonMiddleware().Handler)

	user.BuildUserModule(db, r.PathPrefix("/v1").Subrouter())
	auth.BuildAuthModule(db, r.PathPrefix("/v1/auth").Subrouter())

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
