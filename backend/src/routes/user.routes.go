package routes

import (
	"github.com/gorilla/mux"
	"github.com/tofu345/Building-mgmt-backend/src/apis"
	mdw "github.com/tofu345/Building-mgmt-backend/src/middleware"
	sc "github.com/tofu345/Building-mgmt-backend/src/schemas"
)

func userRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()
	router.Use(mdw.AuthenticationRequired)

	r.HandleFunc("/token", mdw.ValidateSchema(sc.GetTokenPair, apis.GenerateTokenPair)).Methods("POST")
	r.HandleFunc("/token/refresh", mdw.ValidateSchema(sc.RefreshToken, apis.RegenerateAccessToken)).Methods("POST")

	router.HandleFunc("", apis.GetUserList).Methods("GET")
	router.HandleFunc("/info", apis.GetUserInfo).Methods("GET")
}
