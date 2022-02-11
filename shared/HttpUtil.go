package shared

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Status  int
	Message string
}

func ThrowHttpError(w http.ResponseWriter, status int, message string) {
	answer, _ := json.Marshal(HttpError{Message: message, Status: status})
	w.WriteHeader(status)
	w.Write(answer)
	return
}
