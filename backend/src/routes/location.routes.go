package routes

import (
	"github.com/gorilla/mux"
	"github.com/tofu345/Building-mgmt-backend/src/apis"
	mdw "github.com/tofu345/Building-mgmt-backend/src/middleware"
	sc "github.com/tofu345/Building-mgmt-backend/src/schemas"
)

func locationRoutes(r *mux.Router) {
	router := r.PathPrefix("/locations").Subrouter()
	router.Use(mdw.AuthenticationRequired)

	router.HandleFunc("", apis.GetLocations).Methods("GET")
	router.HandleFunc("", mdw.ValidateSchema(sc.CreateLocation, apis.CreateLocation)).Methods("POST")
	router.HandleFunc("/{id}", apis.GetLocation).Methods("GET")
	router.HandleFunc("/{id}", apis.UpdateLocation).Methods("PUT")

	router.HandleFunc("/{id}/rooms", apis.GetLocationRooms).Methods("GET")
	router.HandleFunc("/{id}/rooms", mdw.ValidateSchema(sc.CreateRoom, apis.CreateRoomForLocation)).Methods("POST")
}
