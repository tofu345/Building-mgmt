package routes

import (
	"github.com/gorilla/mux"

	"github.com/tofu345/Building-mgmt-backend/src/apis"
	"github.com/tofu345/Building-mgmt-backend/src/middleware"
)

func RegisterRoutes(r *mux.Router) {
	auth := r.PathPrefix("/").Subrouter()
	auth.Use(middleware.AuthRequired)

	r.HandleFunc("/token", apis.GenerateTokenPair).Methods("POST")
	r.HandleFunc("/token/refresh", apis.RegenerateAccessToken).Methods("POST")
	auth.HandleFunc("/users", apis.GetUserList).Methods("GET")

	auth.HandleFunc("/rooms/{id}", apis.GetRoom).Methods("GET")
	auth.HandleFunc("/rooms/{id}", apis.UpdateRoom).Methods("PUT")

	auth.HandleFunc("/locations", apis.GetLocations).Methods("GET")
	auth.HandleFunc("/locations", apis.CreateLocation).Methods("POST")
	auth.HandleFunc("/locations/{id}", apis.GetLocation).Methods("GET")
	auth.HandleFunc("/locations/{id}", apis.UpdateLocation).Methods("PUT")
	auth.HandleFunc("/locations/{id}/rooms", apis.GetLocationRooms).Methods("GET")
	auth.HandleFunc("/locations/{id}/rooms", apis.CreateRoomForLocation).Methods("POST")
}
