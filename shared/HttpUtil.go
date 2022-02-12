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

func SendHttpResponse(w http.ResponseWriter, status int, data any) {
	response, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(response)
	return
}

func SetJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
