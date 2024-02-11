package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	locationRoutes(r)
	roomRoutes(r)
	userRoutes(r)
}
