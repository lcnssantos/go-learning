package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/shared"
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

	userRouter := user.BuildUserRouter(db)

	r := mux.NewRouter()
	r.Handle("/v1/user", userRouter)

	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}
