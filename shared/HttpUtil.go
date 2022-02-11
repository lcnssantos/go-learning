package shared

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ThrowHttpError(w http.ResponseWriter, status int, message string) {
	answer, _ := json.Marshal(HttpError{Message: message, Status: status})
	w.WriteHeader(status)
	w.Write(answer)
	return
}
