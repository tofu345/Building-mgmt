package routes

import (
	"github.com/gorilla/mux"
	"github.com/tofu345/Building-mgmt-backend/src/apis"
	"github.com/tofu345/Building-mgmt-backend/src/middleware"
)

func roomRoutes(r *mux.Router) {
	router := r.PathPrefix("/rooms").Subrouter()
	router.Use(middleware.AuthenticationRequired)

	router.HandleFunc("/{id}", apis.GetRoom).Methods("GET")
	router.HandleFunc("/{id}", apis.UpdateRoom).Methods("PUT")
}
