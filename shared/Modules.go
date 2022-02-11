package shared

import (
	"github.com/gorilla/mux"
)

type Module struct {
	Routers *mux.Router
}
